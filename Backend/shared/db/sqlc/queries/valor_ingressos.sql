-- name: ConsultaValorIngresso :one
SELECT
    *
FROM
    "valor_ingresso"
WHERE
    "tipo" = ?
LIMIT 1