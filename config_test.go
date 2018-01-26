package main

import (
  "testing"
)

func TestNewConfig(t *testing.T) {
  c, err := NewConfig("config/prod.json")
  t.Log(c, err)
}
