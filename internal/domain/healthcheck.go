package domain

import (
	"fmt"
	"time"
)

type HealthCheck struct {
	Site      string
	IsActive  bool
	Status    int
	Timestamp time.Time
}

type HealthCheckService interface {
	GetWebsiteStatus(url string) (*HealthCheck, error)
}

func (h *HealthCheck) GetLog() string {

	return fmt.Sprintf("[%s] %s - UP: %t (%d) ", h.Timestamp, h.Site, h.IsActive, h.Status)
}
