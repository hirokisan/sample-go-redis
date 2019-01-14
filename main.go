package main

import (
	"encoding/json"
	"log"

	"github.com/gomodule/redigo/redis"
)

type Ticker struct {
	Symbol string
	Bid    string
	Ask    string
}

func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	ticker := &Ticker{Symbol: "USDJPY", Bid: "108.08", Ask: "108.07"}
	serialized, _ := json.Marshal(ticker)

	c.Do("SET", "test", serialized)

	data, _ := redis.Bytes(c.Do("GET", "test"))

	if data != nil {
		deserialized := new(Ticker)
		json.Unmarshal(serialized, deserialized)
		log.Println("deserialized : ", deserialized)
	}
}
