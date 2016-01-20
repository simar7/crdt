func main() {
	flag.Parse()

	checkInput()

	client, err = redis.Dial("tcp", fmt.Sprintf("localhost:%d", redisPort))
	if err != nil {
		log.Fatal("Couldn't open a connection to the Redis Server, please check that redis-server is running.")
	}

	Add(4, "mysql")
	PrintAll(0, 10, "WITHSCORES")
}
