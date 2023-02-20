package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type Address struct {
	Street string `xorm:"'street'"`
	City   string `xorm:"'city'"`
	State  string `xorm:"'state'"`
}

type User struct {
	ID      int64   `xorm:"id pk autoincr "`
	Name    string  `json:"name" xorm:"varchar(128)"`
	Age     int     `json:"age" xorm:"TINYINT"`
	Address Address `json:"address" xorm:"addr json"`
}

var mysqlClient *xorm.Engine

func init() {
	dbDsn := "root:12345678@tcp(localhost:3306)/geminidev?charset=utf8"
	var err error
	mysqlClient, err = xorm.NewEngine("mysql", dbDsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = mysqlClient.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Use xorm's Sync2 function to sync the struct with the database.
	// This will create a new table called "user" in the database if it doesn't exist.
	err = mysqlClient.Sync2(new(User))
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}
	fmt.Println("init successfully")
}

func main() {
	insert()
	read()
}

func insert() {
	// Create a new User struct and insert it into the database.
	user := User{
		Name: "John Doe",
		Age:  25,
		Address: Address{
			Street: "123 Main St",
			City:   "New York",
			State:  "NY",
		},
	}

	_, err := mysqlClient.Insert(&user)
	if err != nil {
		fmt.Println(err) // handle error
	}
}

func read() {
	newUser := new(User)
	has, err := mysqlClient.Table("user").Where("id = 1").Get(newUser)
	if !has {
		fmt.Println("can't get user")
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("user.city = %v\n", newUser.Address.City)
}
