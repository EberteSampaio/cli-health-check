# GO CLI Health Checker

CLI Health Checker é uma ferramenta de linha de comando (CLI) de focado em alta performace contruída em **GO** para verificar a disponibilidade de listas massivas de sites simultaneamente

Este projeto demonstra o uso de **Goroutines**, **Channels** (implícito via WaitGroup) e **Clean Architecture** para criar um verificador de status HTTP rápido e escalável.

---

## Arquitetura

O projeto segue os princípios da **Clean Architecture** para garantir desacoplamento e testabilidade. 

A estrutura de pastas reflete a separação de responsabilidades:

```text
├── cmd/
│   └── cli/
│       └── main.go
├── internal/
│   ├── domain/            # Entidades e Interfaces
│   ├── infra/             # Implementações externas
│   ├── service/           # Lógica de negócio HTTP
│   └── usecase/           # Lógica de Aplicação
└── sites.csv              # Arquivo de entrada
```

## Camadas

- **Domain**: Define o que é um HealthCheck e as interfaces do sistema.
- **Infra**: Responsável por ler dados do disco (CSV).
- **Service**: Realiza requisições HTTP e validar status codes (200 vs 400+).
- **UseCase**: Gerencia a concorrência (sync.WaitGroup), disparando múltiplas threads leves (Goroutines) para processar o lote.

## Funcionalidades 

- **Alta Concorrência**: Verifica centenas de sites em segundos usando Goroutines.
- **Segurança de Thread**: Uso correto de sync.WaitGroup para sincronização.
- **Configurável**: Aceita qualquer arquivo CSV via flag -file.
- **Clean Code**: Código modular, facilitando a troca de leitura de CSV para Banco de Dados no futuro sem quebrar a lógica de negócio.

## Tecnologias 
- **Linguagem:** Go (Golang) 1.24+
- **Bibliotecas:** Standard Library (net/http, sync, encoding/csv, flag).
- Docker
## Pré-requisitos

- Tenha o [Go](https://go.dev/doc/install) instalado na sua máquina.
- Tenha o [Docker](https://docs.docker.com/) instalado na sua máquina.
## 1. Clone o repositório

```bash
  git clone https://github.com/EberteSampaio/cli-health-check.git
  cd cli-health-check
```

## 2. Prepare o arquivo de entrada

Substituia o arquivo chamado sites.csv na raiz do projeto (ou use o exemplo fornecido). O formato deve ser apenas a URL na primeira coluna:

```text
https://google.com
https://github.com
https://stackoverflow.com
https://site-inexistente-teste.com
```

## 3. Rode a imagem docker

```bash
  docker build -t cli-healthcheck .
```
## 4. Execute
Rode o projeto apontando para o arquivo: 

```bash
  docker run --rm -v caminho-omitido/sites.csv:/app/sites.csv cli-healthcheck -file sites.csv
```

## Exemplo de Saída

```text
2025/12/02 11:40:14 Iniciando a verificação de 10 sites
2025/12/02 11:40:15 [2025-12-02 11:40:15.100222196 -0300 -03 m=+0.215042026] https://www.youtube.com - UP: true (200) 
2025/12/02 11:40:15 [2025-12-02 11:40:15.109880968 -0300 -03 m=+0.224700784] https://www.amazon.com - UP: true (200) 
2025/12/02 11:40:15 [2025-12-02 11:40:15.182304102 -0300 -03 m=+0.297123953] https://www.instagram.com - UP: true (200) 
2025/12/02 11:40:15 [2025-12-02 11:40:15.620810471 -0300 -03 m=+0.735630294] https://www.twitter.com - UP: true (200) 
2025/12/02 11:40:15 [2025-12-02 11:40:15.751346744 -0300 -03 m=+0.866166556] https://www.linkedin.com - UP: true (200) 
2025/12/02 11:40:15 [2025-12-02 11:40:15.819338602 -0300 -03 m=+0.934158430] https://www.reddit.com - UP: true (200) 
2025/12/02 11:40:20 [2025-12-02 11:40:20.650569775 -0300 -03 m=+5.765389592] https://www.google.com - UP: true (200) 
2025/12/02 11:40:20 [2025-12-02 11:40:20.746444634 -0300 -03 m=+5.861264440] https://www.yahoo.com - UP: false (429) 
2025/12/02 11:40:20 [2025-12-02 11:40:20.75971408 -0300 -03 m=+5.874533887] https://www.facebook.com - UP: true (200) 
2025/12/02 11:40:24 [ERROR] https://www.wikipedia.org : Erro ao fazer requisição para a url https://www.wikipedia.org
2025/12/02 11:40:24 Processo finalizado em 10.003527331s
```
## Próximos Passo 

- [ ] Adicionar Testes Unitários para a camada de Service e Usecase.

- [ ] Implementar um timeout configurável via flag.

- [ ] Salvar o relatório de saída em um arquivo JSON.

- [x] Dockerizar a aplicação.

Desenvolvido por [Eberte Sampaio](https://github.com/EberteSampaio)
