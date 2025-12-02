package service

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/eberte-sampaio/cli-health-check/internal/domain"
)

type IHeathCheckService struct {
}

func NewIHeathCheckService() domain.HealthCheckService {
	return &IHeathCheckService{}
}
func (s *IHeathCheckService) GetWebsiteStatus(url string) (*domain.HealthCheck, error) {

	client := http.Client{
		Timeout: 10 * time.Second,
	}
	response, err := client.Get(url)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Erro ao fazer requisição para a url %s", url))
	}

	defer response.Body.Close()

	healthCheck := &domain.HealthCheck{
		Site:      url,
		IsActive:  response.StatusCode < 400,
		Status:    response.StatusCode,
		Timestamp: time.Now(),
	}
	return healthCheck, nil
}
