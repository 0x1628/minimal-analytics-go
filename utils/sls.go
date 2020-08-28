package utils

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"reflect"
	"time"

	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"github.com/fragment0/minimal-analytics-go/gen"
	"github.com/gogo/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Callback struct {
}

func (callback *Callback) Success(result *producer.Result) {
	attemptList := result.GetReservedAttempts() // 遍历获得所有的发送记录
	for _, attempt := range attemptList {
		fmt.Println(attempt)
	}
}

func (callback *Callback) Fail(result *producer.Result) {
	fmt.Println(result.IsSuccessful())        // 获得发送日志是否成功
	fmt.Println(result.GetErrorCode())        // 获得最后一次发送失败错误码
	fmt.Println(result.GetErrorMessage())     // 获得最后一次发送失败信息
	fmt.Println(result.GetReservedAttempts()) // 获得producerBatch 每次尝试被发送的信息
	fmt.Println(result.GetRequestId())        // 获得最后一次发送失败请求Id
	fmt.Println(result.GetTimeStampMs())      // 获得最后一次发送失败请求时间
}

type SLSTarget struct {
	config   ConfigItem
	producer *producer.Producer
}

func InitPudaSLS() *SLSTarget {
	producerConfig := producer.GetDefaultProducerConfig()
	config := GetConfig().Puda
	producerConfig.AccessKeyID = config.AccessKeyID
	producerConfig.AccessKeySecret = config.AccessKeySecret
	producerConfig.Endpoint = config.Endpoint

	producerInstance := producer.InitProducer(producerConfig)
	ch := make(chan os.Signal)
	signal.Notify(ch)
	producerInstance.Start()

	return &SLSTarget{
		config:   config,
		producer: producerInstance,
	}
}

func (s *SLSTarget) Send(topic string, source string, content []*sls.LogContent) {
	logcontent := &sls.Log{
		Time:     proto.Uint32(uint32(time.Now().Unix())),
		Contents: content,
	}
	err := s.producer.SendLogWithCallBack(s.config.ProjectName, s.config.LogstoreName, topic, source, logcontent, &Callback{})

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
				Key:   proto.String("DeviceId"),
				Value: proto.String(user.DeviceId),
			})
			content = append(content, &sls.LogContent{
				Key:   proto.String("ServiceId"),
				Value: proto.String(user.ServiceId),
			})
			continue
		}

		if typeName == "Time" {
			time := value.Interface().(*timestamppb.Timestamp)
			content = append(content, &sls.LogContent{
				Key:   proto.String("Time"),
				Value: proto.String(time.String()),
			})
			continue
		}

		content = append(content, &sls.LogContent{
			Key:   proto.String(typeName),
			Value: proto.String(value.Interface().(string)),
		})
	}

	return content
}
