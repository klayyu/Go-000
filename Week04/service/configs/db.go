package configs

import "database/sql"

var db *sql.DB


type DbConfig struct {

}

func NewDb(c *DbConfig) (*sql.DB,error) {
	return &sql.DB{},nil
}