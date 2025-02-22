// Code generated by ent, DO NOT EDIT.

package attachment

import (
	"entgo.io/contrib/entproto/cmd/protoc-gen-ent/internal/todo/ent/predicate"
	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Contents applies equality check predicate on the "contents" field. It's identical to ContentsEQ.
func Contents(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContents), v))
	})
}

// ContentsEQ applies the EQ predicate on the "contents" field.
func ContentsEQ(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContents), v))
	})
}

// ContentsNEQ applies the NEQ predicate on the "contents" field.
func ContentsNEQ(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContents), v))
	})
}

// ContentsIn applies the In predicate on the "contents" field.
func ContentsIn(vs ...string) predicate.Attachment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldContents), v...))
	})
}

// ContentsNotIn applies the NotIn predicate on the "contents" field.
func ContentsNotIn(vs ...string) predicate.Attachment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldContents), v...))
	})
}

// ContentsGT applies the GT predicate on the "contents" field.
func ContentsGT(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContents), v))
	})
}

// ContentsGTE applies the GTE predicate on the "contents" field.
func ContentsGTE(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContents), v))
	})
}

// ContentsLT applies the LT predicate on the "contents" field.
func ContentsLT(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContents), v))
	})
}

// ContentsLTE applies the LTE predicate on the "contents" field.
func ContentsLTE(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContents), v))
	})
}

// ContentsContains applies the Contains predicate on the "contents" field.
func ContentsContains(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContents), v))
	})
}

// ContentsHasPrefix applies the HasPrefix predicate on the "contents" field.
func ContentsHasPrefix(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContents), v))
	})
}

// ContentsHasSuffix applies the HasSuffix predicate on the "contents" field.
func ContentsHasSuffix(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContents), v))
	})
}

// ContentsEqualFold applies the EqualFold predicate on the "contents" field.
func ContentsEqualFold(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContents), v))
	})
}

// ContentsContainsFold applies the ContainsFold predicate on the "contents" field.
func ContentsContainsFold(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContents), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Attachment) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Attachment) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
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
func Not(p predicate.Attachment) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		p(s.Not())
	})
}
