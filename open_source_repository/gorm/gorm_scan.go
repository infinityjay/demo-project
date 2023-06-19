package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	mysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Address struct {
	ID   int64  `gorm:"id primaryKey"`
	Name string `gorm:"name"`
	City string `gorm:"'city'"`
}

type User struct {
	ID   int64  `gorm:"id primaryKey"`
	Name string `gorm:"name"`
	Age  int    `gorm:"age"`
}

type FindUser struct {
	Name    string     `gorm:"name"`
	Age     int        `gorm:"age"`
	Address []*Address `gorm:"address"`
}

var mysqlClient *gorm.DB

func InitMysql() error {
	dsn := "root:12345678@tcp(localhost:3306)/localdemo?charset=utf8"
	var err error
	mysqlClient, err = gorm.Open(mysqlDriver.Open(dsn+"&parseTime=true&loc=Local"), &gorm.Config{
		//Logger: daoLogger.LogMode(logger.Info),
		//QueryFields: true,
		//PrepareStmt: true,
	})
	if err != nil {
		return err
	}

	sqlDB, err := mysqlClient.DB()
	if err != nil {
		return err
	}
	//sqlDB.SingularTable(true)
	//sqlDB.LogMode(true)

	sqlDB.SetMaxIdleConns(20)

	return nil
}

func main() {
	InitMysql()
	err := mysqlClient.AutoMigrate(&User{}, &Address{})
	if err != nil {
		fmt.Println(err)
	}
	users := []*User{
		{
			ID:   1,
			Name: "aaa",
			Age:  1,
		},
		{
			ID:   2,
			Name: "bbb",
			Age:  2,
		},
		{
			ID:   3,
			Name: "ccc",
			Age:  3,
		},
	}
	address := []*Address{
		{
			ID:   1,
			Name: "aaa",
			City: "beijing",
		},
		{
			ID:   2,
			Name: "aaa",
			City: "shanghai",
		},
		{
			ID:   3,
			Name: "bbb",
			City: "guangzhou",
		},
		{
			ID:   4,
			Name: "ccc",
			City: "wuhan",
		},
	}
	mysqlClient.Create(users)
	mysqlClient.Create(address)

	var res []*FindUser
	sql := ""
	err = mysqlClient.Raw(sql).Scan(&res).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (i *FindUser) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	return json.Unmarshal(bytes, i)
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (i FindUser) Value() (driver.Value, error) {
	return json.Marshal(i)
}
