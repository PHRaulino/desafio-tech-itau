-- name: ListaSessoes :many
SELECT
    sessoes.id,
    sessoes.filme_id,
    sessoes.sala_id,
    CAST(CONCAT('Sala ', salas.numero) AS VARCHAR) AS sala_descricao,
    cinemas.nome AS nome_cinema,
    sessoes.status,
    sessoes.data_sessao
FROM
    sessoes
INNER JOIN salas ON sessoes.sala_id = salas.id
INNER JOIN cinemas ON salas.cinema_id = salas.cinema_id
WHERE
    (:filme_id IS NULL OR sessoes.filme_id = :filme_id)
    AND (:sala_id IS NULL OR sessoes.sala_id = :sala_id)
    AND (:data_sessao IS NULL OR sessoes.data_sessao = :data_sessao)
    AND (:cinema_id IS NULL OR salas.cinema_id = :cinema_id);

-- name: CriaSessao :exec
INSERT INTO sessoes
    (
        id,
        filme_id,
        sala_id,
        status,
        data_sessao,
        data_criacao,
        ultima_atualizacao
   )
   VALUES
    (
        :sessao_id,
        :filme_id,
        :sala_id,
        'aberta',
        :data_sessao,
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    );

-- name: ListaAssentos :many
WITH assentos_sessao as (
    SELECT assentos.id,
    assentos.sala_id,
    assentos.fileira,
    assentos.numero FROM assentos
    INNER JOIN sessoes ON assentos.sala_id = sessoes.sala_id
    WHERE sessoes.id = :sessao_id
), ingressos_sessao as (
    SELECT ingressos.assento_id, ingressos.status FROM ingressos
    WHERE ingressos.sessao_id = :sessao_id
    AND ingressos.status NOT IN ('expirado', 'invalido')
), todos_assento as (
SELECT
    assentos_sessao.id as assento_id,
    assentos_sessao.sala_id,
    assentos_sessao.fileira,
    assentos_sessao.numero,
    CAST(CONCAT(assentos_sessao.fileira, assentos_sessao.numero) AS VARCHAR) AS descricao,
    CAST(CASE
        WHEN ingressos_sessao.status IS NULL THEN 'disponivel'
        WHEN ingressos_sessao.status = 'confirmado' THEN 'ocupado'
        WHEN ingressos_sessao.status in ('reservado', 'em pagamento') THEN 'reservado'
    ELSE ingressos_sessao.status
END AS VARCHAR) AS status
FROM assentos_sessao
LEFT JOIN ingressos_sessao
ON assentos_sessao.id = ingressos_sessao.assento_id)
SELECT * FROM todos_assento
ORDER BY status
;

-- name: ListaAssentosReservados :many
SELECT ingressos.id FROM ingressos
WHERE ingressos.sessao_id = :sessao_id
AND ingressos.status = 'reservado';