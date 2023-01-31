package handler

type Subscriber interface {
	Handle(object EmitObject)
	IsSubscribedTo(object EmitObject) bool
}

type SubscriberFunc func(object EmitObject)

func (s SubscriberFunc) Handle(object EmitObject) {
	s(object)
}

func (s SubscriberFunc) IsSubscribedTo(object EmitObject) bool {
	return true
}
