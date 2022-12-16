package __ObserverMode

import "fmt"

// Subscriber 订阅者接口
type Subscriber interface {
	Name() string          //订阅者名称
	Update(message string) //订阅更新方法
}

// shortMessage 信用卡消息短信订阅者
type shortMessage struct{}

func (s *shortMessage) Name() string {
	return "手机短息"
}

func (s *shortMessage) Update(message string) {
	fmt.Printf("通过【%s】发送消息:%s\n", s.Name(), message)
}

// email 信用卡消息邮箱订阅者
type email struct{}

func (e *email) Name() string {
	return "电子邮件"
}

func (e *email) Update(message string) {
	fmt.Printf("通过【%s】发送消息:%s\n", e.Name(), message)
}

// telephone 信用卡消息电话订阅者
type telephone struct{}

func (t *telephone) Name() string {
	return "电话"
}

func (t *telephone) Update(message string) {
	fmt.Printf("通过【%s】告知:%s\n", t.Name(), message)
}
