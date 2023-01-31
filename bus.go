package bus

import (
	"github.com/evolidev/bus/dispatcher"
	"github.com/evolidev/bus/handler"
)

type Bus struct {
	subscriber []handler.Subscriber
	pipelines  []handler.Pipeline
	dispatcher dispatcher.Dispatcher
}

func NewBus(pipelines ...handler.Pipeline) *Bus {
	b := &Bus{
		subscriber: make([]handler.Subscriber, 0),
		pipelines:  pipelines,
		dispatcher: dispatcher.GoRoutine,
	}

	return b
}

func (b *Bus) UseDispatcher(dispatcher dispatcher.Dispatcher) *Bus {
	bus := b.clone()
	bus.dispatcher = dispatcher

	return bus
}

func (b *Bus) Dispatch(o handler.EmitObject) {
	b.dispatcher.Dispatch(b.buildCallTree(o), o)
}

func (b *Bus) buildCallTree(o handler.EmitObject) handler.Subscriber {
	next := b.subscriberHandler()

	for _, pipeline := range b.pipelines {
		tmp := pipeline.Pipeline(next)
		if tmp.IsSubscribedTo(o) {
			next = tmp
		}
	}

	return next
}

func (b *Bus) Subscribe(subscriber handler.Subscriber) {
	b.subscriber = append(b.subscriber, subscriber)
}

func (b *Bus) subscriberHandler() handler.Subscriber {
	return handler.SubscriberFunc(func(object handler.EmitObject) {
		for _, subscriber := range b.subscriber {
			if subscriber.IsSubscribedTo(object) {
				subscriber.Handle(object)
			}
		}
	})
}

func (b *Bus) PipeThrough(pipe handler.Pipeline) *Bus {
	bus := b.clone()
	bus.pipelines = append(b.pipelines, pipe)

	return bus
}

func (b *Bus) clone() *Bus {
	c := NewBus(b.pipelines...)
	c.subscriber = b.subscriber
	c.dispatcher = b.dispatcher

	return c
}

func (b *Bus) PipeOnlyThrough(pipe handler.Pipeline) *Bus {
	bus := b.clone()
	bus.pipelines = make([]handler.Pipeline, 0)
	bus.pipelines = append(bus.pipelines, pipe)

	return bus
}
