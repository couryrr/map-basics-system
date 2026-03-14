package pubsub

type Topic string

type Message struct {
	Id   string
	Data any
}

type Callback func(Message)

type Broker struct {
	callbacks map[Topic][]Callback
}

func NewBroker() Broker {
	return Broker{
		callbacks: make(map[Topic][]Callback),
	}
}

func (b *Broker) Register(topic Topic, subscriber Callback) {
	b.callbacks[topic] = append(b.callbacks[topic], subscriber)
}

func (b *Broker) Send(topic Topic, message Message) {
	for _, callback := range b.callbacks[topic] {
		callback(message)
	}
}
