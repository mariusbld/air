package main

import (
  "log"
  
  "github.com/go-redis/redis"
)

type RedisConfig struct {
  Addr     string `json:"addr"`
  Password string `json:"password"`
  Db       int    `json:"db"`
}

func NewRedisClient(config *RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.Db,
	})

	pong, err := client.Ping().Result()
	if err != nil {
    return nil, err
  }
  log.Println(pong)

  return client, nil
}
