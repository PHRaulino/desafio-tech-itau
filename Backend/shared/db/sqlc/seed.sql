INSERT INTO
  cinemas (
    id,
    nome,
    localizacao,
    data_criacao,
    ultima_atualizacao
  )
VALUES
  (
    'd99ecaf0-a4ff-4176-9a95-433a5f53aefe',
    'Cine SP Center',
    'São Paulo - SP',
    '2025-06-14T20:14:50.829095',
    '2025-06-14T20:14:50.829095'
  );

INSERT INTO
  salas (
    id,
    numero,
    cinema_id,
    data_criacao,
    ultima_atualizacao
  )
VALUES
  (
    '2e19142e-4e6b-4311-b25e-6694a4ba6dd3',
    1,
    'd99ecaf0-a4ff-4176-9a95-433a5f53aefe',
    '2025-06-14T20:14:50.829095',
    '2025-06-14T20:14:50.829095'
  );

-- Assentos
INSERT INTO
  assentos (
    id,
    numero,
    fileria,
    sala_id,
    data_criacao,
    ultima_atualizacao
  )
VALUES
  (
    '75be71fe-ce90-4429-9bc1-5b9cfd3f7ef1',
    1,
    'A',
    '2e19142e-4e6b-4311-b25e-6694a4ba6dd3',
    '2025-06-14T20:14:50.829095',
    '2025-06-14T20:14:50.829095'
  ),
  (
    '0a275958-ee70-4537-81b0-ac9e59cb731a',
    2,
    'A',
    '2e19142e-4e6b-4311-b25e-6694a4ba6dd3',
    '2025-06-14T20:14:50.829095',
    '2025-06-14T20:14:50.829095'
  ),
  (
    '7ddd4187-73f3-42fc-aea8-c11fa3428b93',
    3,
    'A',
    '2e19142e-4e6b-4311-b25e-6694a4ba6dd3',
    '2025-06-14T20:14:50.829095',
    '2025-06-14T20:14:50.829095'
  ),
  (
    '3ee4605c-86a9-450f-835b-19979655a8c9',
    4,
    'A',
    '2e19142e-4e6b-4311-b25e-6694a4ba6dd3',
    '2025-06-14T20:14:50.829095',
    '2025-06-14T20:14:50.829095'
  ),
  (
    'e0b917bb-4196-4bdd-a491-6652896fc99a',
    5,
    'A',
    '2e19142e-4e6b-4311-b25e-6694a4ba6dd3',
    '2025-06-14T20:14:50.829095',
    '2025-06-14T20:14:50.829095'
  );

-- Filme
INSERT INTO
  filmes (
    id,
    nome,
    lancamento,
    classificacao,
    descricao,
    trailer,
    capa,
    data_criacao,
    ultima_atualizacao
  )
VALUES
  (
    'b37485a7-fcac-4d6e-a924-80468c6b33c0',
    'A Volta dos Que Não Foram',
    '2025-07-01',
    '12',
    'Uma comédia sobre improváveis heróis.',
    'https://youtu.be/trailer',
    'https://img.com/capa.jpg',
    '2025-06-14T20:14:50.829095',
    '2025-06-14T20:14:50.829095'
  );

-- Sessão
INSERT INTO
  sessoes (
    id,
    filme_id,
    sala_id,
    data_horario,
    data_criacao,
    ultima_atualizacao
  )
VALUES
  (
    'ceddee5e-b4df-4b12-a22f-e7e76aeb9d82',
    'b37485a7-fcac-4d6e-a924-80468c6b33c0',
    '2e19142e-4e6b-4311-b25e-6694a4ba6dd3',
    '2025-07-01T20:00:00',
    '2025-06-14T20:14:50.829095',
    '2025-06-14T20:14:50.829095'
  );

-- Usuário
INSERT INTO
  usuarios (
    id,
    email,
    data_nascimento,
    data_criacao,
    ultima_atualizacao
  )
VALUES
  (
    '66233cc6-7d3e-4070-92fa-66c64093b006',
    'user@example.com',
    '1990-05-20',
    '2025-06-14T20:14:50.829095',
    '2025-06-14T20:14:50.829095'
  );

-- Produtos
INSERT INTO
  produtos (
    id,
    nome,
    descricao,
    valor,
    data_criacao,
    ultima_atualizacao
  )
VALUES
  (
    '27508be3-bac4-4483-a53f-8925dde988a4',
    'Pipoca Média',
    'Pipoca salgada média.',
    15.0,
    '2025-06-14T20:14:50.829095',
    '2025-06-14T20:14:50.829095'
  ),
  (
    '5b6d14a4-3d73-4981-be91-653037cb4734',
    'Refrigerante Lata',
    'Lata de refrigerante 350ml.',
    8.0,
    '2025-06-14T20:14:50.829095',
    '2025-06-14T20:14:50.829095'
  ),
  (
    '3525dd22-60f6-4891-bbfa-79cfe8481fe5',
    'Chocolate',
    'Chocolate ao leite.',
    6.5,
    '2025-06-14T20:14:50.829095',
    '2025-06-14T20:14:50.829095'
  );

-- Combo
INSERT INTO
  combos (
    id,
    nome,
    descricao,
    valor,
    data_criacao,
    ultima_atualizacao
  )
VALUES
  (
    '9073c461-2fad-4b33-963e-9961ad31868a',
    'Combo Pipoca + Refri',
    'Pipoca média + Refrigerante',
    22.00,
    '2025-06-14T20:14:50.829095',
    '2025-06-14T20:14:50.829095'
  );

-- Produtos no combo
INSERT INTO
  combos_produtos (
    combo_id,
    produto_id,
    data_criacao,
    ultima_atualizacao
  )
VALUES
  (
    '9073c461-2fad-4b33-963e-9961ad31868a',
    '27508be3-bac4-4483-a53f-8925dde988a4',
    '2025-06-14T20:14:50.829095',
    '2025-06-14T20:14:50.829095'
  ),
  (
    '9073c461-2fad-4b33-963e-9961ad31868a',
    '5b6d14a4-3d73-4981-be91-653037cb4734',
    '2025-06-14T20:14:50.829095',
    '2025-06-14T20:14:50.829095'
  );

-- Tipos de ingresso
INSERT INTO
  valor_ingresso (id, valor, tipo, data_criacao, ultima_atualizacao)
VALUES
  (
    'cc320201-6b67-4409-9c4e-2d2ee1c2cb16',
    30.00,
    'inteira',
    '2025-06-14T20:14:50.829095',
    '2025-06-14T20:14:50.829095'
  ),
  (
    '4d381e85-abdc-4b2e-838f-caa62eda59b0',
    15.00,
    'meia',
    '2025-06-14T20:14:50.829095',
    '2025-06-14T20:14:50.829095'
  );