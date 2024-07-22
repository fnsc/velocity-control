package handlers

import "github.com/fnsc/velocity-control/domain"

type DailyLoadCountHandler struct {
	BaseHandler
	dailyCounts map[int64]map[string]int
}

func NewDailyLoadCountHandler() *DailyLoadCountHandler {
	return &DailyLoadCountHandler{dailyCounts: make(map[int64]map[string]int)}
}

func (h *DailyLoadCountHandler) Handle(request domain.Request) domain.Response {
	dateKey := request.Time.Format("2006-01-02")
	_, exists := h.dailyCounts[request.CustomerID]
	if !exists {
		h.dailyCounts[request.CustomerID] = make(map[string]int)
	}

	h.dailyCounts[request.CustomerID][dateKey]++
	limits := domain.NewLimits()
	if h.dailyCounts[request.CustomerID][dateKey] > limits.DailyLoad {
		return domain.Response{
			ID:         request.ID,
			CustomerID: request.CustomerID,
			Accepted:   false,
		}
	}

	return h.BaseHandler.Handle(request)
}
