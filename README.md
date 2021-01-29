# Routes API

### O que é?

É uma API para cadastrar rotas comerciais, clientes e vendedores. Os vendedores são associados a uma rota e nessa rota possui os clientes que o vendedor entra em contato.

Essa API foi implementada usando a linguagem Go, usando Docker e Postgresql (Postgis).

### Arquitetura

A ideia foi seguir algo na linha da arquitetura hexagonal e arquitetura limpa.

![hexagonal](https://apiumhub.com/wp-content/uploads/2018/10/Screenshot-2018-10-30-at-08.45.49.png "Hexagonal")

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
