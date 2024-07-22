package handlers

import "github.com/fnsc/velocity-control/domain"

type DailyLimitHandler struct {
	BaseHandler
	dailyLoads map[int64]map[string]float64
}

func NewDailyLimitHandler() *DailyLimitHandler {
	return &DailyLimitHandler{dailyLoads: make(map[int64]map[string]float64)}
}

func (h *DailyLimitHandler) Handle(request domain.Request) domain.Response {
	dateKey := request.Time.Format("2006-01-02")
	_, exists := h.dailyLoads[request.CustomerID]
	if !exists {
		h.dailyLoads[request.CustomerID] = make(map[string]float64)
	}

	limits := domain.NewLimits()

	h.dailyLoads[request.CustomerID][dateKey] += request.LoadAmount
	if h.dailyLoads[request.CustomerID][dateKey] > limits.Daily {
		return domain.Response{
			ID:         request.ID,
			CustomerID: request.CustomerID,
			Accepted:   false,
		}
	}

	return h.BaseHandler.Handle(request)
}
