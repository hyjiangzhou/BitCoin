package main

import (
	"../bolt"
	"log"
	"fmt"
)

func  main()  {
	//1. 打开数据库（如果没有会创建)
	//110
	//rwx
	db, err := bolt.Open("test.db", 0600, nil)

	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		//2. 找到我们的桶，通过桶的名字
		// Returns nil if the bucket does not exist.
		bucket := tx.Bucket([]byte("bucketName1"))

		//如果没有找到，先创建
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte("bucketName1"))
			if err != nil {
				log.Panic(err)
			}
		}

		//3. 写数据
		err = bucket.Put([]byte("111"), []byte("hello"))
		if err != nil {
			log.Panic(err)
		}

		err = bucket.Put([]byte("222"), []byte("world"))
		if err != nil {
			log.Panic(err)
		}

		//4. 读数据
		data1 := bucket.Get([]byte("111"))
		data2 := bucket.Get([]byte("222"))
		data3 := bucket.Get([]byte("333"))

		fmt.Printf("data1 : %s\n", data1)
		fmt.Printf("data2 : %s\n", data2)
		fmt.Printf("data3 : %s\n", data3)

		return nil
	})

	db.Close()
}


