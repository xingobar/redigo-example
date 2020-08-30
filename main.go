package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

// reference: https://blog.csdn.net/wangshubo1989/article/details/75050024
func main() {
	fmt.Println("[*] run redis example")
	c, err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		panic(err)
	}

	defer c.Close()

	_, err = c.Do("SET", "testkey" ,"testvalue")
	if err != nil {
		fmt.Println("redis set failed: ", err)
	}

	value, err := redis.String(c.Do("GET", "testkey"))
	if err != nil {
		fmt.Println("redis get key failed: ", err)
	}
	fmt.Println("redis get testkey value: ", value)

	// 設定過期時間

	// 5秒後過期
	_, err = c.Do("SET", "testkey", "testvalue", "EX", 5)
	if err != nil {
		fmt.Println("set expire time failed: " , err)
	}

	time.Sleep(6 * time.Second)
	value, err = redis.String(c.Do("GET", "testkey"))
	if err != nil {
		fmt.Println("get key failed:  ", err)
	}

	// 確認 key 是否存在
	exists, err := redis.Bool(c.Do("EXISTS", "testkey"))
	if err != nil {
		fmt.Println("failed: ", err)
	} else {
		fmt.Println("key exists or not: ", exists)
	}

	_, err = c.Do("SET", "testkey", "testvalue")
	if err != nil {
		fmt.Println("redis set failed: " , err)
	}

	exists, err = redis.Bool(c.Do("EXISTS", "testkey"))
	if err != nil {
		fmt.Println("failed: ", err)
	} else {
		fmt.Println("key exists or not: ", exists)
	}

	// 刪除 key
	_, err = c.Do("DEL", "testkey")
	if err != nil {
		fmt.Println("redis delete key failed: ", err)
	}

	value, err = redis.String(c.Do("GET", "testkey"))
	if err != nil {
		fmt.Println("redis get value failed: ", err)
	}

	// write json
	imap := map[string]string {
		"username": "test",
		"phone": "value",
	}

	v, _ := json.Marshal(imap)
	n, err := c.Do("SETNX", "profile", v)
	if err != nil {
		fmt.Println("redis setnx failed: " , err)
	} else {
		fmt.Println("success: ", n)
	}

	var imapGet map[string]string

	valueGet, err := redis.Bytes(c.Do("GET", "profile"))
	if err != nil {
		fmt.Println("redis get failed: ", err)
	}

	errSha := json.Unmarshal(valueGet, &imapGet)
	if errSha != nil {
		fmt.Println("unmarshal failed: ", err)
	}

	fmt.Println("username: ", imapGet["username"])
}


