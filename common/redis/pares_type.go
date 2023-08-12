package redis

import "github.com/gomodule/redigo/redis"

type ParesType struct {
	reply interface{}
	err   error
}

func GetParesType(reply interface{}, err error) *ParesType {
	return &ParesType{
		reply,
		err,
	}
}

func (p *ParesType) Sting() (string, error) {
	return redis.String(p.reply, p.err)
}

func (p *ParesType) Int() (int, error) {
	return redis.Int(p.reply, p.err)
}
