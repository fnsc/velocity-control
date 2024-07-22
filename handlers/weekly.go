package handlers

import (
	"fmt"

	"github.com/fnsc/velocity-control/domain"
)

type WeeklyLimitHandler struct {
	BaseHandler
	weeklyLoads map[int64]map[string]float64
}

func NewWeeklyLimitHandler() *WeeklyLimitHandler {
	return &WeeklyLimitHandler{weeklyLoads: make(map[int64]map[string]float64)}
}

func (h *WeeklyLimitHandler) Handle(request domain.Request) domain.Response {
	year, week := request.Time.ISOWeek()
	weekKey := fmt.Sprintf("%d-W%d", year, week)
	_, exists := h.weeklyLoads[request.CustomerID]
	if !exists {
		h.weeklyLoads[request.CustomerID] = make(map[string]float64)
	}

	limits := domain.NewLimits()

	h.weeklyLoads[request.CustomerID][weekKey] += request.LoadAmount
	if h.weeklyLoads[request.CustomerID][weekKey] > limits.Weekly {
		return domain.Response{
			ID:         request.ID,
			CustomerID: request.CustomerID,
			Accepted:   false,
		}
	}

	return h.BaseHandler.Handle(request)
}
