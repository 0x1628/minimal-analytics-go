package util

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"time"

	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"github.com/fragment0/minimal-analytics-go/gen"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Callback struct {
}

func (callback *Callback) Success(result *producer.Result) {
	/*
		attemptList := result.GetReservedAttempts() // 遍历获得所有的发送记录
		for _, attempt := range attemptList {
			fmt.Println(attempt)
		} */
}

func (callback *Callback) Fail(result *producer.Result) {
	fmt.Println(result.IsSuccessful())        // 获得发送日志是否成功
	fmt.Println(result.GetErrorCode())        // 获得最后一次发送失败错误码
	fmt.Println(result.GetErrorMessage())     // 获得最后一次发送失败信息
	fmt.Println(result.GetReservedAttempts()) // 获得producerBatch 每次尝试被发送的信息
	fmt.Println(result.GetRequestId())        // 获得最后一次发送失败请求Id
	fmt.Println(result.GetTimeStampMs())      // 获得最后一次发送失败请求时间
}

type SlsTarget struct {
	config   SlsConfig
	producer *producer.Producer
}

var slsTarget SlsTarget

func InitSls() {
	config := GetConfig()
	producerConfig := producer.GetDefaultProducerConfig()
	producerConfig.AccessKeyID = config.Sls.AccessKeyID
	producerConfig.AccessKeySecret = config.Sls.AccessKeySecret
	producerConfig.Endpoint = config.Sls.Endpoint

	producerInstance := producer.InitProducer(producerConfig)
	ch := make(chan os.Signal)
	signal.Notify(ch)
	producerInstance.Start()

	slsTarget = SlsTarget{
		config:   config.Sls,
		producer: producerInstance,
	}
}

func GetSlsTarget() *SlsTarget {
	return &slsTarget
}

func (s *SlsTarget) getSlsLogStore(topic string) string {
	switch topic {
	case "app_crash":
		return s.config.CrashLogstore
	case "app_register":
		return s.config.RegisterLogstore
	case "app_event":
		return s.config.EventLogstore
	}
	return ""
}

func (s *SlsTarget) Send(topic string, source string, content []*sls.LogContent) {
	logcontent := &sls.Log{
		Time:     proto.Uint32(uint32(time.Now().Unix())),
		Contents: content,
	}
	logstore := s.getSlsLogStore(topic)
	if logstore == "" {
		log.Println("error log topic")
		return
	}
	err := s.producer.SendLogWithCallBack(s.config.ProjectName, logstore, topic, source, logcontent, &Callback{})

	if err != nil {
		log.Println(err)
	}
}

func MakeLogContent(data interface{}) []*sls.LogContent {
	content := []*sls.LogContent{}

	v := reflect.ValueOf(data)
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i)
		if !value.CanInterface() {
			continue
		}

		typeName := v.Type().Field(i).Name
		if typeName == "User" {
			user := value.Interface().(*gen.MAUser)
			content = append(content, &sls.LogContent{
				Key:   proto.String("device_id"),
				Value: proto.String(user.DeviceId),
			})
			content = append(content, &sls.LogContent{
				Key:   proto.String("service_id"),
				Value: proto.String(user.ServiceId),
			})
			content = append(content, &sls.LogContent{
				Key:   proto.String("extra_id"),
				Value: proto.String(user.ExtraId),
			})
			continue
		}

		if typeName == "Time" {
			t := value.Interface().(*timestamppb.Timestamp)
			content = append(content, &sls.LogContent{
				Key:   proto.String("time"),
				Value: proto.String(time.Unix(t.Seconds, 8).String()),
			})
			continue
		}

		jsonName := v.Type().Field(i).Tag.Get("json")
		jsonName = strings.Split(jsonName, ",")[0]
		if jsonName != "" {
			typeName = jsonName
		}

		content = append(content, &sls.LogContent{
			Key:   proto.String(typeName),
			Value: proto.String(value.Interface().(string)),
		})
	}

	content = append(content, &sls.LogContent{
		Key:   proto.String("sys_time"),
		Value: proto.String(fmt.Sprintf("%v", time.Now().Unix())),
	})

	return content
}
