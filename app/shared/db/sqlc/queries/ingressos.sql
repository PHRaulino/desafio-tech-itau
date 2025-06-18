-- name: ConsultaValorIngresso :one
SELECT
    *
FROM
    "valor_ingresso"
WHERE
    "tipo" = ?
LIMIT
    1;

-- name: CriaIngresso :exec
INSERT INTO ingressos
    (
        id,
        sessao_id,
        assento_id,
        usuario_id,
        status,
        valor,
        data_criacao,
        ultima_atualizacao
    )
    VALUES
    (
        :ingresso_id,
        :sessao_id,
        :assento_id,
        :usuario_id,
        'reservado',
        :valor,
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    );

-- name: AtualizaStatusIngresso :exec
UPDATE
    ingressos
SET
    status = :status,
    ultima_atualizacao = CURRENT_TIMESTAMP
WHERE
    id = :ingresso_id;

-- name: ListaIngressos :many
SELECT
    ingressos.id AS ingresso_id,
    ingressos.sessao_id,
    ingressos.assento_id,
    ingressos.usuario_id,
    ingressos.status,
    ingressos.valor
FROM ingressos
WHERE
    (:sessao_id IS NULL OR ingressos.sessao_id = :sessao_id)
    AND (:assento_id IS NULL OR ingressos.sessao_id = :assento_id)
    AND (:usuario_id IS NULL OR ingressos.sessao_id = :usuario_id);