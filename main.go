package main

import (
	"strconv"
	"strings"

	"github.com/fragment0/minimal-analytics-go/route"
	"github.com/fragment0/minimal-analytics-go/utils"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {
	config := utils.InitConfig()

	r := gin.New()
	route.Route(r)
	s := []string{config.Server.Host, ":", strconv.Itoa(config.Server.Port)}
	endless.ListenAndServe(strings.Join(s, ""), r)
}
