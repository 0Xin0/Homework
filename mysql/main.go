package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //要导入相应驱动包，否则会报错
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initDB() {
	var err error

	dsn := "root:123456@tcp(127.0.0.1:3306)/student?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}

	fmt.Println("connecting to MySQL...")
	return
} //连接

func insert(name string, age int) {
	sqlStr := "insert into student(name, age) value (?,?)"
	ret, err := db.Exec(sqlStr, name, age)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
} //添加

type student struct {
	id   int
	age  int
	name string
}

func queryMulti(num int) {
	sqlStr := "select id, name, age from student where id >= ?"
	rows, err := db.Query(sqlStr, num)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u student
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
} //查询

func main() {
	//初始化连接
	initDB()
	insert("初始化明", 0)
	insert("小明", 10)
	insert("中明", 20)
	insert("大明", 30)
	insert("大大明", 40)
	insert("超大明", 50)
	insert("老明", 60)
	insert("大老明", 70)
	insert("老老老明", 80)
	insert("巨老明", 90)
	insert("终极明", 100)
	queryMulti(1)
}
