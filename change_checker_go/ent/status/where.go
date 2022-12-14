// Code generated by ent, DO NOT EDIT.

package status

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"gitlab.42paris.fr/utilities/notion_service/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// DbID applies equality check predicate on the "db_id" field. It's identical to DbIDEQ.
func DbID(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDbID), v))
	})
}

// RowID applies equality check predicate on the "row_id" field. It's identical to RowIDEQ.
func RowID(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRowID), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// SavedAt applies equality check predicate on the "saved_at" field. It's identical to SavedAtEQ.
func SavedAt(v time.Time) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSavedAt), v))
	})
}

// DbIDEQ applies the EQ predicate on the "db_id" field.
func DbIDEQ(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDbID), v))
	})
}

// DbIDNEQ applies the NEQ predicate on the "db_id" field.
func DbIDNEQ(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDbID), v))
	})
}

// DbIDIn applies the In predicate on the "db_id" field.
func DbIDIn(vs ...string) predicate.Status {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDbID), v...))
	})
}

// DbIDNotIn applies the NotIn predicate on the "db_id" field.
func DbIDNotIn(vs ...string) predicate.Status {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDbID), v...))
	})
}

// DbIDGT applies the GT predicate on the "db_id" field.
func DbIDGT(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDbID), v))
	})
}

// DbIDGTE applies the GTE predicate on the "db_id" field.
func DbIDGTE(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDbID), v))
	})
}

// DbIDLT applies the LT predicate on the "db_id" field.
func DbIDLT(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDbID), v))
	})
}

// DbIDLTE applies the LTE predicate on the "db_id" field.
func DbIDLTE(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDbID), v))
	})
}

// DbIDContains applies the Contains predicate on the "db_id" field.
func DbIDContains(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDbID), v))
	})
}

// DbIDHasPrefix applies the HasPrefix predicate on the "db_id" field.
func DbIDHasPrefix(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDbID), v))
	})
}

// DbIDHasSuffix applies the HasSuffix predicate on the "db_id" field.
func DbIDHasSuffix(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDbID), v))
	})
}

// DbIDEqualFold applies the EqualFold predicate on the "db_id" field.
func DbIDEqualFold(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDbID), v))
	})
}

// DbIDContainsFold applies the ContainsFold predicate on the "db_id" field.
func DbIDContainsFold(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDbID), v))
	})
}

// RowIDEQ applies the EQ predicate on the "row_id" field.
func RowIDEQ(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRowID), v))
	})
}

// RowIDNEQ applies the NEQ predicate on the "row_id" field.
func RowIDNEQ(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRowID), v))
	})
}

// RowIDIn applies the In predicate on the "row_id" field.
func RowIDIn(vs ...string) predicate.Status {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRowID), v...))
	})
}

// RowIDNotIn applies the NotIn predicate on the "row_id" field.
func RowIDNotIn(vs ...string) predicate.Status {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRowID), v...))
	})
}

// RowIDGT applies the GT predicate on the "row_id" field.
func RowIDGT(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRowID), v))
	})
}

// RowIDGTE applies the GTE predicate on the "row_id" field.
func RowIDGTE(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRowID), v))
	})
}

// RowIDLT applies the LT predicate on the "row_id" field.
func RowIDLT(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRowID), v))
	})
}

// RowIDLTE applies the LTE predicate on the "row_id" field.
func RowIDLTE(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRowID), v))
	})
}

// RowIDContains applies the Contains predicate on the "row_id" field.
func RowIDContains(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRowID), v))
	})
}

// RowIDHasPrefix applies the HasPrefix predicate on the "row_id" field.
func RowIDHasPrefix(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRowID), v))
	})
}

// RowIDHasSuffix applies the HasSuffix predicate on the "row_id" field.
func RowIDHasSuffix(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRowID), v))
	})
}

// RowIDEqualFold applies the EqualFold predicate on the "row_id" field.
func RowIDEqualFold(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRowID), v))
	})
}

// RowIDContainsFold applies the ContainsFold predicate on the "row_id" field.
func RowIDContainsFold(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRowID), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Status {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Status {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStatus), v))
	})
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStatus), v))
	})
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStatus), v))
	})
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStatus), v))
	})
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStatus), v))
	})
}

// SavedAtEQ applies the EQ predicate on the "saved_at" field.
func SavedAtEQ(v time.Time) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSavedAt), v))
	})
}

// SavedAtNEQ applies the NEQ predicate on the "saved_at" field.
func SavedAtNEQ(v time.Time) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSavedAt), v))
	})
}

// SavedAtIn applies the In predicate on the "saved_at" field.
func SavedAtIn(vs ...time.Time) predicate.Status {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSavedAt), v...))
	})
}

// SavedAtNotIn applies the NotIn predicate on the "saved_at" field.
func SavedAtNotIn(vs ...time.Time) predicate.Status {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSavedAt), v...))
	})
}

// SavedAtGT applies the GT predicate on the "saved_at" field.
func SavedAtGT(v time.Time) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSavedAt), v))
	})
}

// SavedAtGTE applies the GTE predicate on the "saved_at" field.
func SavedAtGTE(v time.Time) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSavedAt), v))
	})
}

// SavedAtLT applies the LT predicate on the "saved_at" field.
func SavedAtLT(v time.Time) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSavedAt), v))
	})
}

// SavedAtLTE applies the LTE predicate on the "saved_at" field.
func SavedAtLTE(v time.Time) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSavedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Status) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Status) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Status) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		p(s.Not())
	})
}
