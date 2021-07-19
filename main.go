package main

import (
	"github.com/fragment0/minimal-analytics-go/server"
	"github.com/fragment0/minimal-analytics-go/util"
)

func main() {
	util.InitSls()
	server.Start()
}
