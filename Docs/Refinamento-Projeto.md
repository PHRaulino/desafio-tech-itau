# Case Tecnico Itaú

## Desafio

Você deverá propor uma arquitetura na AWS e desenvolver uma solução funcional (em sua stack de preferência) para um sistema de bilhetagem com as seguintes características:
    • O sistema deve permitir solicitação, reserva e compra de ingressos.
    • Durante o processo de compra, o sistema deve oferecer produtos adicionais como pipoca, chocolate, refrigerante, etc
    • A solução deve conter uma única base de código (um único projeto/solution), mesmo que a arquitetura proposta seja orientada a microserviços. Isso facilitará a apresentação e a avaliação do seu trabalho.

## Antes de começar

Perguntas a serem respondidas como base..

1. Quem é o publico alvo?
1. Qual o consumo da aplicação
1. Quantas sessões podemos ter?
1. Faturamento de bilheteria, ou ticket médio por combo ?
1. Picos de utilização ? (Pré estreias de filmes) filmes aguardados vs 
1. Faixa etaria do publico? Faz sentido ?
1. Perfil do consumidor?
1. Nota do Filme vs Buscas do termo cinema no google ?

Perguntas Fundamentais para Início do Projeto de Bilheteria de Cinema

🎯 1. Público e Perfil de Uso
	•	Quem é o público-alvo? (Idade, localização, hábitos digitais, frequência no cinema)
	•	Qual é o perfil de consumo? (Vai só, casal, família? Compra ingresso e combo?)

📈 2. Consumo e Carga da Aplicação
	•	Qual o volume médio de acessos por dia/mês?
	•	Qual a taxa de conversão de visitantes em compradores?
	•	Quantas sessões podem existir simultaneamente?


💸 3. Faturamento e Produtos
	•	Qual é o ticket médio por compra? (Somando ingresso + combos)
	•	Qual é o faturamento mensal de uma bilheteria comum?
	•	Qual é a proporção de combos vendidos em relação aos ingressos?

🔺 4. Picos e Eventos Especiais
	•	Existem picos previsíveis de uso? (Pré-estreias, estreias em feriados, grandes lançamentos)
	•	Qual o comportamento durante lançamentos populares? (ex: Vingadores, Barbie, etc)
	•	Quais horários concentram o maior tráfego? (ex: 18h às 22h)

🔍 5. Indicadores Externos (mercado e tendências)
	•	Há relação entre nota do filme e procura por ingressos? (usar IMDb + Google Trends)
	•	Existe correlação entre termos buscados como “cinema perto de mim” e demanda real?
	•	Quais regiões do país consomem mais cinema? (IBGE ou Datafolha podem ter isso)

⸻

🧩 Sugestões adicionais
	•	Quais erros ou gargalos sistemas de bilheteria costumam enfrentar? (concorrência na reserva, filas virtuais, lentidão em picos)
	•	Quais funcionalidades o sistema pode ter para se destacar? (ex: recomendação de combos, upsell inteligente)
	•	Qual será o comportamento caso o cliente abandone o carrinho? O ingresso volta ao estoque?