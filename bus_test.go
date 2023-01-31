package bus

import (
	"github.com/evolidev/bus/dispatcher"
	"github.com/evolidev/bus/handler"
	"testing"
	"time"
)

func TestBus_Dispatch(t *testing.T) {
	t.Parallel()
	t.Run("Dispatched object should not be picked up if no one has subscribed to it", func(t *testing.T) {
		bus := newBus()
		bus.Subscribe(&TestSubscriber{})

		e := &NoOneSubscribedToItEvent{}
		bus.Dispatch(e)

		if e.Emitted == true {
			t.Errorf("Event got not emitted")
		}
	})
	t.Run("Emitted object should be caught by handler", func(t *testing.T) {
		bus := newBus()
		bus.Subscribe(&TestSubscriber{})

		e := &TestEvent{}
		bus.Dispatch(e)

		if e.Emitted == false {
			t.Errorf("Event got not emitted")
		}
	})
	t.Run("Catch all subscriber should get all events", func(t *testing.T) {
		bus := newBus()
		sub := &CatchAllSubscriber{Handled: make([]handler.EmitObject, 0)}
		bus.Subscribe(sub)

		bus.Dispatch(&TestEvent{})
		bus.Dispatch(&NoOneSubscribedToItEvent{})

		if len(sub.Handled) != 2 {
			t.Errorf("Not all events were caught")
		}
	})
}

func TestPipeline_Called(t *testing.T) {
	t.Parallel()
	t.Run("Check if global pipeline was called", func(t *testing.T) {
		pipe := &TestPipeline{}
		bus := newBus(pipe)
		bus.Subscribe(&TestSubscriber{})

		e := &TestEvent{}
		bus.Dispatch(e)

		if e.Emitted == false {
			t.Errorf("Event got not emitted")
		}
		if pipe.cnt != 1 {
			t.Errorf("Pipeline got not executed")
		}
	})

	t.Run("Check if per dispatch pipeline get called", func(t *testing.T) {
		pipe := &TestPipeline{}
		bus := newBus()
		bus.Subscribe(&TestSubscriber{})

		e := &TestEvent{}
		bus.PipeThrough(pipe).Dispatch(e)

		if e.Emitted == false {
			t.Errorf("Event got not emitted")
		}
		if pipe.cnt != 1 {
			t.Errorf("Pipeline got not executed")
		}
	})

	t.Run("Check if no global pipelines get executed if pipeOnlyThrough was called", func(t *testing.T) {
		pipe := &TestPipeline{}
		bus := newBus(pipe)
		bus.Subscribe(&TestSubscriber{})
		handled := false

		e := &TestEvent{}
		bus.PipeOnlyThrough(handler.PipelineFunc(func(next handler.Subscriber) handler.Subscriber {
			return handler.SubscriberFunc(func(object handler.EmitObject) {
				next.Handle(object)
				handled = true
			})
		})).Dispatch(e)

		if e.Emitted == false {
			t.Errorf("Event got not emitted")
		}
		if pipe.cnt != 0 {
			t.Errorf("Pipeline got executed")
		}
		if handled == false {
			t.Errorf("Pipeline callback got not executed")
		}
	})
}

func TestDispatcher(t *testing.T) {
	b := NewBus()
	o := &TestEventChannel{ch: make(chan int)}
	b.Subscribe(&TestEventChannelSubscriber{})

	b.Dispatch(o)

	select {
	case ret := <-o.ch:
		if ret != 1 {
			t.Errorf("wrong value returned")
		}
		return
	case <-time.After(1 * time.Second):
		t.Errorf("nothing received")
	}
}

func newBus(pipelines ...handler.Pipeline) *Bus {
	b := NewBus(pipelines...)

	return b.UseDispatcher(dispatcher.Simple)
}

type TestPipeline struct {
	cnt int
}

func (t *TestPipeline) Pipeline(next handler.Subscriber) handler.Subscriber {
	return handler.SubscriberFunc(func(object handler.EmitObject) {
		t.cnt++
		next.Handle(object)
	})
}

type NoOneSubscribedToItEvent struct {
	Emitted bool
}

func (n *NoOneSubscribedToItEvent) Time() time.Time {
	return time.Time{}
}

type TestEvent struct {
	Emitted bool
}

func (t *TestEvent) Time() time.Time {
	return time.Time{}
}

type TestSubscriber struct {
}

func (t *TestSubscriber) Handle(object handler.EmitObject) {
	object.(*TestEvent).Emitted = true
}

func (t *TestSubscriber) IsSubscribedTo(object handler.EmitObject) bool {
	_, ok := object.(*TestEvent)

	return ok
}

type CatchAllSubscriber struct {
	Handled []handler.EmitObject
}

func (c *CatchAllSubscriber) Handle(object handler.EmitObject) {
	c.Handled = append(c.Handled, object)
}

func (c *CatchAllSubscriber) IsSubscribedTo(object handler.EmitObject) bool {
	return true
}

type TestEventChannel struct {
	ch chan int
}

func (t *TestEventChannel) Time() time.Time {
	return time.Time{}
}

type TestEventChannelSubscriber struct {
}

func (t *TestEventChannelSubscriber) Handle(object handler.EmitObject) {
	o := object.(*TestEventChannel)

	o.ch <- 1
}

func (t *TestEventChannelSubscriber) IsSubscribedTo(object handler.EmitObject) bool {
	_, ok := object.(*TestEventChannel)

	return ok
}
