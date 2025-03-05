# Go RabbitMQ Consumer

Este projeto é um consumer de mensagens do RabbitMQ escrito em Go. As mensagens recebidas são armazenadas em um banco de dados MongoDB. O serviço é executado como um processo independente e espera que o RabbitMQ já esteja rodando.

## Requisitos

- Go 1.20
- Docker e Docker Compose
- RabbitMQ rodando previamente

## Configuração

Antes de executar a aplicação, é necessário configurar as variáveis de ambiente. Um exemplo de configuração pode ser encontrado no arquivo `.env.example`.

Crie um arquivo `.env` baseado no `.env.example` e edite conforme necessário:

```sh
cp .env.example .env
```

## Como executar

1. Suba o MongoDB utilizando o Docker Compose:

   ```sh
   docker-compose up -d
   ```

2. Certifique-se de que o RabbitMQ já está rodando.
3. Execute a aplicação Go:

   ```sh
   go run main.go
   ```

## Considerações

- A aplicação **não inicia o RabbitMQ**, ele deve estar rodando previamente.
- O MongoDB é iniciado via Docker Compose.
- Utilize o `.env` para definir as credenciais e conexões necessárias.

## Licença

Este projeto é distribuído sob a licença MIT.

