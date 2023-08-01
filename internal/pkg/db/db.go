package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"
	"ordersystem/ent"
	"os"
	"sync"
)

type db_client struct {
	Session *ent.Client
	sync.Once
}

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

var Db *db_client

func (c *Config) getConnectionString() string {
	// mysql connection string format
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", c.User, c.Password, c.Host, c.Port, c.Database)
}

func NewDb() (*ent.Client, error) {
	bytes, err := os.ReadFile("internal/config/mysql.yaml")
	if err != nil {
		return nil, err
	}
	c := new(Config)
	err = yaml.Unmarshal(bytes, c)
	if err != nil {
		return nil, err
	}
	Db = new(db_client)
	var returnChan error

	Db.Do(func() {
		client, err := ent.Open("mysql", c.getConnectionString())
		if err != nil {
			returnChan = err
		}
		Db.Session = client
	})
	if returnChan != nil {
		return nil, returnChan
	}
	return Db.Session, nil
}
