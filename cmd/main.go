package main

import (
	"flag"
	"log"
	"time"

	"github.com/eberte-sampaio/cli-health-check/internal/infra"
	"github.com/eberte-sampaio/cli-health-check/internal/service"
	"github.com/eberte-sampaio/cli-health-check/internal/usecase"
)

func main() {
	start := time.Now()
	filePath := flag.String("file", "sites.csv", "caminho do arquivo CSV")
	flag.Parse()

	if *filePath == "" {
		log.Fatalf("A flag -file é obrigatória")
	}

	loader := infra.NewCsvLoader()
	urls, err := loader.Load(*filePath)
	if err != nil {
		log.Fatalf("Falha crítica ao carregar o arquivo: %v", err)
	}

	healthCheckService := service.NewIHeathCheckService()

	batchProcessor := usecase.NewCheckBathUseCase(healthCheckService)

	log.Printf("Iniciando a verificação de %d sites", len(urls))
	batchProcessor.Run(urls)

	elapse := time.Since(start)
	log.Printf("Processo finalizado em %s", elapse)
}
