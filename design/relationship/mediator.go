package relationship

import "fmt"

type Colleague interface {
	Send(to string, msg string)
	receive(from string, msg string)
}

type ColleagueA struct {
	Name     string
	Mediator Mediator
}

func (c *ColleagueA) Send(to string, msg string) {
	fmt.Printf("%s发送数据给%s,msg:%s\n", c.Name, to, msg)
	c.Mediator.Relay(c.Name, to, msg)
}

func (c *ColleagueA) receive(from string, msg string) {
	fmt.Printf("%s接收数据来自%s,msg:%s\n", c.Name, from, msg)
}

type Mediator interface {
	Add(key string, e Colleague)
	Relay(from, to, msg string)
}

type MediatorA struct {
	ObjMap map[string]Colleague
}

func (m *MediatorA) Add(key string, e Colleague) {
	m.ObjMap[key] = e
}

func (m *MediatorA) Relay(from, to, msg string) {
	m.ObjMap[to].receive(from, msg)
}
