package main

import (
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	/*
	step1：安装
		go get "github.com/boltdb/bolt"


	step2：打开一个数据库

	go run xxx.go
		--->直接执行，相对于当前的工程下

	go build -o 别名 xxx.go

	 */
	 //db,err:=bolt.Open("./day01_03_boltdb/QB.db",0600,nil)
	db,err:=bolt.Open("./QB.db",0600,nil)//存在就打开，不存在就创建
	 if err != nil{
	 	log.Panic(err)
	 }
	 defer db.Close()

	 /*
	 db.Update(),读和写
	 db.View()，只读


	  */

}

/*
执行程序：
	1、程序上右键执行： QB.db会被创建在当前目录根目录下
	go run main.go
	在程序中产生的文件，相对于工程的路径

	2、将程序进行编译——>产生可执行的文件  QB.db会被创建在当前文件所在目录下
	go build main.go                       能正常生成demo01_安装和打开.exe
	go build -o QB main.go //QB为名字了    不能生成QB.exe？可能是win版本高的原因，一般情况下没图标等信息的exe打包会出问题

	mac，linux——>xxx
	windows——>xxx.exe
	相当于当前的可执行文件的位置。

*/
