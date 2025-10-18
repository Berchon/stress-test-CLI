# 🔁 Stress Test CLI — Testes de Carga HTTP (Goroutines + Relatório)

[![Go](https://img.shields.io/badge/Go-1.24-blue)](https://golang.org/)  
[![Docker](https://img.shields.io/badge/Docker-Enabled-blue)](https://www.docker.com/)  

Sistema CLI desenvolvido em Go, com a funcionalidade principal de executar testes de carga em serviços web. O sistema realiza múltiplas requisições HTTP concorrentes para uma URL definida pelo usuário e gera um relatório detalhado de resultados.



## 📌 Objetivo do Projeto

Criar uma aplicação CLI que permita:
- Realizar um número configurável de requisições HTTP para uma URL alvo;
- Definir o nível de concorrência para execução das requisições;
- Gerar relatórios detalhados de sucesso, falhas, latência e percentis;
- Permitir execução via Go ou Docker.

## 🧾 Requisitos do Projeto

**Objetivo**: Criar uma ferramenta CLI em Go para testar a performance de serviços web por meio de requisições HTTP simultâneas.

#### Você deverá desenvolver:

- Função para receber os parâmetros via CLI (`--url`, `--requests`, `--concurrency`);
- Distribuição das requisições de acordo com o nível de concorrência;
- Registro de respostas HTTP, incluindo status e latência;
- Geração de relatório detalhado ao final da execução, incluindo:
  - Total de requests realizados;
  - Requests com HTTP 200;
  - Distribuição de outros códigos de status (400, 500, Others, Errors);
  - Estatísticas de latência (média, máxima, mínima, desvio padrão);
  - Percentis (P50, P90, P95, P99);
  - Tempo total de execução.

#### Entrega:

- Código-fonte completo da aplicação;
- Documentação explicando como rodar via Go e Docker;
- Dockerfile funcional para execução em container.

## 🧩 O que foi implementado (resumo)

- CLI para receber URL, total de requests e concorrência (`cmd/cli/main.go`);
- Orquestrador do load test em `internal/business/usecase/load_test.go`;
- Serviço HTTP (`internal/business/service/http_service.go`) para enviar requisições e medir latência;
- Entidade e DTO para armazenar resultados (`internal/business/entity` e `internal/business/dto`);
- Relatório final formatado conforme o padrão abaixo:

```bash
Executing stress test for URL: https://www.google.com . . .

 Stress Test Report
--------------------
 Target URL:           https://www.google.com
 Total Duration:       45.225s
 Total Requests:       600

Successful Requests:  83 (HTTP 200)

Failed Requests:
  Status Code |  Count
 ------------------------
  400         |        0
  500         |        0
  Others      |        0
  Errors      |      517

Latency (ms):
  Status Code |   Avg    |   Max    |   Min    |  StdDev
 ---------------------------------------------------------
  200         |    809.9 |   3843.0 |    187.0 |    978.2
  400         |      0.0 |      0.0 |      0.0 |      0.0
  500         |      0.0 |      0.0 |      0.0 |      0.0
  Others      |      0.0 |      0.0 |      0.0 |      0.0
  Total       |   1842.4 |   4059.0 |     79.0 |    966.2
  Errors      |   2008.2 |   4059.0 |     79.0 |    855.9

Percentiles (ms):
  Status Code |   P50    |   P90    |   P95    |   P99
 ---------------------------------------------------------
  200         |    374.0 |   2193.0 |   3385.0 |   3843.0
  400         |      0.0 |      0.0 |      0.0 |      0.0
  500         |      0.0 |      0.0 |      0.0 |      0.0
  Others      |      0.0 |      0.0 |      0.0 |      0.0
  Total       |   1767.0 |   3243.0 |   3374.0 |   3802.0
  Errors      |   1796.0 |   3304.0 |   3374.0 |   3766.0

Report generated at:  Sat, 18 Oct 2025 04:57:16 -03
```

## ⚙️ Parâmetros da CLI

| Parâmetro       | Descrição                      | Obrigatório / Exemplo                  |
| --------------- | ------------------------------ | -------------------------------------- |
| `--url`         | URL do serviço a ser testado   | Obrigatório — `https://www.google.com` |
| `--requests`    | Número total de requests       | Obrigatório — `600`                    |
| `--concurrency` | Número de chamadas simultâneas | Obrigatório — `30`                     |

## ▶️ Como Rodar o Projeto

### 💾 Clonar repositório

```bash
git clone git@github.com:Berchon/stress-test-CLI.git
cd stress-test-cli
```

### <img src="assets/go_logo.svg" height="16px"> Rodar via Go

Para rodar o projeto diretamente no seu ambiente de desenvolvimento com Go, use o seguinte comando:
```bash
go run cmd/cli/main.go --url=https://www.google.com --requests=600 --concurrency=30
```
>Basta substituir a `URL`, o número de requisições (`--requests`) e a concorrência (`--concurrency`) pelos parâmetros desejados.

### 🐳 Rodar via Docker

Se preferir rodar o projeto dentro de um container Docker, siga os passos abaixo:

**1. Build da imagem Docker:**
```bash
# Build da imagem
docker build -t stress-test-cli .
```

**2. Executar o container:**
```bash
# Executar o container
docker run stress-test-cli --url=https://www.google.com --requests=600 --concurrency=30
```
Esse comando executa o teste de carga diretamente no Docker, utilizando os parâmetros passados via CLI.

### 🔗 Testando com httpbin.org
**httpbin.org** é uma ferramenta útil para simulação de APIs e pode ser usada para testar latência, timeout e status codes de maneira controlada. Abaixo estão alguns exemplos para você testar diferentes cenários:

* **Simular latência (tempo de resposta configurado):**
```bash
go run cmd/cli/main.go --url=https://httpbin.org/delay/1 --requests=20 --concurrency=5
```
A URL `https://httpbin.org/delay/1` retorna um **HTTP 200** após um atraso de 1 segundos, permitindo testar a latência das requisições.

* **Simular timeout (requisição que excede o tempo de espera):**
```bash
go run cmd/cli/main.go --url=https://httpbin.org/delay/6 --requests=20 --concurrency=5
```

Aqui, a URL `https://httpbin.org/delay/6` faz com que a resposta demore 6 segundos, forçando `timeout` nas requisições (já configurado para 5 segundos no cliente HTTP).

* **Testar diferentes status codes aleatórios:**
```bash
go run cmd/cli/main.go --url=https://httpbin.org/status/200,300,400,500 --requests=100 --concurrency=5
```

A URL `https://httpbin.org/status/200,300,400,500` retorna status codes aleatórios entre **200**, **300**, **400**, e **500**, o que é útil para testar como o sistema lida com diferentes respostas HTTP.

Esses exemplos são extremamente úteis para testar como a aplicação reage a diferentes cenários de resposta da API e comportamento de rede.

## 🧰 Tecnologias utilizadas

- Golang 1.24
- Docker
- Goroutines para execução concorrente de requests

## 📂 Estrutura Completa do Projeto
```bash
stress-test-cli/
├── .git/                          # histórico git (repositório presente no ZIP)
├── .gitignore                     # padrões de arquivos ignorados
├── Dockerfile                     # Dockerfile existente (ver abaixo para sugestão)
├── README.md                      # README presente (resumo do projeto)
├── go.mod                         # módulo go
├── go.sum                         # somas de dependência
├── cmd/
│   └── cli/
│       └── main.go                # Entrypoint CLI: inicializa a aplicação e parseia flags
├── internal/
│   ├── business/
│   │   ├── dto/
│   │   │   ├── report.go          # DTO: estrutura do relatório final (P50, P90, etc)
│   │   │   ├── request_result.go  # DTO: estrutura que representa resultado de uma request
│   │   │   └── stress_test_params.go # DTO: parâmetros do teste (url, requests, concurrency)
│   │   ├── entity/
│   │   │   └── request_result.go   # Entidade: armazenamento/representação de cada request (timestamps, status)
│   │   ├── service/
│   │   │   └── http_client_service.go # Serviço que dispara requests HTTP e mede latência
│   │   └── usecase/
│   │       └── load_stress-test.go  # Orquestrador do load test: cria goroutines, agrega resultados
│   │
│   └── infra/
│       ├── cli/
│       │   ├── cli.go              # Código para inicializar CLI (comandos/execução)
│       │   ├── flags.go            # Definição de flags/argumentos CLI
│       │   ├── help.go             # Funções de ajuda/usage
│       │   ├── help_text.go        # Texto de help estático
│       │   └── root.go             # Root command / montagem do CLI (se estiver usando cobra-like)
│       ├── report/
│       │   └── console_report.go   # Gera relatório (impressão no console)
│       └── statistics/
│           ├── aggregate.go       # Funções de agregação de latências/sucessos/erros
│           ├── compute.go         # Cálculos (percentis, média, desvios)
│           └── statistics.go      # Tipos/coleções para estatísticas

```

👨‍💻 Autor

Projeto desenvolvido por Berchon — https://github.com/Berchon