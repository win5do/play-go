package main

import "context"

type IMiddleware interface {
	Handle(ctx context.Context, in interface{}) (interface{}, error)
}

type Chain []IMiddleware

type engine struct {
	Handlers Chain
}

func (e *engine) Use(h IMiddleware) {
	e.Handlers = append(e.Handlers, h)
}

func (e *engine) Handle(ctx context.Context, in interface{}) (interface{}, error) {
	var r interface{}
	var err error
	for i := range e.Handlers {
		r, err = e.Handlers[i].Handle(ctx, in)
		if err != nil {
			return nil, err
		}
		in = r
	}

	return r, err
}
