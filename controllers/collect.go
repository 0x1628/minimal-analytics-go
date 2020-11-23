package controllers

import (
	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/fragment0/minimal-analytics-go/gen"
	"github.com/fragment0/minimal-analytics-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
)

var slsMap map[string]*utils.SLSTarget = make(map[string]*utils.SLSTarget)

func getSls(c *gin.Context) *utils.SLSTarget {
	ns := c.Param("namespace")

	sls := slsMap[ns]
	if sls == nil {
		if ns == "puda" {
			sls = utils.InitPudaSLS()
			slsMap["puda"] = sls
		}
	}

	return sls
}

func parseData(c *gin.Context, target proto.Message) error {
	body, err := c.GetRawData()
	if err != nil {
		return err
	}
	err = proto.UnmarshalMerge(body, target)
	if err != nil {
		return err
	}
	return nil
}

func Events(c *gin.Context) {
	s := getSls(c)

	if s == nil {
		c.Status(404)
		return
	}

	maevent := &gen.MAEvent{}
	err := parseData(c, maevent)
	if err != nil || maevent.User == nil {
		c.Status(403)
		return
	}

	contents := utils.MakeLogContent(*maevent)
	contents = append(contents, &sls.LogContent{
		Key:   proto.String("type"),
		Value: proto.String("events"),
	})

	s.Send("app_events", c.ClientIP(), contents)
}

func Register(c *gin.Context) {
	s := getSls(c)

	if s == nil {
		c.Status(404)
		return
	}

	maregister := &gen.MARegister{}
	err := parseData(c, maregister)
	if err != nil || maregister.User == nil {
		c.Status(403)
		return
	}

	contents := utils.MakeLogContent(*maregister)
	contents = append(contents, &sls.LogContent{
		Key:   proto.String("type"),
		Value: proto.String("register"),
	})
	s.Send("app_register", c.ClientIP(), contents)
}

func Crash(c *gin.Context) {
	s := getSls(c)

	if s == nil {
		c.Status(404)
	}

	macrash := &gen.MACrash{}
	err := parseData(c, macrash)
	if err != nil || macrash.User == nil {
		c.Status(403)
		return
	}

	contents := utils.MakeLogContent(*macrash)
	contents = append(contents, &sls.LogContent{
		Key: proto.String("type"),
		Value: proto.String("crash"),
	})
	s.Send("app_crash", c.ClientIP(), contents)
}
