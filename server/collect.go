package server

import (
	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/fragment0/minimal-analytics-go/gen"
	"github.com/fragment0/minimal-analytics-go/util"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"net/http"
)

func parseData(request *http.Request, target proto.Message) error {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	err = proto.Unmarshal(body, target)
	if err != nil {
		return err
	}
	return nil
}

func CollectEvent(writer http.ResponseWriter, request *http.Request) {
	var maevent gen.MAEvent
	err := parseData(request, &maevent)

	if err != nil || maevent.User == nil {
		writer.WriteHeader(403)
		return
	}

	contents := util.MakeLogContent(maevent)
	contents = append(contents, &sls.LogContent{
		Key: proto.String("type"),
		Value: proto.String("event"),
	})

	util.GetSlsTarget().Send("app_event", request.RemoteAddr, contents)

	writer.WriteHeader(204)
}

func CollectRegister(writer http.ResponseWriter, request *http.Request) {
	var maregister gen.MARegister
	err := parseData(request, &maregister)

	if err != nil || maregister.User == nil {
		writer.WriteHeader(403)
		return
	}

	contents := util.MakeLogContent(maregister)
	contents = append(contents, &sls.LogContent{
		Key: proto.String("type"),
		Value: proto.String("register"),
	})

	util.GetSlsTarget().Send("app_register", request.RemoteAddr, contents)

	writer.WriteHeader(204)
}

func CollectCrash(writer http.ResponseWriter, request *http.Request) {
	var macrash gen.MACrash
	err := parseData(request, &macrash)

	if err != nil || macrash.User == nil {
		writer.WriteHeader(403)
		return
	}

	contents := util.MakeLogContent(macrash)
	contents = append(contents, &sls.LogContent{
		Key: proto.String("type"),
		Value: proto.String("crash"),
	})

	util.GetSlsTarget().Send("app_crash", request.RemoteAddr, contents)

	writer.WriteHeader(204)
}
