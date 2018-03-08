package main

import "time"

type AdminUser struct {
	Id       int64`pk`
	Account  string`unique`
	Password string
}
type UserData struct {
	Id       int64`pk`
	Username string`unique`
	Password string
}
type OnlineUser struct {
	Id        int64
	Username  string`unique`
	Logintime time.Time`created`
}
