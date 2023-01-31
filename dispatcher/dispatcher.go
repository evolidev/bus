package dispatcher

import "github.com/evolidev/bus/handler"

type Dispatcher interface {
	Dispatch(subscriber handler.Subscriber, object handler.EmitObject)
}

type DispatcherFunc func(subscriber handler.Subscriber, object handler.EmitObject)

func (df DispatcherFunc) Dispatch(subscriber handler.Subscriber, object handler.EmitObject) {
	df(subscriber, object)
}
