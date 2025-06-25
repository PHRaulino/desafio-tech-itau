Apresentação - Case Técnico Sistema de Bilhetagem

1. Demonstração do Código (20 min)

Abertura

“Vou mostrar primeiro a aplicação funcionando e em seguida passo pelos pontos estruturais e técnicos do projeto.”

Demonstração funcional
	•	Usar Swagger (docs/index.html) ou scripts de teste (bin/curls.sh) ou o Insomnia
	•	Exemplo de fluxo:
	•	GET /filmes → GET /sessoes → POST /reserva → POST /pedido → POST /checkout → POST /pagamento → GET /usuario/pedidos

Explicação da Estrutura
	•	Projeto monolítico, mas com separação por domínio em pkgs/
	•	Cada domínio tem:
	•	core: modelos + interfaces
	•	usecases: regras de negócio
	•	adapters: infraestrutura (SQLite)
	•	handlers: camada HTTP com injeção via Wire
	•	errors: tratamento de erros contextualizados

Destaques técnicos
	•	sqlc: acesso seguro ao banco via queries tipadas
	•	wire: injeção de dependência automática
	•	JWT com middleware (middlewares/auth.go)
	•	SQLite local, mas estrutura portável para PostgreSQL
	•	OpenAPI documentado com exemplos realistas
	•	Testes realizados com Insomnia

Sobre testes e logs

“A arquitetura está preparada para testes unitários e logs estruturados, mas foquei em entregar o core funcional dentro do prazo. A estrutura modular permite evoluir facilmente com testes isolados por usecase.”

2. Discussão da Arquitetura AWS (20 min)

Visão Geral

“Essa arquitetura foi pensada para ser escalável, segura e orientada a domínio. O foco foi na orquestração de serviços em Fargate e desacoplamento via filas.”

Componentes

Entrada
	•	SPA Angular (S3 + CloudFront)
	•	API Gateway com JWT ou Cognito

Execução e Microserviços
	•	ECS Fargate com auto scaling + Load Balancer
	•	Microserviços separados por domínio: filmes, sessões, pedidos, etc.
	•	Cada serviço comunica-se via HTTP ou fila SQS

Persistência
	•	PostgreSQL com réplicas de leitura
	•	Redis para TTL e locking (ex: assentos)
	•	DynamoDB para dados auxiliares

Processamento Assíncrono
	•	SQS para notificações e pagamentos
	•	Lambdas para tratamento dos eventos
	•	DLQs com reprocessamento controlado e alertas
	•	Falhas recorrentes são registradas com contexto para futura análise ou recuperação manual

Observabilidade e Segurança
	•	IAM (menor privilégio), CloudWatch, KMS
	•	Middleware JWT centralizado
	•	Estrutura preparada para logs estruturados e tracing distribuído com ferramentas como AWS X-Ray ou OpenTelemetry, bastando adicionar as dependências e propagar contextos

CI/CD
	•	GitHub Actions com build único e portável
	•	Estrutura apta para deploy por serviço com ECR + ECS

Planejamento Inicial
	•	Desenho manual com fluxo de telas, domínios e rotas antes do primeiro commit
“Esse esboço me guiou para manter coerência entre a experiência do usuário, as rotas da API e os domínios do backend.”

3. Estratégia Arquitetural (DDD, Hexagonal, Go)

Uso de DDD
	•	Organização por domínios (filmes, ingressos, pedidos…) com boundaries claros
	•	usecases e core isolam regras de negócio e modelos
	•	Regras centrais são independentes de HTTP ou banco

Arquitetura Hexagonal
	•	core + usecases representam o centro da aplicação
	•	handlers e adapters são periféricos (entrada e infraestrutura)
	•	Permite troca de banco, protocolo, e reuso do domínio sem refatorar regras

Go como linguagem
	•	Binários leves, rápidos, ideais para deploy em container ou serverless
	•	Facilita paralelismo (ex: goroutines para reserva)
	•	Estrutura clara e simples, favorece leitura e manutenção
	•	Reuso entre projetos: como não está em internal, pode ser importado em outras soluções Go (ex: workers, CLI)
	•	Permite manter uma base de código única para vários serviços, cada um podendo ter sua própria main.go

4. Troubleshooting

Organização e rastreabilidade
	•	Separação por domínio facilita identificar rapidamente onde ocorreu a falha
	•	usecases isolados permitem debug direto de regras sem dependência externa

Execução local
	•	Banco SQLite e Swagger/Insomnia permitem simular todos os fluxos localmente
	•	Reproduzir erros é simples e rápido

Tratamento de erros
	•	Retorno de mensagens claras e contextualizadas (ex: assento já reservado, pedido expirado)
	•	Falhas críticas geram logs e podem ser direcionadas para reprocessamento (DLQ)

Preparado para observabilidade
	•	Estrutura pronta para receber logs estruturados, trace IDs e exportação para X-Ray ou OpenTelemetry
	•	middlewares, handlers e usecases com pontos claros de entrada para inserir logs e spans

5. Encerramento

“O projeto me permitiu aplicar conceitos sólidos de engenharia e arquitetura, desde a estruturação de domínio até a definição de uma arquitetura escalável e flexível. Acredito que ele representa bem meu estilo de desenvolvimento e a preocupação com clareza, performance e sustentabilidade do código.”

“Fico à disposição para detalhar qualquer parte — desde aspectos técnicos até decisões de negócio ou trade-offs.”