## Sistema de Pedidos com Comunicação Assíncrona

Este é um sistema de consumo de pedidos que utiliza microserviços para o processamento de pedidos com comunicação assíncrona utilizando RabbitMQ.

## Tecnologias Utilizadas

- Go
- RabbitMQ
- PostgreSQL
- Docker

## Estrutura do Projeto

- `config/`: Configurações e inicializações.
- `consumer/`: Consumo e processamento de mensagens RabbitMQ.
- `handler/`: Manipuladores de endpoints.
- `models/`: Estruturas e funções de banco de dados.
- `router/`: Definição de rotas.

## Pré-requisitos

- Docker e Docker Compose
- Go
- PostgreSQL

## Configuração do Ambiente

1. **Clone o repositório**:
    ```sh
    git clone https://github.com/enyasantos/go-async-order-system.git
    cd go-async-order-system
    ```

2. **Crie e configure o arquivo `.env`**:
    ```sh
    cp .env.example .env
    ```

3. **Inicie os containers Docker**:
    ```sh
    docker-compose up -d
    ```

4. **Instale as dependências do Go**:
    ```sh
    go mod tidy
    ```

## Inicialização do Projeto

1. **Compile e inicie o projeto**:
    ```sh
    go run main.go
    ```

2. **Servidor HTTP** estará disponível em `http://localhost:8080`.

## Endpoints Disponíveis

- `GET /orders/{orderId}`: Retorna o preço total do pedido especificado.
- `GET /orders`: Lista todos os pedidos.

## Licença

Este projeto é licenciado sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.

