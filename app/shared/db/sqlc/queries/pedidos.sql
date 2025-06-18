-- name: ListaItensPorPedido :many
WITH
    filmes_sessao AS (
        SELECT
            sessoes.id,
            filmes.nome AS nome_filme,
            salas.numero || ' (' || cinemas.nome || ')' AS nome_sala
        FROM
            sessoes
            INNER JOIN filmes ON sessoes.filme_id = filmes.id
            INNER JOIN salas ON sessoes.sala_id = salas.id
            INNER JOIN cinemas ON salas.cinema_id = cinemas.id
    ),
    pedidos_produtos_detalhe AS (
        SELECT
            pedidos_produtos.nome,
            pedidos_produtos.descricao,
            CAST(pedidos_produtos.quantidade AS INTEGER) quantidade,
            pedidos_produtos.total,
            'reservado' AS status,
            pedidos_produtos.tipo,
            '' AS ingresso_id,
            '' AS sessao_id,
            '' AS assento_id
        FROM
            pedidos_produtos
        WHERE
            pedidos_produtos.pedido_id = :pedido_id
    ),
    pedidos_ingressos_detalhe AS (
        SELECT
            filmes_sessao.nome_filme AS nome,
            filmes_sessao.nome_filme || ' - ' || assentos.fileira || assentos.numero || ' (Sala: ' || filmes_sessao.nome_sala || ')' AS descricao,
            1 as quantidade,
            ingressos.valor as total,
            ingressos.status AS status,
            'ingresso' AS tipo,
            ingressos.id AS ingresso_id,
            ingressos.sessao_id AS sessao_id,
            ingressos.assento_id AS assento_id
        FROM
            pedidos_ingressos
            INNER JOIN ingressos ON pedidos_ingressos.ingresso_id = ingressos.id
            INNER JOIN filmes_sessao ON ingressos.sessao_id = filmes_sessao.id
            INNER JOIN assentos ON ingressos.assento_id = assentos.id
        WHERE
            pedidos_ingressos.pedido_id = :pedido_id
    )
SELECT
    *
FROM
    pedidos_produtos_detalhe
UNION ALL
SELECT
    *
FROM
    pedidos_ingressos_detalhe;

-- name: CriaPedido :exec
INSERT INTO
    pedidos (
        id,
        usuario_id,
        status,
        data_criacao,
        ultima_atualizacao
    )
VALUES
    (
        :id,
        :usuario_id,
        'pendente',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    );

-- name: ConsultaTotalPedido :one
WITH total_ingressos AS (
    SELECT
        SUM(ingressos.valor) AS total
    FROM
        pedidos_ingressos
        INNER JOIN ingressos ON pedidos_ingressos.ingresso_id = ingressos.id
    WHERE
        pedidos_ingressos.pedido_id = :pedido_id
), total_produtos AS (
    SELECT
        SUM(pedidos_produtos.total) AS total
    FROM
        pedidos_produtos
    WHERE
        pedidos_produtos.pedido_id = :pedido_id
)
SELECT
    CAST(COALESCE(total_ingressos.total, 0) + COALESCE(total_produtos.total, 0) AS REAL)AS total
FROM
    total_ingressos,
    total_produtos;

-- name: AtualizaStatusPedido :exec
UPDATE
    pedidos
SET
    status = :status,
    ultima_atualizacao = CURRENT_TIMESTAMP
WHERE
    id = :pedido_id;

-- name: AdicionaProdutoPedido :exec
INSERT INTO pedidos_produtos (
    pedido_id,
    produto_id,
    nome,
    descricao,
    quantidade,
    total,
    tipo,
    data_criacao,
    ultima_atualizacao
)
SELECT
    :pedido_id,
    produtos.id,
    produtos.nome,
    produtos.descricao,
    :quantidade,
    produtos.valor * :quantidade,
    'avulso',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
FROM produtos
WHERE produtos.id = :produto_id;

-- name: AdicionaComboPedido :exec
INSERT INTO pedidos_produtos (
    pedido_id,
    combo_id,
    nome,
    descricao,
    quantidade,
    total,
    tipo,
    data_criacao,
    ultima_atualizacao
)
SELECT
    :pedido_id,
    c.id,
    c.nome,
    c.descricao,
    :quantidade,
    c.valor * :quantidade,
    'combo',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
FROM combos c
WHERE c.id = :combo_id;

-- name: AdicionaProdutosComboPedido :exec

INSERT INTO pedidos_produtos (
    pedido_id,
    produto_id,
    combo_id,
    nome,
    descricao,
    quantidade,
    total,
    tipo,
    data_criacao,
    ultima_atualizacao
)
SELECT
    :pedido_id,
    p.id,
    c.id,
    p.nome,
    p.descricao,
    :quantidade,
    0,
    'combo',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
FROM combos c
JOIN combos_produtos cp ON c.id = cp.combo_id
JOIN produtos p ON cp.produto_id = p.id
WHERE c.id = :combo_id;

-- name: AdicionaIngressoPedido :exec
INSERT INTO pedidos_ingressos (
    pedido_id,
    ingresso_id,
    data_criacao,
    ultima_atualizacao
)
SELECT
    :pedido_id,
    i.id,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
FROM ingressos i
WHERE i.id = :ingresso_id;

-- name: RemoveProdutoPedido :exec
DELETE FROM pedidos_produtos
WHERE pedido_id = :pedido_id AND produto_id = :produto_id AND tipo = :item_tipo;

-- name: RemoveComboPedido :exec
DELETE FROM pedidos_produtos
WHERE pedido_id = :pedido_id AND combo_id = :combo_id AND tipo = :item_tipo;

-- name: RemoveIngressoPedido :exec
DELETE FROM pedidos_ingressos
WHERE pedido_id = :pedido_id AND ingresso_id = :ingresso_id;

-- name: ConsultaPedido :one
SELECT
    p.id,
    p.usuario_id,
    p.status,
    p.data_criacao,
    p.ultima_atualizacao
FROM
    pedidos p
WHERE
    p.id = :pedido_id;

-- name: VerificaQuantidadeItemPedido :one
SELECT
    pedidos_produtos.quantidade
FROM
    pedidos_produtos
WHERE
    pedidos_produtos.pedido_id = :pedido_id
    AND pedidos_produtos.tipo = :item_tipo
    AND (
        pedidos_produtos.produto_id = :item_id
        OR pedidos_produtos.combo_id = :item_id
    );

-- name: BuscaPedidoPendente :one
SELECT
    id
FROM
    pedidos
WHERE
    usuario_id = :usuario_id
    AND status in ('pendente', 'em pagamento')