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
Estudos e analises para definir arquitetura AWS:
    - Publico Alvo
    - Venda de ingressos vs Bomboniere
    - De acordo com o publico quais produtos poderiam ser listado primeiro?
    - Média de utilização (Pré e Pós Lancamentos)

#### Feature Bonus caso dê tempo
    - Serviço que classifica um lançamento e escala preditivamente a infra para atender a demanda (Metrica Simples Google trends ou IA)
    - Classificação de items após a reserva de ingressos
    - Login Oauth Google

### O que irei entregar ao Final deste case

MVP Funcional para apresentação Back e Front
Estudo sobre o Cinema
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

### Frontend

- Criar paleta de cores para o projeto
- Criar wireframe para o fluxos que o usuário irá fazer
- Criar design no figma (Opcional se der tempo)
- Criar criar infra do front, adicionar o build em rota do servidor para empacotar tudo junto
- Criar components reutilizáveis
- Desenvolver as telas

### Arquitetura AWS

- Analisar estudo
	- Publico Alvo
    - Venda de ingressos vs Bomboniere
    - De acordo com o publico quais produtos poderiam ser listado primeiro?
    - Média de utilização (Pré e Pós Lancamentos)
- Criar desenho a partir do dados
- Finalizar documentações
