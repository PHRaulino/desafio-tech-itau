-- name: ListaSessoes :many
SELECT
    *
FROM
    "sessoes";

SELECT 
    sessoes.id,
    sessoes.filme_id ,
    sessoes.sala_id ,
    sessoes.data_horario ,
    sessoes.data_criacao ,
    sessoes.ultima_atualizacao, 
    salas.numero
FROM "sessoes" sessoes
inner join "salas" salas 
ON salas.id = sessoes.sala_id;


SELECT 
    sessoes.id,
    sessoes.filme_id ,
    sessoes.sala_id ,
    sessoes.data_horario ,  
    assentos.numero,
    assentos.fileira,
    assentos.sala_id
FROM "sessoes" sessoes
inner join "assentos" assentos 
ON assentos.id = sessoes.sala_id;

