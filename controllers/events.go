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

func Events(c *gin.Context) {
	s := getSls(c)

	if s == nil {
		c.Status(404)
		return
	}

	maevent := &gen.MAEvent{}
	body, err := c.GetRawData()
	if err != nil {
		c.Status(403)
		return
	}
	err = proto.UnmarshalMerge(body, maevent)
	if err != nil || maevent.User == nil {
		c.Status(403)
		return
	}

	contents := utils.MakeLogContent(maevent)
	contents = append(contents, &sls.LogContent{
		Key:   proto.String("type"),
		Value: proto.String("events"),
	})

	s.Send("test", c.ClientIP(), contents)
}

func Register(c *gin.Context) {
	s := getSls(c)

	if s == nil {
		c.Status(404)
		return
	}

	contents := utils.MakeLogContent(gen.MARegister{
		User: &gen.MAUser{
			DeviceId:  "123",
			ServiceId: "456",
		},
		Os: "mac",
	})
	contents = append(contents, &sls.LogContent{
		Key:   proto.String("type"),
		Value: proto.String("register"),
	})
	s.Send("test", c.ClientIP(), contents)
}

func BatchLog(c *gin.Context) {

}
