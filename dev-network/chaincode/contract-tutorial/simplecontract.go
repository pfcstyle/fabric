package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SimpleContract contract for handling writing and reading from the world state
type SimpleContract struct {
	contractapi.Contract
}

// OpenDB to open bolt db
func OpenDB() (*bolt.DB, error) {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

// CloseDB to close bolt db
func CloseDB(db *bolt.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// UpdateBucket to update Bucket
func UpdateBucket() {
	db, err := OpenDB()
	if err != nil {
		CloseDB(db)
		return
	}
	// 创建表
	err = db.Update(func(tx *bolt.Tx) error {

		// 创建BlockBucket表
		b := tx.Bucket([]byte("BlockBucket"))
		if b == nil {
			b, err = tx.CreateBucket([]byte("BlockBucket"))
			if err != nil {
				return fmt.Errorf("create bucket: %s", err)
			}
		}

		// 往表里面存储数据
		if b != nil {
			err := b.Put([]byte("l"), []byte("Send 100 BTC To 强哥......"))
			if err != nil {
				log.Panic("数据存储失败......")
			}
		}

		// 返回nil，以便数据库处理相应操作
		return nil
	})
	//更新失败
	if err != nil {
		log.Panic(err)
	}
	CloseDB(db)
}

// ViewBucket to view Bucket
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

// Create adds a new key with value to the world state
func (sc *SimpleContract) Create(ctx CustomTransactionContextInterface, key string, value string) error {
	existing := ctx.GetData()

	if existing != nil {
		return fmt.Errorf("Cannot create world state pair with key %s. Already exists", key)
	}

	err := ctx.GetStub().PutState(key, []byte(value))

	if err != nil {
		return errors.New("Unable to interact with world state")
	}

	UpdateBucket()

	return nil
}

// Update changes the value with key in the world state
func (sc *SimpleContract) Update(ctx CustomTransactionContextInterface, key string, value string) error {
	existing := ctx.GetData()

	if existing == nil {
		return fmt.Errorf("Cannot update world state pair with key %s. Does not exist", key)
	}

	err := ctx.GetStub().PutState(key, []byte(value))

	if err != nil {
		return errors.New("Unable to interact with world state")
	}

	return nil
}

// Read returns the value at key in the world state
func (sc *SimpleContract) Read(ctx CustomTransactionContextInterface, key string) (string, error) {
	existing := ctx.GetData()

	if existing == nil {
		return "", fmt.Errorf("Cannot read world state pair with key %s. Does not exist", key)
	}

	ViewBucket()
	return string(existing), nil
}
