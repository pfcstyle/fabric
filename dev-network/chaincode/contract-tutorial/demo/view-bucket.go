package demo

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func ViewBucket() {
	db, err := OpenDB()
	if err != nil {
		CloseDB(db)
		return
	}
	// 创建表
	err = db.View(func(tx *bolt.Tx) error {

		// 创建BlockBucket表
		b := tx.Bucket([]byte("BlockBucket"))
		if b == nil {
			return nil
		} else {
			data := b.Get([]byte("l"))
			fmt.Printf("l:%s\n", data)
			data = b.Get([]byte("ll"))
			fmt.Printf("ll:%s\n", data)
		}

		// 返回nil，以便数据库处理相应操作
		return nil
	})
	//view失败
	if err != nil {
		log.Panic(err)
	}
	CloseDB(db)
}
