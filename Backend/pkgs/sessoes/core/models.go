package core

import "time"

type Sessao struct {
	ID            string    `json:"id"`
	FilmeID       string    `json:"filme_id"`
	SalaID        string    `json:"sala_id"`
	SalaDescricao string    `json:"sala_descricao"`
	Cinema        string    `json:"cinema"`
	Status        string    `json:"status"`
	DataSessao    time.Time `json:"data_sessao"`
}

type CriaSessao struct {
	FilmeID    string    `json:"filme_id"`
	SalaID     string    `json:"sala_id"`
	DataSessao time.Time `json:"data_sessao"`
}
type CriaReserva struct {
	SessaoID     string `json:"sessao_id"`
	UsuarioID    string `json:"usuario_id"`
	TipoIngresso string `json:"tipo_ingresso"`
}

type SessaoAssento struct {
	AssentoID string `json:"assento_id"`
	SalaID    string `json:"sala_id"`
	Status    string `json:"status"`
	Fileira   string `json:"fileira"`
	Numero    int64  `json:"numero"`
	Descricao string `json:"descricao"`
}

type BuscaSessao struct {
	FilmeID    *string    `json:"filme_id"`
	CinemaID   *string    `json:"cinema_id"`
	SalaID     *string    `json:"sala_id"`
	DataSessao *time.Time `json:"data_sessao"`
}
