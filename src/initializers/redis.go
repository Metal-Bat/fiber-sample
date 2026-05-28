package initializers

import (
	"log"
	"os"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func SetUpRedis() {
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatal(err)
	}

	RedisClient = redis.NewClient(opt)
	if err := redisotel.InstrumentTracing(RedisClient); err != nil {
		log.Fatalf("failed to instrument redis tracing: %v", err)
	}
}
