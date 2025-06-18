-- name: ListaProdutos :many
SELECT
    *
FROM
    "produtos";

-- name: ListaProdutosPorCombo :many
SELECT
    *
FROM
    "produtos" produtos
    INNER JOIN "combos_produtos" combos ON produtos.id = combos.produto_id
WHERE
    combos.combo_id = ?;

-- name: ListaCombos :many
SELECT
    *
FROM
    "combos";