package __ObserverMode

import "fmt"

// MsgType 信用卡消息类型
type MsgType int

const (
	ConsumeType MsgType = iota // 消费消息类型
	BillType                   // 账单消息类型
	ExpireType                 // 逾期消息类型
)

// CreditCard 信用卡
type CreditCard struct {
	holder          string                   // 持卡人
	consumeSum      float32                  // 消费总金额
	subscriberGroup map[MsgType][]Subscriber // 根据消息类型分组订阅者
}

// NewCreditCard 指定持卡人创建信用卡
func NewCreditCard(holder string) *CreditCard {
	return &CreditCard{
		holder:          holder,
		subscriberGroup: make(map[MsgType][]Subscriber),
	}
}

// Subscribe 支持订阅多种消息类型
func (c *CreditCard) Subscribe(subscriber Subscriber, msgTypes ...MsgType) {
	for _, msgType := range msgTypes {
		c.subscriberGroup[msgType] = append(c.subscriberGroup[msgType], subscriber)
	}
}

// Unsubscribe 解除订阅多种消息类型
func (c *CreditCard) Unsubscribe(subscriber Subscriber, msgTypes ...MsgType) {
	for _, msgType := range msgTypes {
		if subs, ok := c.subscriberGroup[msgType]; ok {
			c.subscriberGroup[msgType] = removeSubscriber(subs, subscriber)
		}
	}
}

func removeSubscriber(subscribers []Subscriber, toRemove Subscriber) []Subscriber {
	length := len(subscribers)
	for i, subscriber := range subscribers {
		if toRemove.Name() == subscriber.Name() {
			subscribers[length-1], subscribers[i] = subscribers[i], subscribers[length-1]
			return subscribers[:length-1]
		}
	}
	return subscribers
}

// Consume 信用卡消费
func (c *CreditCard) Consume(money float32) {
	c.consumeSum += money
	c.notify(ConsumeType, fmt.Sprintf("尊敬的持卡人%s,您当前消费%.2f元;", c.holder, money))
}

// SendBill 发送信用卡账单
func (c *CreditCard) SendBill() {
	c.notify(BillType, fmt.Sprintf("尊敬的持卡人%s,您本月账单已出，消费总额%.2f元;", c.holder, c.consumeSum))
}

// Expire 逾期通知
func (c *CreditCard) Expire() {
	c.notify(ExpireType, fmt.Sprintf("尊敬的持卡人%s,您本月账单已逾期，请及时还款，总额%.2f元;", c.holder, c.consumeSum))
}

// notify 根据消息类型通知订阅者
func (c *CreditCard) notify(msgType MsgType, message string) {
	if subs, ok := c.subscriberGroup[msgType]; ok {
		for _, sub := range subs {
			sub.Update(message)
		}
	}
}
