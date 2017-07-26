package goutil

import (
	"encoding/json"

	"github.com/nsqio/go-nsq"
)

// PushNSQ 推送到NSQ消息队列中
func PushNSQ(server, topic string, body interface{}) *nsq.Producer {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer(server, config)
	bodyJSON, _ := json.Marshal(body)
	w.Publish(topic, bodyJSON)

	return w
}
