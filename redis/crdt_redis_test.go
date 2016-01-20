package crdt

import (
	"flag"
	"fmt"
	"github.com/mediocregopher/radix.v2/redis"
	"log"
	"testing"
)

func init() {
	flag.IntVar(&redisPort, "port", 0, "Port of the redis server.")
	flag.StringVar(&zset, "zset", "", "Name of the ZSET to store into redis.")

	flag.Parse()

	client, err = redis.Dial("tcp", fmt.Sprintf("localhost:%d", redisPort))
	if err != nil {
		log.Fatal("Couldn't open a connection to the Redis Server, please check that redis-server is running.")
	}
}

func TestAdd(t *testing.T) {

	expected := "Resp(Int 0)"
	actual := Add(4, "mysql")
	if actual != expected {
		t.Errorf("TestAdd failed\n Expected Value: %s \n Actual Value: %s", expected, actual)
	}
}
