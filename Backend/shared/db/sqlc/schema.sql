CREATE TABLE usuarios (
    id TEXT PRIMARY KEY,
    email TEXT NOT NULL,
    data_nascimento DATETIME NOT NULL,
    data_criacao DATETIME NOT NULL,
    ultima_atualizacao DATETIME NOT NULL
);

CREATE TABLE cinemas (
    id TEXT PRIMARY KEY,
    nome TEXT NOT NULL,
    localizacao TEXT NOT NULL,
    data_criacao DATETIME NOT NULL,
    ultima_atualizacao DATETIME NOT NULL
);

CREATE TABLE salas (
    id TEXT PRIMARY KEY,
    numero INTEGER NOT NULL,
    cinema_id TEXT NOT NULL,
    data_criacao DATETIME NOT NULL,
    ultima_atualizacao DATETIME NOT NULL,
    FOREIGN KEY (cinema_id) REFERENCES cinemas(id)
);

CREATE TABLE assentos (
    id TEXT PRIMARY KEY,
    numero INTEGER NOT NULL,
    fileria TEXT NOT NULL,
    sala_id TEXT NOT NULL,
    data_criacao DATETIME NOT NULL,
    ultima_atualizacao DATETIME NOT NULL,
    FOREIGN KEY (sala_id) REFERENCES salas(id)
);

CREATE TABLE filmes (
    id TEXT PRIMARY KEY,
    nome TEXT NOT NULL,
    lancamento DATE NOT NULL,
    classificacao TEXT NOT NULL,
    descricao TEXT NOT NULL,
    trailer TEXT NOT NULL,
    capa TEXT NOT NULL,
    data_criacao DATETIME NOT NULL,
    ultima_atualizacao DATETIME NOT NULL
);

CREATE TABLE sessoes (
    id TEXT PRIMARY KEY,
    filme_id TEXT NOT NULL,
    sala_id TEXT NOT NULL,
    data_horario TIMESTAMP NOT NULL,
    data_criacao DATETIME NOT NULL,
    ultima_atualizacao DATETIME NOT NULL,
    FOREIGN KEY (filme_id) REFERENCES filmes(id),
    FOREIGN KEY (sala_id) REFERENCES salas(id)
);

CREATE TABLE ingressos (
    id TEXT PRIMARY KEY,
    sessao_id TEXT NOT NULL,
    assento_id TEXT NOT NULL,
    usuario_id TEXT NOT NULL,
    status TEXT NOT NULL CHECK(status IN ('disponivel', 'utilizado')),
    valor NUMERIC NOT NULL,
    data_criacao DATETIME NOT NULL,
    ultima_atualizacao DATETIME NOT NULL,
    FOREIGN KEY (sessao_id) REFERENCES sessoes(id),
    FOREIGN KEY (assento_id) REFERENCES assentos(id),
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id),
    UNIQUE (sessao_id, assento_id)
);

CREATE TABLE produtos (
    id TEXT PRIMARY KEY,
    nome TEXT NOT NULL,
    descricao TEXT NOT NULL,
    valor NUMERIC NOT NULL,
    data_criacao DATETIME NOT NULL,
    ultima_atualizacao DATETIME NOT NULL
);

CREATE TABLE combos (
    id TEXT PRIMARY KEY,
    nome TEXT NOT NULL,
    descricao TEXT NOT NULL,
    valor NUMERIC NOT NULL,
    data_criacao DATETIME NOT NULL,
    ultima_atualizacao DATETIME NOT NULL
);

CREATE TABLE combos_produtos (
    combo_id TEXT NOT NULL,
    produto_id TEXT NOT NULL,
    data_criacao DATETIME NOT NULL,
    ultima_atualizacao DATETIME NOT NULL,
    PRIMARY KEY (combo_id, produto_id),
    FOREIGN KEY (combo_id) REFERENCES combos(id),
    FOREIGN KEY (produto_id) REFERENCES produtos(id)
);

CREATE TABLE pedidos (
    id TEXT PRIMARY KEY,
    usuario_id TEXT NOT NULL,
    data TIMESTAMP NOT NULL,
    valor_total NUMERIC NOT NULL,
    status TEXT NOT NULL CHECK(status IN ('pago', 'pendente', 'cancelado')),
    data_criacao DATETIME NOT NULL,
    ultima_atualizacao DATETIME NOT NULL,
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id)
);

CREATE TABLE pedidos_ingressos (
    pedido_id TEXT NOT NULL,
    ingresso_id TEXT NOT NULL,
    data_criacao DATETIME NOT NULL,
    ultima_atualizacao DATETIME NOT NULL,
    PRIMARY KEY (pedido_id, ingresso_id),
    FOREIGN KEY (pedido_id) REFERENCES pedidos(id),
    FOREIGN KEY (ingresso_id) REFERENCES ingressos(id)
);

CREATE TABLE pedidos_produtos (
    pedido_id TEXT NOT NULL,
    produto_id TEXT NOT NULL,
    data_criacao DATETIME NOT NULL,
    ultima_atualizacao DATETIME NOT NULL,
    PRIMARY KEY (pedido_id, produto_id),
    FOREIGN KEY (pedido_id) REFERENCES pedidos(id),
    FOREIGN KEY (produto_id) REFERENCES produtos(id)
);

CREATE TABLE pedidos_combos (
    pedido_id TEXT NOT NULL,
    combo_id TEXT NOT NULL,
    data_criacao DATETIME NOT NULL,
    ultima_atualizacao DATETIME NOT NULL,
    PRIMARY KEY (pedido_id, combo_id),
    FOREIGN KEY (pedido_id) REFERENCES pedidos(id),
    FOREIGN KEY (combo_id) REFERENCES combos(id)
);

CREATE TABLE valor_ingresso (
    id TEXT PRIMARY KEY,
    valor NUMERIC NOT NULL,
    tipo TEXT NOT NULL CHECK(tipo IN ('inteira', 'meia')),
    data_criacao DATETIME NOT NULL,
    ultima_atualizacao DATETIME NOT NULL
);