// pacote conversores para auxiliar na conversão de tipos SQL
package conversores

import (
	"database/sql"
	"time"
)
//
// Funções utilitárias para conversão de tipos SQL.
//

// NullStringParaPonteiro converte um sql.NullString para um ponteiro de string.
func NullStringParaPonteiro(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

// NullBoolParaPonteiro converte um sql.NullBool para um ponteiro de bool.
func NullBoolParaPonteiro(nb sql.NullBool) *bool {
	if nb.Valid {
		return &nb.Bool
	}
	return nil
}

// NullFloatParaPonteiro converte um sql.NullFloat64 para um ponteiro de float64.
func NullFloatParaPonteiro(nf sql.NullFloat64) *float64 {
	if nf.Valid {
		return &nf.Float64
	}
	return nil
}

// ParaNullInt64 converte um ponteiro de int64 para sql.NullInt64.
func ParaNullInt64(s *int64) sql.NullInt64 {
	if s == nil {
		return sql.NullInt64{Int64: 0, Valid: false}
	}
	return sql.NullInt64{Int64: *s, Valid: true}
}

// ParaNullFloat64 converte um ponteiro de float64 para sql.NullFloat64.
func ParaNullFloat64(f *float64) sql.NullFloat64 {
	if f == nil {
		return sql.NullFloat64{Float64: 0, Valid: false}
	}
	return sql.NullFloat64{Float64: *f, Valid: true}
}

// ParaNullTime converte um ponteiro de time.Time para sql.NullTime.
func ParaNullTime(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{Time: time.Time{}, Valid: false}
	}
	return sql.NullTime{Time: *t, Valid: true}
}

// ParaNullString converte um ponteiro de string para sql.NullString.
func ParaNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: *s, Valid: true}
}

// ParaNullBool converte um ponteiro de bool para sql.NullBool.
func ParaNullBool(b *bool) sql.NullBool {
	if b == nil {
		return sql.NullBool{Bool: false, Valid: false}
	}
	return sql.NullBool{Bool: *b, Valid: true}
}
// NullStringParaPonteiro converte um sql.NullString para um ponteiro de string.
func NullStringToPointer(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String // Retorna o ponteiro para o valor de String
	}
	return nil // Retorna nil se o valor for NULL
}

// NullBoolToPointer converts a sql.NullBool to a pointer to a bool.
func NullBoolToPointer(nb sql.NullBool) *bool {
	if nb.Valid {
		return &nb.Bool // Retorna o ponteiro para o valor de Bool
	}
	return nil // Retorna nil se o valor for NULL
}

// NullFloatToPointer converts a sql.NullFloat64 to a pointer to a float64.
func NullFloatToPointer(nf sql.NullFloat64) *float64 {
	if nf.Valid {
		return &nf.Float64 // Retorna o ponteiro para o valor de Float64
	}
	return nil // Retorna nil se o valor for NULL
}

// ToNullInt64 converts a pointer to an int64 to a sql.NullInt64.
func ToNullInt64(s *int64) sql.NullInt64 {
	if s == nil {
		return sql.NullInt64{Int64: 0, Valid: false}
	}
	return sql.NullInt64{Int64: *s, Valid: true}
}

// ToNullFloat64 converts a pointer to a float64 to a sql.NullFloat64.
func ToNullFloat64(f *float64) sql.NullFloat64 {
	if f == nil {
		return sql.NullFloat64{Float64: 0, Valid: false}
	}
	return sql.NullFloat64{Float64: *f, Valid: true}
}

// ToNullTime converts a pointer to a time.Time to a sql.NullTime.
func ToNullTime(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{Time: time.Time{}, Valid: false}
	}
	return sql.NullTime{Time: *t, Valid: true}
}

// ToNullString converts a pointer to a string to a sql.NullString.
func ToNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: *s, Valid: true}
}

// ToNullBool converts a pointer to a bool to a sql.NullBool.
func ToNullBool(b *bool) sql.NullBool {
	if b == nil {
		return sql.NullBool{Bool: false, Valid: false}
	}
	return sql.NullBool{Bool: *b, Valid: true}
}
