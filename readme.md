# Go Rate Limiter Challenge

Este projeto implementa um rate limiter em Go que pode ser configurado para limitar o número máximo de requisições por segundo com base em um endereço IP específico ou em um token de acesso. A aplicação utiliza Redis para armazenar as informações de limitação e permite configuração via variáveis de ambiente ou arquivo .env.

## Estrutura do projeto

```
├── cmd
│   └── server
│       └── main.go
├── configs
│   └── config.go
├── internal
│   └── middleware
│       └── ratelimiter.go
├── pkg
│   └── ratelimiter
│       ├── limiter.go
│       ├── redis.go
│       └── strategy.go
├── test
│   └── ratelimiter_test.go
├── Taskfile.yml
├── docker-compose.yml
├── Dockerfile
├── .air.toml
├── .env
└── README.md
```

### Funcionalidades

- Limitação de requisições por IP e token de acesso.
- Configuração do número máximo de requisições por segundo.
- Opção para configurar o tempo de bloqueio após exceder o limite.
- Configuração via variáveis de ambiente ou arquivo `.env`.
- Armazenamento das informações de limitação no Redis.
- Middleware para fácil integração com servidores web.
- Suporte para troca de mecanismo de persistência.

### Requisitos

- Docker
- Docker Compose
- Go 1.21.3
- Taskfile (opcional)
- Air (opcional)

### Configuração do Ambiente

1. **Clone o repositório**

   ```sh
   git clone https://github.com/rzeradev/rate-limiter.git
   cd rate-limiter
   ```

2. **Crie e configure o aquivo .env**

   ```sh
   cp .env.example .env
   # Modifique o arquivo .env como necessário
   ```

3. **Inicie o servidor Redis com Docker Compose**

   ```sh
   docker-compose up -d
   ```

4. **Execute o projeto**

   ```sh
   task
   ```

   ou

   ```sh
   air
   ```

   ou

   ```sh
    go run ./cmd/server/main.go
   ```

5. **O servidor vai rodar no endereço `http://localhost:8080`**

### Rodando os testes

1. **Execute os testes**
   ```sh
   task test
   ```
   ou
   ```sh
   go test -v ./test
   ```

## Configuração

A configuração necessária pode ser feita no arquivo `.env`:

```
REDIS_ADDR=localhost:6379
REDIS_PASS=
REDIS_DB=0
IP_MAX_REQ=5
TOKEN_MAX_REQ=10
BLOCK_TIME=60 #in seconds
RATE_LIMIT_DUR=second #can be per "second" or per "minute"
```

## Padrão Strategy

O rate limiter utiliza o padrão de projeto Strategy, permitindo a fácil troca entre diferentes mecanismos de persistência. Atualmente, ele usa Redis, mas pode ser estendido para usar outros armazenamentos, como em memória, sql, nosql ou outro.

## Middleware

O rate limiter é implementado como um middleware que pode ser injetado no framework Fiber (ou qualquer outro framework) para controlar a taxa de requisições recebidas com base no endereço IP ou token de acesso.

### Conclusão

Este projeto fornece uma implementação completa de um rate limiter em Go com suporte a Redis e middleware para fácil integração com servidores web. A configuração é simples e flexível, permitindo ajustes via variáveis de ambiente ou arquivos `.env`. Os testes unitários garantem a robustez e a eficácia do rate limiter.

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).
