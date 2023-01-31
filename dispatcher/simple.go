package dispatcher

import (
	"github.com/evolidev/bus/handler"
)

var Simple = DispatcherFunc(func(subscriber handler.Subscriber, object handler.EmitObject) {
	subscriber.Handle(object)
})
