package database

import "gorm.io/gorm"

type Database interface {
	GetConnection() *gorm.DB
}

type database struct {
	connection *gorm.DB
}

func NewDatabase(connection *gorm.DB) Database {
	return &database{connection: connection}
}

func (d *database) GetConnection() *gorm.DB {
	return d.connection
}
