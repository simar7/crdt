package crdt

import (
	"flag"
	"fmt"
	"github.com/mediocregopher/radix.v2/redis"
	"log"
)

var redisPort int
var zset string
var client *redis.Client
var err error

func Add(value int, elemName string) {
	resp := client.Cmd(fmt.Sprintf("ZADD"), zset, value, elemName)
	if resp.Err != nil {
		log.Fatal(fmt.Sprintf("Command failed: %s\n", resp.Err))
	} else {
		fmt.Printf("%s\n", resp.String())
	}
}

func PrintAll(minValue int, maxValue int, constraints string) {
	resp := client.Cmd(fmt.Sprintf("ZRANGE"), zset, minValue, maxValue, constraints)
	if resp.Err != nil {
		log.Fatal(fmt.Sprintf("Command failed: %s\n", resp.Err))
	} else {
		fmt.Printf("%s\n", resp.String())
	}
}

func checkInput() {
	if redisPort == 0 {
		log.Fatal("Please specify the port at which the redis server is running using --port option.")
	} else if zset == "" {
		log.Fatal("Please specify the name of the ZSET to store using the --zset option.")
	}
}

func init() {
	flag.IntVar(&redisPort, "port", 0, "Port of the redis server.")
	flag.StringVar(&zset, "zset", "", "Name of the ZSET to store into redis.")
}

func main() {
	flag.Parse()

	checkInput()

	client, err = redis.Dial("tcp", fmt.Sprintf("localhost:%d", redisPort))
	if err != nil {
		log.Fatal("Couldn't open a connection to the Redis Server, please check that redis-server is running.")
	}

	//PrintAll(0, 10, "WITHSCORES")
}
