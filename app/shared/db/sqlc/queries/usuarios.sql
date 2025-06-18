-- name: ListaUsuarios :many
SELECT * FROM usuarios;

-- name: CriaUsuario :exec
INSERT INTO usuarios
    (
        nome,
        email,
        id
    )
VALUES
    (
        :nome,
        :email,
        :id
    );

-- name: BuscaUsuario :one
SELECT * FROM usuarios
WHERE usuarios.id = :usuario_id;