package main

import (
	"context"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"time"
)

type Person struct {
	Id   int64
	Name string
	Age  int
}

func init() {
	orm.RegisterModel(new(Person))

	err := orm.RegisterDataBase("default", "sqlite3", "./test.db")
	fmt.Println(err)
	err = orm.RegisterDriver("sqlite3", orm.DRSqlite)
	if err != nil {
		return
	}
	//自动建表
	err = orm.RunSyncdb("default", false, true) //第一个true 是 强制重建表结构， 第二个true是 显示命令操作是的信息
	if err != nil {
		fmt.Println("init database table error:", err)
	}

}

func main() {

	var err error
	deletePoint()
	for i := 0; i < 1000; i++ {
		persons := make([]*Person, 0)

		for j := 0; j < 10000; j++ {
			persons = append(persons, &Person{
				Name: strconv.Itoa(j + i),
				Age:  i * j,
			})
		}

		personChannel <- persons
	}
	close(personChannel)

	//_, err = orm.NewOrm().Raw("delete from sqlite_sequence where name='person';").Exec()
	//fmt.Println("err",err)
	_, err = orm.NewOrm().Raw("update sqlite_sequence SET seq = 1024 where name = 'person';").Exec()
	fmt.Println("err", err)

	if err != nil {
		return
	}
	//InsertP(persons)
	go InsertByChannel()
	for true {
		time.Sleep(time.Second)
		fmt.Println(",,,,,,,,")
	}
}

func deletePoint() {
	exec, err := orm.NewOrm().Raw("delete from person;").Exec()
	fmt.Println(exec, err)
}

type Message struct {
	ID     int32
	PN     string
	Number int32
	Flag   string
}

func InsertP(persons []*Person) {
	o := orm.NewOrm()
	t := time.Now()

	err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		var err error
		for i := 0; i < len(persons); i++ {
			_, err = txOrm.Insert(persons[i])
			if err != nil {
				fmt.Println(err)
			}
		}
		return err
	})
	fmt.Println(time.Since(t), "insert point length:", len(persons))

	if err != nil {
		fmt.Println(err)
	}

}

var personChannel = make(chan []*Person, 1000)

func InsertByChannel() {
	now := time.Now()
	for people := range personChannel {
		InsertP(people)
	}
	fmt.Println(time.Since(now))

}
