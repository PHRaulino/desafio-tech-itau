package conversores

import (
	"database/sql"
	"time"
)

// Conversores de Null para ponteiro

func NullStringParaPonteiro(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

func NullBoolParaPonteiro(nb sql.NullBool) *bool {
	if nb.Valid {
		return &nb.Bool
	}
	return nil
}

func NullFloat64ParaPonteiro(nf sql.NullFloat64) *float64 {
	if nf.Valid {
		return &nf.Float64
	}
	return nil
}

// Conversores de ponteiro para Null

func ParaNullInt64(i *int64) sql.NullInt64 {
	if i == nil {
		return sql.NullInt64{Int64: 0, Valid: false}
	}
	return sql.NullInt64{Int64: *i, Valid: true}
}

func ParaNullFloat64(f *float64) sql.NullFloat64 {
	if f == nil {
		return sql.NullFloat64{Float64: 0, Valid: false}
	}
	return sql.NullFloat64{Float64: *f, Valid: true}
}

func ParaNullTime(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{Time: time.Time{}, Valid: false}
	}
	return sql.NullTime{Time: *t, Valid: true}
}

func ParaNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: *s, Valid: true}
}

func ParaNullBool(b *bool) sql.NullBool {
	if b == nil {
		return sql.NullBool{Bool: false, Valid: false}
	}
	return sql.NullBool{Bool: *b, Valid: true}
}
