# Desafio Tech Itaú

## Desafio

Você deverá propor uma arquitetura na AWS e desenvolver uma solução funcional (em sua stack de preferência) para um sistema de bilhetagem com as seguintes características:

- O sistema deve permitir solicitação, reserva e compra de ingressos.
- Durante o processo de compra, o sistema deve oferecer produtos adicionais como pipoca, chocolate, refrigerante, etc
- A solução deve conter uma única base de código (um único projeto/solution), mesmo que a arquitetura proposta seja orientada a microserviços. Isso facilitará a apresentação e a avaliação do seu trabalho.

## Refinamento do projeto

### Objetivo
Criar um serviço que permita o usuário realizar a solicitação, reserva e compra de ingressose também oferecer produtos adicionais

### Algumas decisões antes de iniciar o projeto

O Contexto do serviço: a solução será exclusiva para o contexto de cinema
Stacks de desenvolvimento: Go (Backend) - Angular (Frontend)
Escala da aplicação: Regionalizada focada em SP
Principio: O desenvolvimento do sistema será orientado a partir da definição da API, que funcionará como o contrato principal entre serviços, frontend, ou qualquer integração externa.

### O que irei entregar ao Final deste case

MVP Funcional para apresentação (Back)
Desenho de arquitetura do projeto para AWS


## Tarefas

### Backend

- Criar desenho da api e o OpenAPI
- Criar collection para testes
- Organizar as regras de dominío
- Criar os repositórios para os serviços
- Desenvolver a API
- Criar os testes
- Criar script de seed (para popular os dados de cinema, sessão e etc)
