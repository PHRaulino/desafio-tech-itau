-- Cinemas
INSERT INTO cinemas (id, nome, localizacao, ultima_atualizacao) VALUES
  ('cinema-1', 'Cine SP Center', 'São Paulo - SP', CURRENT_TIMESTAMP),
  ('cinema-2', 'Cine Mooca Plaza', 'São Paulo - SP', CURRENT_TIMESTAMP);

-- Salas
INSERT INTO salas (id, numero, cinema_id, ultima_atualizacao) VALUES
  ('sala-1', 1, 'cinema-1', CURRENT_TIMESTAMP),
  ('sala-2', 2, 'cinema-1', CURRENT_TIMESTAMP),
  ('sala-3', 1, 'cinema-2', CURRENT_TIMESTAMP);

-- Assentos
INSERT INTO assentos (id, numero, fileria, sala_id, ultima_atualizacao) VALUES
  ('assento-1', 1, 'A', 'sala-1', CURRENT_TIMESTAMP),
  ('assento-2', 2, 'A', 'sala-1', CURRENT_TIMESTAMP),
  ('assento-3', 1, 'B', 'sala-1', CURRENT_TIMESTAMP),
  ('assento-4', 1, 'A', 'sala-2', CURRENT_TIMESTAMP);

-- Filmes
INSERT INTO filmes (id, nome, data_lancamento, classificacao, descricao, trailer, capa, expectativa_publico, ultima_atualizacao) VALUES
  ('filme-1', 'Aventura Galáctica', '2024-12-01', '12 anos', 'Filme de ação no espaço.', 'https://trailer.com/1', 'https://image.com/1.jpg', 80, CURRENT_TIMESTAMP),
  ('filme-2', 'Comédia Romântica', '2024-11-20', 'Livre', 'Um amor improvável com muita comédia.', 'https://trailer.com/2', 'https://image.com/2.jpg', 65, CURRENT_TIMESTAMP);

-- Sessões
INSERT INTO sessoes (id, filme_id, sala_id, data_horario, ultima_atualizacao) VALUES
  ('sessao-1', 'filme-1', 'sala-1', '2025-06-15T20:00:00', CURRENT_TIMESTAMP),
  ('sessao-2', 'filme-2', 'sala-2', '2025-06-15T18:00:00', CURRENT_TIMESTAMP);

-- Usuários
INSERT INTO usuarios (id, email, data_nascimento, ultima_atualizacao) VALUES
  ('user-1', 'joao@email.com', '1995-03-22', CURRENT_TIMESTAMP),
  ('user-2', 'ana@email.com', '2000-07-15', CURRENT_TIMESTAMP);

-- Produtos
INSERT INTO produtos (id, nome, descricao, valor, ultima_atualizacao) VALUES
  ('produto-1', 'Pipoca Média', 'Pipoca com manteiga', 15.00, CURRENT_TIMESTAMP),
  ('produto-2', 'Refrigerante 500ml', 'Bebida gelada', 8.50, CURRENT_TIMESTAMP),
  ('produto-3', 'Chocolate', 'Barra de chocolate', 6.00, CURRENT_TIMESTAMP);

-- Combos
INSERT INTO combos (id, nome, descricao, valor, chave_classificacao, ultima_atualizacao) VALUES
  ('combo-1', 'Combo Casal', '2 pipocas médias + 2 refrigerantes', 40.00, 'casal', CURRENT_TIMESTAMP),
  ('combo-2', 'Combo Infantil', '1 pipoca pequena + 1 chocolate', 20.00, 'infantil', CURRENT_TIMESTAMP);

-- Relacionamento combos_produtos
INSERT INTO combos_produtos (combo_id, produto_id, ultima_atualizacao) VALUES
  ('combo-1', 'produto-1', CURRENT_TIMESTAMP),
  ('combo-1', 'produto-2', CURRENT_TIMESTAMP),
  ('combo-2', 'produto-1', CURRENT_TIMESTAMP),
  ('combo-2', 'produto-3', CURRENT_TIMESTAMP);

-- Pedido de exemplo
INSERT INTO pedidos (id, usuario_id, data, valor_total, status, ultima_atualizacao) VALUES
  ('pedido-1', 'user-1', '2025-06-14T13:00:00', 63.50, 'pago', CURRENT_TIMESTAMP);

-- Pedido com itens
INSERT INTO pedidos_produtos (pedido_id, produto_id, ultima_atualizacao) VALUES
  ('pedido-1', 'produto-1', CURRENT_TIMESTAMP),
  ('pedido-1', 'produto-2', CURRENT_TIMESTAMP);

INSERT INTO pedidos_combos (pedido_id, combo_id, chave_classificacao, ultima_atualizacao) VALUES
  ('pedido-1', 'combo-2', 'infantil', CURRENT_TIMESTAMP);

-- Ingresso
INSERT INTO ingressos (id, sessao_id, assento_id, usuario_id, status, valor, ultima_atualizacao) VALUES
  ('ingresso-1', 'sessao-1', 'assento-1', 'user-1', 'disponivel', 30.00, CURRENT_TIMESTAMP);

-- Relacionamento ingresso com pedido
INSERT INTO pedidos_ingressos (pedido_id, ingresso_id, ultima_atualizacao) VALUES
  ('pedido-1', 'ingresso-1', CURRENT_TIMESTAMP);