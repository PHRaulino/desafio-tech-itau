CREATE TABLE usuarios (
    id TEXT PRIMARY KEY,
    email TEXT,
    data_nascimento DATETIME,
    ultima_atualizacao DATETIME
);

CREATE TABLE cinemas (
    id TEXT PRIMARY KEY,
    nome TEXT,
    localizacao TEXT,
    ultima_atualizacao DATETIME
);

CREATE TABLE salas (
    id TEXT PRIMARY KEY,
    numero INTEGER,
    cinema_id TEXT,
    ultima_atualizacao DATETIME,
    FOREIGN KEY (cinema_id) REFERENCES cinemas(id)
);

CREATE TABLE assentos (
    id TEXT PRIMARY KEY,
    numero INTEGER,
    fileria TEXT,
    sala_id TEXT,
    ultima_atualizacao DATETIME,
    FOREIGN KEY (sala_id) REFERENCES salas(id)
);

CREATE TABLE filmes (
    id TEXT PRIMARY KEY,
    nome TEXT,
    data_lancamento DATE,
    classificacao TEXT,
    descricao TEXT,
    trailer TEXT,
    capa TEXT,
    expectativa_publico INTEGER,
    ultima_atualizacao DATETIME
);

CREATE TABLE sessoes (
    id TEXT PRIMARY KEY,
    filme_id TEXT,
    sala_id TEXT,
    data_horario TIMESTAMP,
    ultima_atualizacao DATETIME,
    FOREIGN KEY (filme_id) REFERENCES filmes(id),
    FOREIGN KEY (sala_id) REFERENCES salas(id)
);

CREATE TABLE ingressos (
    id TEXT PRIMARY KEY,
    sessao_id TEXT,
    assento_id TEXT,
    usuario_id TEXT,
    status TEXT CHECK(status IN ('disponivel', 'utilizado')),
    valor NUMERIC,
    ultima_atualizacao DATETIME,
    FOREIGN KEY (sessao_id) REFERENCES sessoes(id),
    FOREIGN KEY (assento_id) REFERENCES assentos(id),
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id),
    UNIQUE (sessao_id, assento_id)
);

CREATE TABLE produtos (
    id TEXT PRIMARY KEY,
    nome TEXT,
    descricao TEXT,
    valor NUMERIC,
    ultima_atualizacao DATETIME
);

CREATE TABLE combos (
    id TEXT PRIMARY KEY,
    nome TEXT,
    descricao TEXT,
    valor NUMERIC,
    chave_classificacao TEXT,
    ultima_atualizacao DATETIME
);

CREATE TABLE combos_produtos (
    combo_id TEXT,
    produto_id TEXT,
    ultima_atualizacao DATETIME,
    PRIMARY KEY (combo_id, produto_id),
    FOREIGN KEY (combo_id) REFERENCES combos(id),
    FOREIGN KEY (produto_id) REFERENCES produtos(id)
);

CREATE TABLE pedidos (
    id TEXT PRIMARY KEY,
    usuario_id TEXT,
    data TIMESTAMP,
    valor_total NUMERIC,
    status TEXT CHECK(status IN ('pago', 'pendente', 'cancelado')),
    ultima_atualizacao DATETIME,
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id)
);

CREATE TABLE pedidos_ingressos (
    pedido_id TEXT,
    ingresso_id TEXT,
    ultima_atualizacao DATETIME,
    PRIMARY KEY (pedido_id, ingresso_id),
    FOREIGN KEY (pedido_id) REFERENCES pedidos(id),
    FOREIGN KEY (ingresso_id) REFERENCES ingressos(id)
);

CREATE TABLE pedidos_produtos (
    pedido_id TEXT,
    produto_id TEXT,
    ultima_atualizacao DATETIME,
    PRIMARY KEY (pedido_id, produto_id),
    FOREIGN KEY (pedido_id) REFERENCES pedidos(id),
    FOREIGN KEY (produto_id) REFERENCES produtos(id)
);

CREATE TABLE pedidos_combos (
    pedido_id TEXT,
    combo_id TEXT,
    chave_classificacao TEXT,
    ultima_atualizacao DATETIME,
    PRIMARY KEY (pedido_id, combo_id),
    FOREIGN KEY (pedido_id) REFERENCES pedidos(id),
    FOREIGN KEY (combo_id) REFERENCES combos(id)
);