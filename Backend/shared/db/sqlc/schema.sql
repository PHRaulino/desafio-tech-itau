CREATE TABLE
    usuarios (
        id TEXT PRIMARY KEY,
        email TEXT NOT NULL UNIQUE,
        data_nascimento DATETIME NOT NULL,
        data_criacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        ultima_atualizacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    cinemas (
        id TEXT PRIMARY KEY,
        nome TEXT NOT NULL UNIQUE,
        localizacao TEXT NOT NULL,
        data_criacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        ultima_atualizacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    salas (
        id TEXT PRIMARY KEY,
        numero INTEGER NOT NULL,
        cinema_id TEXT NOT NULL,
        data_criacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        ultima_atualizacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (cinema_id) REFERENCES cinemas (id)
    );

CREATE TABLE
    assentos (
        id TEXT PRIMARY KEY,
        numero INTEGER NOT NULL,
        fileira TEXT NOT NULL,
        sala_id TEXT NOT NULL,
        data_criacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        ultima_atualizacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (sala_id) REFERENCES salas (id)
    );

CREATE TABLE
    filmes (
        id TEXT PRIMARY KEY,
        nome TEXT NOT NULL UNIQUE,
        lancamento DATE NOT NULL,
        classificacao TEXT NOT NULL,
        descricao TEXT NOT NULL,
        trailer TEXT NOT NULL,
        capa TEXT NOT NULL,
        data_criacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        ultima_atualizacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    sessoes (
        id TEXT PRIMARY KEY,
        filme_id TEXT NOT NULL,
        sala_id TEXT NOT NULL,
        status TEXT NOT NULL CHECK (status IN ('finalizada', 'aberta', 'lotada')),
        data_sessao TIMESTAMP NOT NULL,
        data_criacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        ultima_atualizacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (filme_id) REFERENCES filmes (id),
        FOREIGN KEY (sala_id) REFERENCES salas (id)
    );

CREATE TABLE
    ingressos (
        id TEXT PRIMARY KEY,
        sessao_id TEXT NOT NULL,
        assento_id TEXT NOT NULL,
        usuario_id TEXT NOT NULL,
        status TEXT NOT NULL CHECK (
            status IN ('disponivel', 'utilizado', 'reservado')
        ),
        valor NUMERIC NOT NULL,
        data_criacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        ultima_atualizacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (sessao_id) REFERENCES sessoes (id),
        FOREIGN KEY (assento_id) REFERENCES assentos (id),
        FOREIGN KEY (usuario_id) REFERENCES usuarios (id),
        UNIQUE (sessao_id, assento_id)
    );

CREATE TABLE
    produtos (
        id TEXT PRIMARY KEY,
        nome TEXT NOT NULL UNIQUE,
        descricao TEXT NOT NULL,
        valor NUMERIC NOT NULL,
        data_criacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        ultima_atualizacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    combos (
        id TEXT PRIMARY KEY,
        nome TEXT NOT NULL UNIQUE,
        descricao TEXT NOT NULL,
        valor NUMERIC NOT NULL,
        data_criacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        ultima_atualizacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    combos_produtos (
        combo_id TEXT NOT NULL,
        produto_id TEXT NOT NULL,
        data_criacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        ultima_atualizacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (combo_id, produto_id),
        FOREIGN KEY (combo_id) REFERENCES combos (id),
        FOREIGN KEY (produto_id) REFERENCES produtos (id)
    );

CREATE TABLE
    pedidos (
        id TEXT PRIMARY KEY,
        usuario_id TEXT NOT NULL,
        status TEXT NOT NULL CHECK (status IN ('pago', 'pendente', 'cancelado')),
        data_criacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        ultima_atualizacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (usuario_id) REFERENCES usuarios (id)
    );

CREATE TABLE
    pedidos_ingressos (
        pedido_id TEXT NOT NULL,
        ingresso_id TEXT NOT NULL,
        data_criacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        ultima_atualizacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (pedido_id, ingresso_id),
        FOREIGN KEY (pedido_id) REFERENCES pedidos (id),
        FOREIGN KEY (ingresso_id) REFERENCES ingressos (id)
    );

CREATE TABLE
    pedidos_produtos (
        pedido_id TEXT NOT NULL,
        produto_id TEXT,
        combo_id TEXT,
        nome TEXT NOT NULL,
        descricao TEXT NOT NULL,
        quantidade NUMERIC NOT NULL,
        total NUMERIC NOT NULL,
        tipo TEXT NOT NULL CHECK (tipo IN ('combo', 'avulso')),
        data_criacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        ultima_atualizacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (pedido_id) REFERENCES pedidos (id),
        FOREIGN KEY (produto_id) REFERENCES produtos (id),
        FOREIGN KEY (combo_id) REFERENCES combos (id)
    );

CREATE TABLE
    valor_ingresso (
        id TEXT PRIMARY KEY,
        valor NUMERIC NOT NULL,
        tipo TEXT NOT NULL UNIQUE CHECK (tipo IN ('inteira', 'meia')),
        data_criacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        ultima_atualizacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
    );