package usecase

import (
	"log"
	"sync"

	"github.com/eberte-sampaio/cli-health-check/internal/domain"
)

type CheckBathUseCase struct {
	service domain.HealthCheckService
}

func NewCheckBathUseCase(service domain.HealthCheckService) *CheckBathUseCase {
	return &CheckBathUseCase{service: service}
}

func (u *CheckBathUseCase) Run(urls []string) {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)

		go func(websiteUrl string) {
			defer wg.Done()

			healthCheck, err := u.service.GetWebsiteStatus(websiteUrl)

			if err != nil {
				log.Printf("[ERROR] %s : %v", websiteUrl, err)
				return
			}

			log.Printf(healthCheck.GetLog())
		}(url)
	}
	wg.Wait()
}
