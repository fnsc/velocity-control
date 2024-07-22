package handlers

import "github.com/fnsc/velocity-control/domain"

type LoadHandler interface {
	SetNext(handler LoadHandler)
	Handle(request domain.Request) domain.Response
}

type BaseHandler struct {
	next LoadHandler
}

func (h *BaseHandler) SetNext(handler LoadHandler) {
	h.next = handler
}

func (h *BaseHandler) Handle(request domain.Request) domain.Response {
	if h.next != nil {
		return h.next.Handle(request)
	}

	return domain.Response{
		ID:         request.ID,
		CustomerID: request.CustomerID,
		Accepted:   true,
	}
}
