package main

import (
  "encoding/json"
  "os"
)

type Config struct {
  Http struct {
    Port string `json:"port"`
  } `json:"http"`
  Redis RedisConfig `json:"redis"`
}

func NewConfig(path string) (*Config, error) {
  f, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  var c Config
  err = json.NewDecoder(f).Decode(&c)
  if err != nil {
    return nil, err
  }
  return &c, nil
}
