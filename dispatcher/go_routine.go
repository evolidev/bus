package dispatcher

import (
	"github.com/evolidev/bus/handler"
)

var GoRoutine = DispatcherFunc(func(subscriber handler.Subscriber, object handler.EmitObject) {
	go func(sub handler.Subscriber, ob handler.EmitObject) {
		sub.Handle(ob)
	}(subscriber, object)
})

//type GoRoutine struct {
//}
//
//func (g *GoRoutine) Dispatch(subscriber bus.Subscriber, object bus.EmitObject) {
//
//}
