package main

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
)

func main() {
	//db, err := bolt.Open("./day01_03_boltdb/QB.db", 0600, nil)
	db, err := bolt.Open("./QB.db", 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	//读取数据：
	err = db.View(func(tx *bolt.Tx) error {
		//打开桶
		b := tx.Bucket([]byte("QBbucket"))

		if b != nil {
			//读取数据
			data := b.Get([]byte("k")) //[]byte
			fmt.Println(data)
			fmt.Println(string(data))

			data2 := b.Get([]byte("kk"))
			fmt.Println(data2)
			fmt.Println(string(data2))
		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

}
