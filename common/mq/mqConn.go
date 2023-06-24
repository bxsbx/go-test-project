package mq

import (
	"fmt"
	beegoConfig "github.com/astaxie/beego/config"
	"github.com/streadway/amqp"
	"log"
	"sync"
	"time"
)

type mqConn struct {
	Name            string
	MqChannels      []*mqChannel
	MqConn          *amqp.Connection
	Mu              sync.Mutex
	AvailableConn   int
	CurrentConn     int
	IsFull          chan bool
	MaxOpenConn     int
	MaxIdleConn     int
	ConnMaxLifetime int
}

type mqChannel struct {
	channel  *amqp.Channel
	LifeTime int64
	InUser   bool
}

var mqPoolMap = make(map[string]*mqConn)

const (
	DEFAULT = "default"
)

func InitMq(cfg beegoConfig.Configer) {
	mqPoolMap[DEFAULT] = newMqPool(defaultMqConfig(cfg), DEFAULT)
}

func DefaultMq() *mqConn {
	return mqPoolMap[DEFAULT]
}

func GetMq(key string) *mqConn {
	return mqPoolMap[key]
}

func newMqPool(cfg mqConfig, poolName string) *mqConn {
	if cfg.ConnMaxLifetime < 10 {
		cfg.ConnMaxLifetime = 10
	}
	if cfg.MaxOpenConn < 10 {
		cfg.MaxIdleConn = 10
	}
	if cfg.MaxIdleConn < 5 {
		cfg.MaxIdleConn = 5
	}
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.UserName, cfg.Password, cfg.Host, cfg.Port)
	conn, err := amqp.Dial(url)
	if nil != err {
		log.Fatalf("mq初始化失败, err:%v", err)
		return nil
	}
	return &mqConn{
		MqChannels:      make([]*mqChannel, cfg.MaxOpenConn),
		MqConn:          conn,
		IsFull:          make(chan bool, 1),
		Name:            poolName,
		MaxOpenConn:     cfg.MaxOpenConn,
		MaxIdleConn:     cfg.MaxIdleConn,
		ConnMaxLifetime: cfg.ConnMaxLifetime,
	}
}

func (q *mqConn) newChannel() (*mqChannel, error) {
	channel, err := q.MqConn.Channel()
	if err != nil {
		return nil, err
	}
	return &mqChannel{
		channel: channel,
	}, nil
}

// 从连接池获取连接
func (q *mqConn) Get() (mqChannel *mqChannel, err error) {
	for q.CurrentConn >= q.MaxOpenConn && q.AvailableConn <= 0 {
		<-q.IsFull
	}
	q.Mu.Lock()
	defer q.Mu.Unlock()
	if q.AvailableConn > 0 {
		for i := 0; i < q.MaxOpenConn; i++ {
			mqChannel = q.MqChannels[i]
			if !mqChannel.InUser {
				mqChannel.InUser = true
				q.AvailableConn--
				break
			}
		}
	} else {
		mqChannel, err = q.newChannel()
		mqChannel.InUser = true
	}
	return
}

// 如果连接数大于最大空闲连接且过了最大空闲时间则释放连接
func (q *mqConn) timedRelease(mqChannel *mqChannel) {
	<-time.After(time.Duration(q.ConnMaxLifetime) * time.Second)
	q.Mu.Lock()
	defer q.Mu.Unlock()
	if !mqChannel.InUser && time.Now().Unix() > mqChannel.LifeTime {
		for i := 0; i < q.MaxOpenConn; i++ {
			if q.MqChannels[i] == mqChannel {
				mqChannel.channel.Close()
				q.MqChannels[i] = nil
				break
			}
		}
		q.CurrentConn--
		q.AvailableConn--
	}
}

// 放回连接到连接池
func (q *mqConn) Release(mqChannel *mqChannel) {
	for i := 0; i < q.MaxOpenConn; i++ {
		if q.MqChannels[i] == mqChannel {
			q.Mu.Lock()
			count := q.AvailableConn
			q.AvailableConn++
			mqChannel.InUser = false
			mqChannel.LifeTime = -1
			if q.CurrentConn > q.MaxIdleConn {
				mqChannel.LifeTime = time.Now().Add(time.Duration(q.ConnMaxLifetime) * time.Second).Unix()
				go q.timedRelease(mqChannel)
			}
			q.Mu.Unlock()
			for q.CurrentConn >= q.MaxOpenConn && count <= 0 {
				q.IsFull <- true
			}
			return
		}
	}
	mqChannel.channel.Close()
}

// 关闭连接池
func (q *mqConn) Close() {

	delete(mqPoolMap, q.Name)
	q.Mu.Lock()
	defer q.Mu.Unlock()

	for i := 0; i < q.MaxOpenConn; i++ {
		if q.MqChannels[i] != nil {
			q.MqChannels[i].channel.Close()
		}
	}
	q.MqConn.Close()
	q.MqChannels = nil
}
