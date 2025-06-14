## Visão Geral

A solução é estruturada com foco em escalabilidade e organização modular de código, aproveitando as boas práticas de separação de responsabilidades. A linguagem utilizada é Go, e mesmo com uma única base de código, os serviços são organizados de forma isolada em pacotes, simulando uma arquitetura de microserviços dentro de um monorepo.

Essa organização permite que, no futuro, cada pacote possa ser extraído e convertido em um serviço independente com mínimo acoplamento.

---

## Estrutura de Pacotes por Domínio

Cada entidade principal do sistema está representada por um pacote Go independente, contendo suas regras de domínio e casos de uso (use cases). Os pacotes definidos no projeto são:

- **Produtos**
- **Usuário**
- **Pagamento**
- **Pedidos**
- **Reservas**
- **Sessões**
- **Filmes**

Cada pacote é responsável por:
- Definir suas **entidades** e regras de negócio.
- Implementar seus **casos de uso**.
- Realizar **validações** específicas do domínio.
---
## Compartilhamento de Camadas

Apesar da separação por domínio, as seguintes camadas são compartilhadas entre os pacotes:

### Repositories

- Interfaces que definem os contratos de acesso a dados (PostgreSQL, Redis, etc).
- Implementações concretas compartilhadas entre os serviços, quando aplicável.

### Adapters

- Implementações de integração externa e persistência.
- Exemplos:
    - Adaptador SQL (PostgreSQL)
    - Adaptador Redis (cache e reserva de assentos)

---
## Interface de Entrada

A camada de entrega (delivery) é composta por handlers HTTP que expõem os casos de uso. Essa camada é responsável por:

- Mapear as rotas da API REST
- Realizar a autenticação/autorizacção (se aplicável)
- Realizar parse/validação dos dados de entrada
- Encaminhar para o caso de uso correspondente

---

## Benefícios dessa Estrutura

- **Escalabilidade**: facilita transformação de pacotes em microserviços reais no futuro.
- **Organização**: separa responsabilidades de forma clara e intuitiva.
- **Testabilidade**: cada pacote pode ser testado isoladamente.
- **Baixo acoplamento**: cada módulo tem dependências mínimas entre si.