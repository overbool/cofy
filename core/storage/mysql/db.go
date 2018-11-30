package mysql

import (
	"fmt"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var DB *gorm.DB

type Database struct {
	db *gorm.DB
}

func New() (*Database, error) {
	username := viper.GetString("db.username")
	password := viper.GetString("db.password")
	addr := viper.GetString("db.addr")
	name := viper.GetString("db.name")

	db, err := openDB(username, password, addr, name)
	if err != nil {
		return nil, err
	}

	DB = db

	return &Database{
		db: db,
	}, nil
}

func (d *Database) DB() *gorm.DB {
	return d.db
}

func (d *Database) Close() {
	err := d.db.Close()
	if err != nil {
		panic(err)
	}
}

func openDB(username, password, addr, name string) (*gorm.DB, error) {
	args := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local")

	db, err := gorm.Open("mysql", args)
	if err != nil {
		return nil, err
	}

	db.LogMode(viper.GetBool("db.log"))
	db.DB().SetMaxIdleConns(0)

	return db, nil
}
