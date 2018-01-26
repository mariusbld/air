package main

import (
  "log"
  "net/http"
  "os"

  "github.com/gorilla/mux"
  "github.com/go-redis/redis"
)

var server *Server

type Server struct {
  config *Config
  redisClient *redis.Client
}

func NewServer() (*Server, error) {
  configFile := os.Getenv("AIR_CONFIG")
  if configFile == "" {
    configFile = "config/dev.json"
  }
  config, err := NewConfig(configFile)
  if err != nil {
    return nil, err
  }

  redisClient, err := NewRedisClient(&config.Redis)
  if err != nil {
    return nil, err
  }

  return &Server{ config, redisClient }, nil
}

func main() {
  server, err := NewServer()
  if err != nil {
    panic(err)
  }

  r := mux.NewRouter()
  r.HandleFunc("/login/{token}", LoginHandler).Methods("POST")
  r.HandleFunc("/broadcast/{action}", BroadcastHandler).Methods("GET")
  r.HandleFunc("/follow/{email}", FollowHandler).Methods("GEt")
  log.Fatal(http.ListenAndServe(":" + server.config.Http.Port, r))
}
