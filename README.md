# Routes API

### O que é?

É uma API para cadastrar rotas comerciais, clientes e vendedores. Os vendedores são associados a uma rota e nessa rota possui os clientes que o vendedor entra em contato.

Essa API foi implementada usando a linguagem Go, usando Docker e Postgresql (Postgis).

### Arquitetura

A ideia foi seguir algo na linha da arquitetura hexagonal e arquitetura limpa.

![hexagonal](https://apiumhub.com/wp-content/uploads/2018/10/Screenshot-2018-10-30-at-08.45.49.png "Hexagonal")

## Modelagem do banco de dados

![db](doc/db.png "Modelagem")

**config**

- _db_: aqui ficam as migrações de banco de dados, resolvi colocar dentro do projeto pensando no caso de ter algum outro desenvolvedor é interessante que todos tenham o banco com as alterações
- http-routes é um arquivo onde é definido as rotas e qual controller vai ser usado para aquela requisição

**controllers**: aqui é onde ficam os controllers que vão lidar com as requisições

**entities**: as entidades da aplicação (administradores, vendedores, rotas e clientes)

**infra**: aqui fica toda a parte de serviços como banco de dados, JWT, toda parte que não é regra de negócio mas que a nossa aplicação precisa para funcionar

**middlewares**: o middleware de autenticação para verificar se é um administrador que está fazendo a requisição

**repositories**: o padão repositório usado para acessar banco de dados

**usecases**: as nossas regras de negócio ficam aqui em vez de ficar no controller ou na entity.

### Como rodar

1- Com o docker instalado pode usar o comando _docker-compose up_

2 - Caso queira criar uma migration pode usar uma cli chamada [golang-migrate](https://github.com/golang-migrate/migrate) e tem um Makefile basta rodar o command _make migrateup_ ou _make migratedown_

### Requisições

![login](doc/login.png "Login")

![auth_fails](doc/auth_fails.png "Malformed Token")

![invalid_token](doc/invalid_token.png "Invalid Token")

![create_seller](doc/create_seller.png "Create Seller")

![get_sellers](doc/get_sellers.png "Get Sellers")

![delete_seller](doc/delete_seller.png "Delete Sellers")
