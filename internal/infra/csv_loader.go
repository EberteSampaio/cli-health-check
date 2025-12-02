package infra

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CsvLoader struct {
}

func NewCsvLoader() *CsvLoader {
	return &CsvLoader{}
}

func (l *CsvLoader) Load(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir o arquivo  %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("erro ao ler CSV: %w", err)
	}

	var urls []string
	for _, record := range records {
		if len(record) > 0 {
			urls = append(urls, record[0])
		}
	}

	return urls, nil
}
