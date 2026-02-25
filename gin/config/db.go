package config

import "strconv"

type DB struct {
	Type     string
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	Params   string
}

func (d DB) DSN() string {
	return d.User + ":" + d.Password + "@tcp(" + d.Host + ":" + strconv.Itoa(d.Port) + ")/" + d.Name + "?" + d.Params
}
