package handler

type Pipeline interface {
	Pipeline(handler Subscriber) Subscriber
}

type PipelineFunc func(handler Subscriber) Subscriber

func (s PipelineFunc) Pipeline(next Subscriber) Subscriber {
	return s(next)
}
