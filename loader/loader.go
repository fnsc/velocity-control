package loader

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/fnsc/velocity-control/domain"
)

func ParseRequest(transaction string) (domain.Request, error) {
	var request domain.Request
	var tempRequest domain.LoadRequest

	if err := json.Unmarshal([]byte(transaction), &tempRequest); err != nil {
		return request, err
	}

	id, err := strconv.ParseInt(tempRequest.ID, 0, 64)
	if err != nil {
		return request, err
	}

	customerId, err := strconv.ParseInt(tempRequest.CustomerID, 0, 64)
	if err != nil {
		return request, err
	}

	stringAmount := strings.TrimPrefix(tempRequest.LoadAmount, "$")
	amount, err := strconv.ParseFloat(stringAmount, 64)
	if err != nil {
		return request, err
	}

	timestamp, err := time.Parse(time.RFC3339, tempRequest.Time)
	if err != nil {
		return request, err
	}

	request.ID = id
	request.CustomerID = customerId
	request.LoadAmount = amount
	request.Time = timestamp

	return request, nil
}
