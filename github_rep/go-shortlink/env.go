package main

import (
	"go_shortlink/api"
	"go_shortlink/utils"
)

// Env  ...
type Env struct {
	S api.Storage
}

func getEnv() *Env {
	r := utils.NewRedisCli("127.0.0.1:6379", "", 3)
	return &Env{S: r}
}
