# 🎟️ Sistema de Bilhetagem – Case Técnico

## 🔍 Visão Geral

Esta solução foi projetada com foco em **escalabilidade**, **modularidade** e **organização de código**, aproveitando os princípios de separação de responsabilidades e orientação a domínios.

A aplicação foi desenvolvida em **Go**, utilizando uma única base de código (`monorepo`) organizada em pacotes isolados que simulam uma arquitetura de microserviços. Essa estrutura permite que, futuramente, os pacotes possam ser extraídos como serviços independentes com **mínimo acoplamento**.

---

## 🧱 Estrutura Modular por Domínio

Cada domínio principal do sistema está representado por um pacote Go independente, contendo:

- Definição das **entidades**
- Implementação dos **casos de uso (use cases)**
- **Validações** específicas do domínio

### 📦 Pacotes existentes:

- `produtos`
- `usuarios`
- `pagamentos`
- `pedidos`
- `sessoes`
- `ingressos`
- `filmes`

---

## 🔗 Camadas Compartilhadas

Algumas camadas são compartilhadas entre os domínios para promover reaproveitamento de código e consistência:

### 🔄 Repositories

- Interfaces que definem contratos para acesso a dados (PostgreSQL, Redis, etc).
- Implementações concretas reutilizáveis por múltiplos domínios.

### 🔌 Adapters

- Cuidam da integração com recursos externos ou infraestrutura.
- Exemplos:
  - Adaptador SQL para PostgreSQL/SQLite
  - Adaptador Redis para cache e controle de assentos

---

## 🌐 Interface de Entrada

A camada de **entrega** (delivery) é composta por handlers HTTP responsáveis por:

- Mapeamento das rotas REST da API
- Autenticação e autorização (quando necessário)
- Validação e parsing de entrada
- Encaminhamento das chamadas para os casos de uso

---

## ✅ Benefícios da Arquitetura

- **Escalável**: fácil evolução para uma arquitetura baseada em microserviços reais.
- **Organizada**: estrutura clara e separada por contexto de negócio.
- **Testável**: pacotes isolados com lógica de domínio facilmente testável.
- **Desacoplada**: baixa dependência entre módulos.

---

## 🧭 Diagramas

### 🗂️ Arquitetura da Solução

> Representação de alto nível da estrutura de serviços e integrações.

![Arquitetura da Solução](Docs/imgs/Arquitetura%20Case.png)

---

### 🧩 Modelagem de Domínio

> Diagrama ER com relacionamentos entre entidades do sistema.

![Modelagem de Domínio](Docs/imgs/Modelagem%20Case%20Itau%20-%20PH.png)