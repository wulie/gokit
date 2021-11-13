package main

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
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

	persons := make([]*Person, 0)
	for i := 0; i < 50; i++ {
		persons = append(persons, &Person{
			Name: strconv.Itoa(i),
			Age:  i,
		})
	}

	_, err = orm.NewOrm().Raw("update sqlite_sequence SET seq = 1024 where name = 'person';").Exec()
	if err != nil {
		return
	}
	//exec, err := engine.Exec("update sqlite_sequence SET seq = 1024 where name = 'person';")
	//if err != nil {
	//	return
	//}

	//for i := 0; i < len(persons); i++ {
	//	//p.Id= int64(i+1034)
	//	_, err := orm.NewOrm().Insert(persons[i])
	//	fmt.Println(err)
	//	if err != nil {
	//		//fmt.Println(err,".......",i)
	//		return
	//	}
	//}
	_, err = orm.NewOrm().InsertMulti(len(persons),
		persons)
	fmt.Println(err)

	fmt.Println(err)
}
