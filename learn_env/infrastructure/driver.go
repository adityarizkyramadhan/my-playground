package infrastructure

import "os"

type Driver struct {
	Port string
	Host string
	User string
}

func NewDriver() *Driver {
	return &Driver{
		Port: os.Getenv("DB_PORT"),
		Host: os.Getenv("DB_HOST"),
		User: os.Getenv("DB_USER"),
	}
}
