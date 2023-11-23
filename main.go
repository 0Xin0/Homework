package main

import (
	"fmt"
	"os"
	"time"
)

type timestampWriter struct {
	timestamp time.Time
}

func (tw *timestampWriter) Write(p []byte) (n int, err error) {
	msg := fmt.Sprintf("%v[%s]%s\n", time.Now(), time.Since(tw.timestamp), string(p))
	file, err := os.OpenFile("log", os.O_RDWR, 0666)
	if err != nil {
		return 0, err
	}
	n, err = file.WriteString(msg)
	return
}
func main() {
	_, err := os.OpenFile("log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	//创建一个带时间戳的写入器
	logWriter := &timestampWriter{timestamp: time.Now()}

	//模拟用户操作并记录日志
	_, a := fmt.Fprintln(logWriter, "用户登录")
	fmt.Println(a)
	fmt.Println(time.Now(), time.Now().Unix(), "用户登录")
	time.Sleep(2 * time.Second)
	_, b := fmt.Fprintln(logWriter, "用户执行操作A")
	fmt.Println(b)
	fmt.Println(time.Now(), time.Now().Unix(), "用户执行操作A")
	time.Sleep(1 * time.Second)
	_, c := fmt.Fprintln(logWriter, "用户执行操作B")
	fmt.Println(c)
	fmt.Println(time.Now(), time.Now().Unix(), "用户执行操作B")

	//tw := &timestampWriter{timestamp: time.Now()}
	//for {
	//	time.Sleep(time.Second)
	//	_, e := fmt.Fprint(tw, "init log successfully")
	//	fmt.Println(e)
	//}
}
