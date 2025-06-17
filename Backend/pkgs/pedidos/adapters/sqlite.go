package adapters

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/phraulino/cinetuber/pkgs/pedidos/core"
	"github.com/phraulino/cinetuber/shared/conversores"
	sqlcRepositorio "github.com/phraulino/cinetuber/shared/db/repositorios"
)

type SQLLiteRepoPedidos struct {
	db      *sql.DB
	queries *sqlcRepositorio.Queries
}

func NewSQLLiteRepoPedidos(db *sql.DB) *SQLLiteRepoPedidos {
	return &SQLLiteRepoPedidos{
		db:      db,
		queries: sqlcRepositorio.New(db),
	}
}

func (r *SQLLiteRepoPedidos) CriaPedido(ctx context.Context, usuarioID string) (string, error) {
	pedidoID := uuid.New().String()

	err := r.queries.CriaPedido(ctx, sqlcRepositorio.CriaPedidoParams{
		ID:        pedidoID,
		UsuarioID: usuarioID,
	})
	if err != nil {
		return "", err
	}

	return pedidoID, nil
}

func (r *SQLLiteRepoPedidos) BuscaPedidoPendente(ctx context.Context, usuarioID string) string {
	pedidoID, err := r.queries.BuscaPedidoPendente(ctx, usuarioID)
	if err != nil {
		return ""
	}
	return pedidoID
}

func (r *SQLLiteRepoPedidos) ConsultaPedido(ctx context.Context, pedidoID string) (*core.Pedido, error) {
	pedidoSqlc, err := r.queries.ConsultaPedido(ctx, pedidoID)
	if err != nil {
		return nil, err
	}

	totalPedidoSqlc, err := r.queries.ConsultaTotalPedido(ctx, pedidoID)
	if err != nil {
		return nil, err
	}

	itensPedidoSqlc, err := r.queries.ListaItensPorPedido(ctx, pedidoID)
	if err != nil {
		return nil, err
	}

	itensPedido := make([]*core.ItemPedido, 0, len(itensPedidoSqlc))

	for _, itemPedido := range itensPedidoSqlc {

		item := &core.ItemPedido{
			Nome:       itemPedido.Nome,
			Descricao:  itemPedido.Descricao,
			Tipo:       itemPedido.Tipo,
			Quantidade: itemPedido.Quantidade,
			Status:     itemPedido.Status,
			Total:      itemPedido.Total,
		}

		if itemPedido.Tipo == "ingresso" {
			item.DadosIngresso = &core.DadosIngresso{
				AssentoID: itemPedido.AssentoID,
				SessaoID:  itemPedido.SessaoID,
			}
		}

		itensPedido = append(itensPedido, item)
	}

	pedido := &core.Pedido{
		ID:         pedidoSqlc.ID,
		UsuarioID:  pedidoSqlc.UsuarioID,
		Status:     pedidoSqlc.Status,
		DataPedido: pedidoSqlc.DataCriacao,
		Total:      totalPedidoSqlc,
		Itens:      itensPedido,
	}

	return pedido, nil
}

func (r *SQLLiteRepoPedidos) VerificaQuantidadeItemPedido(ctx context.Context, pedidoID, ItemID, itemTipo string) float64 {
	itensPedidoSqlc, err := r.queries.VerificaQuantidadeItemPedido(ctx, sqlcRepositorio.VerificaQuantidadeItemPedidoParams{
		PedidoID: pedidoID,
		ItemID:   conversores.ParaNullString(&ItemID),
		ItemTipo: itemTipo,
	})
	if err != nil {
		return 0
	}
	return itensPedidoSqlc
}

func (r *SQLLiteRepoPedidos) AtualizaStatusPedido(ctx context.Context, pedidoID, status string) error {
	err := r.queries.AtualizaStatusPedido(ctx, sqlcRepositorio.AtualizaStatusPedidoParams{
		PedidoID: pedidoID,
		Status:   status,
	})
	return err
}

func (r *SQLLiteRepoPedidos) AdicionaProdutoPedido(ctx context.Context, pedidoID, produtoID string, quantidade float64) error {
	err := r.queries.AdicionaProdutoPedido(ctx, sqlcRepositorio.AdicionaProdutoPedidoParams{
		PedidoID:   pedidoID,
		ProdutoID:  produtoID,
		Quantidade: quantidade,
	})
	return err
}

func (r *SQLLiteRepoPedidos) AdicionaProdutosComboPedido(ctx context.Context, pedidoID, comboID string, quantidade float64) error {
	err := r.queries.AdicionaProdutosComboPedido(ctx, sqlcRepositorio.AdicionaProdutosComboPedidoParams{
		PedidoID:   pedidoID,
		ComboID:    comboID,
		Quantidade: quantidade,
	})

	err = r.queries.AdicionaComboPedido(ctx, sqlcRepositorio.AdicionaComboPedidoParams{
		PedidoID:   pedidoID,
		ComboID:    comboID,
		Quantidade: quantidade,
	})
	return err
}

func (r *SQLLiteRepoPedidos) AdicionaIngressoPedido(ctx context.Context, pedidoID, ingressoID string) error {
	err := r.queries.AdicionaIngressoPedido(ctx, sqlcRepositorio.AdicionaIngressoPedidoParams{
		PedidoID:   pedidoID,
		IngressoID: ingressoID,
	})
	return err
}

func (r *SQLLiteRepoPedidos) RemoveProdutoPedido(ctx context.Context, pedidoID, produtoID, itemTipo string) error {
	err := r.queries.RemoveProdutoPedido(ctx, sqlcRepositorio.RemoveProdutoPedidoParams{
		PedidoID:  pedidoID,
		ProdutoID: conversores.ParaNullString(&produtoID),
		ItemTipo:  itemTipo,
	})
	return err
}

func (r *SQLLiteRepoPedidos) RemoveComboPedido(ctx context.Context, pedidoID, comboID, itemTipo string) error {
	err := r.queries.RemoveComboPedido(ctx, sqlcRepositorio.RemoveComboPedidoParams{
		PedidoID: pedidoID,
		ComboID:  conversores.ParaNullString(&comboID),
		ItemTipo: itemTipo,
	})
	return err
}

func (r *SQLLiteRepoPedidos) RemoveIngressoPedido(ctx context.Context, pedidoID, ingressoID string) error {
	err := r.queries.RemoveProdutoPedido(ctx, sqlcRepositorio.RemoveProdutoPedidoParams{
		PedidoID:  pedidoID,
		ProdutoID: conversores.ParaNullString(&ingressoID),
	})
	return err
}
