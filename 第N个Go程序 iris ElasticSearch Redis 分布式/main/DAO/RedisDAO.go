package DAO

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type Redis struct {

}

func (d *Redis) Dail() redis.Conn{
	c, err := redis.Dial("tcp", "172.17.0.2:6379")//, options)
	if err != nil {
		fmt.Println(err)
	}
	return c
}

func (d *Redis) Set(key string,value string) interface{}{
	c := d.Dail()
	defer c.Close()
	v, err := c.Do("SET", key, value)
	if err != nil {
		fmt.Println(err)
	}
	return v
}

func (d *Redis) Get(key string) string {
	c := d.Dail()
	defer c.Close()
	v, err := redis.String(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
	}
	return v
}
