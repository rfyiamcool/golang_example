package main

import (
	"fmt"
	"github.com/layidao/alimq"
	"time"
)

func main() {
	alimq.ACCESS_KEY = "xxxxxxx"
	alimq.SECRET_KEY = "xxxxxxx"
	alimq.URL_PREFIX = "http://publictest-rest.ons.aliyun.com"
	for {
		List()
	}

}

// 获取消息
func List() {
	t := GetCurrentTime() * 1000

	messages := alimq.Messages{
		Topic:      "TOPIC_TEST_GIFT",
		Tag:        "http",
		ConsumerId: "CID_TEST_GIFT",
		Time:       t,
	}

	list, _ := messages.List()
	if list != nil {
		for k, v := range *list {
			fmt.Println(k, " ---> ", v.Body)

			if v.MsgHandle != "" {
				// 删除消息
				if done, err := messages.Delete(v.MsgHandle); done {
					fmt.Println("删除成功")
				} else {
					fmt.Println(err)
				}
			}

		}
	}
}

// 发送消息
func Send() {
	t := GetCurrentTime() * 1000

	message := alimq.SendMessage{
		Topic:      "TOPIC_TEST_GIFT",
		Tag:        "http",
		ProducerId: "PID_TEST_GIFT",
		Key:        "http",
		Body:       "aaaaa",
		Time:       t,
	}

	messageId, err := message.Send()
	fmt.Println(messageId)
	fmt.Println(err)
}
