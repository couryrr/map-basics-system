package queue

import rl "github.com/gen2brain/raylib-go/raylib"

type Topic string

type EventKind int

type Event struct {
	Kind     EventKind
	Position rl.Vector2
	Key      int
	Consumed bool
}

type callback func(*Event)

type subscriber struct {
	priority int
	callback callback
}

type EventQueue struct {
	subscribers []subscriber
	pending     []*Event
}

func NewEventQueue() EventQueue {
	return EventQueue{
		subscribers: make([]subscriber, 0),
	}
}

func (eq *EventQueue) Subscribe(priority int, callback callback) {
	eq.subscribers = append(eq.subscribers, subscriber{
		priority: priority,
		callback: callback,
	})
}

func (eq *EventQueue) Push(event *Event) {
	eq.pending = append(eq.pending, event)
}

func (eq *EventQueue) Drain() {
	for _, subscriber := range eq.subscribers {
		for _, event := range eq.pending {
			if event.Consumed {
				break
			}
			subscriber.callback(event)
		}
	}
	eq.pending = eq.pending[:0]
}
