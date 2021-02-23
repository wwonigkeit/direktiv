// Code generated by entc, DO NOT EDIT.

package server

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/vorteil/direktiv/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// IP applies equality check predicate on the "ip" field. It's identical to IPEQ.
func IP(v string) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIP), v))
	})
}

// ExtIP applies equality check predicate on the "extIP" field. It's identical to ExtIPEQ.
func ExtIP(v string) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExtIP), v))
	})
}

// NatsPort applies equality check predicate on the "natsPort" field. It's identical to NatsPortEQ.
func NatsPort(v int) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNatsPort), v))
	})
}

// MemberPort applies equality check predicate on the "memberPort" field. It's identical to MemberPortEQ.
func MemberPort(v int) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMemberPort), v))
	})
}

// Added applies equality check predicate on the "added" field. It's identical to AddedEQ.
func Added(v time.Time) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdded), v))
	})
}

// IPEQ applies the EQ predicate on the "ip" field.
func IPEQ(v string) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIP), v))
	})
}

// IPNEQ applies the NEQ predicate on the "ip" field.
func IPNEQ(v string) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIP), v))
	})
}

// IPIn applies the In predicate on the "ip" field.
func IPIn(vs ...string) predicate.Server {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Server(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIP), v...))
	})
}

// IPNotIn applies the NotIn predicate on the "ip" field.
func IPNotIn(vs ...string) predicate.Server {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Server(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIP), v...))
	})
}

// IPGT applies the GT predicate on the "ip" field.
func IPGT(v string) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIP), v))
	})
}

// IPGTE applies the GTE predicate on the "ip" field.
func IPGTE(v string) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIP), v))
	})
}

// IPLT applies the LT predicate on the "ip" field.
func IPLT(v string) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIP), v))
	})
}

// IPLTE applies the LTE predicate on the "ip" field.
func IPLTE(v string) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIP), v))
	})
}

// IPContains applies the Contains predicate on the "ip" field.
func IPContains(v string) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIP), v))
	})
}

// IPHasPrefix applies the HasPrefix predicate on the "ip" field.
func IPHasPrefix(v string) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIP), v))
	})
}

// IPHasSuffix applies the HasSuffix predicate on the "ip" field.
func IPHasSuffix(v string) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIP), v))
	})
}

// IPEqualFold applies the EqualFold predicate on the "ip" field.
func IPEqualFold(v string) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIP), v))
	})
}

// IPContainsFold applies the ContainsFold predicate on the "ip" field.
func IPContainsFold(v string) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIP), v))
	})
}

// ExtIPEQ applies the EQ predicate on the "extIP" field.
func ExtIPEQ(v string) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExtIP), v))
	})
}

// ExtIPNEQ applies the NEQ predicate on the "extIP" field.
func ExtIPNEQ(v string) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExtIP), v))
	})
}

// ExtIPIn applies the In predicate on the "extIP" field.
func ExtIPIn(vs ...string) predicate.Server {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Server(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExtIP), v...))
	})
}

// ExtIPNotIn applies the NotIn predicate on the "extIP" field.
func ExtIPNotIn(vs ...string) predicate.Server {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Server(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExtIP), v...))
	})
}

// ExtIPGT applies the GT predicate on the "extIP" field.
func ExtIPGT(v string) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExtIP), v))
	})
}

// ExtIPGTE applies the GTE predicate on the "extIP" field.
func ExtIPGTE(v string) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExtIP), v))
	})
}

// ExtIPLT applies the LT predicate on the "extIP" field.
func ExtIPLT(v string) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExtIP), v))
	})
}

// ExtIPLTE applies the LTE predicate on the "extIP" field.
func ExtIPLTE(v string) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExtIP), v))
	})
}

// ExtIPContains applies the Contains predicate on the "extIP" field.
func ExtIPContains(v string) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldExtIP), v))
	})
}

// ExtIPHasPrefix applies the HasPrefix predicate on the "extIP" field.
func ExtIPHasPrefix(v string) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldExtIP), v))
	})
}

// ExtIPHasSuffix applies the HasSuffix predicate on the "extIP" field.
func ExtIPHasSuffix(v string) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldExtIP), v))
	})
}

// ExtIPEqualFold applies the EqualFold predicate on the "extIP" field.
func ExtIPEqualFold(v string) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldExtIP), v))
	})
}

// ExtIPContainsFold applies the ContainsFold predicate on the "extIP" field.
func ExtIPContainsFold(v string) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldExtIP), v))
	})
}

// NatsPortEQ applies the EQ predicate on the "natsPort" field.
func NatsPortEQ(v int) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNatsPort), v))
	})
}

// NatsPortNEQ applies the NEQ predicate on the "natsPort" field.
func NatsPortNEQ(v int) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNatsPort), v))
	})
}

// NatsPortIn applies the In predicate on the "natsPort" field.
func NatsPortIn(vs ...int) predicate.Server {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Server(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNatsPort), v...))
	})
}

// NatsPortNotIn applies the NotIn predicate on the "natsPort" field.
func NatsPortNotIn(vs ...int) predicate.Server {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Server(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNatsPort), v...))
	})
}

// NatsPortGT applies the GT predicate on the "natsPort" field.
func NatsPortGT(v int) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNatsPort), v))
	})
}

// NatsPortGTE applies the GTE predicate on the "natsPort" field.
func NatsPortGTE(v int) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNatsPort), v))
	})
}

// NatsPortLT applies the LT predicate on the "natsPort" field.
func NatsPortLT(v int) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNatsPort), v))
	})
}

// NatsPortLTE applies the LTE predicate on the "natsPort" field.
func NatsPortLTE(v int) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNatsPort), v))
	})
}

// MemberPortEQ applies the EQ predicate on the "memberPort" field.
func MemberPortEQ(v int) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMemberPort), v))
	})
}

// MemberPortNEQ applies the NEQ predicate on the "memberPort" field.
func MemberPortNEQ(v int) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMemberPort), v))
	})
}

// MemberPortIn applies the In predicate on the "memberPort" field.
func MemberPortIn(vs ...int) predicate.Server {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Server(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMemberPort), v...))
	})
}

// MemberPortNotIn applies the NotIn predicate on the "memberPort" field.
func MemberPortNotIn(vs ...int) predicate.Server {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Server(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMemberPort), v...))
	})
}

// MemberPortGT applies the GT predicate on the "memberPort" field.
func MemberPortGT(v int) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMemberPort), v))
	})
}

// MemberPortGTE applies the GTE predicate on the "memberPort" field.
func MemberPortGTE(v int) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMemberPort), v))
	})
}

// MemberPortLT applies the LT predicate on the "memberPort" field.
func MemberPortLT(v int) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMemberPort), v))
	})
}

// MemberPortLTE applies the LTE predicate on the "memberPort" field.
func MemberPortLTE(v int) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMemberPort), v))
	})
}

// AddedEQ applies the EQ predicate on the "added" field.
func AddedEQ(v time.Time) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdded), v))
	})
}

// AddedNEQ applies the NEQ predicate on the "added" field.
func AddedNEQ(v time.Time) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAdded), v))
	})
}

// AddedIn applies the In predicate on the "added" field.
func AddedIn(vs ...time.Time) predicate.Server {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Server(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAdded), v...))
	})
}

// AddedNotIn applies the NotIn predicate on the "added" field.
func AddedNotIn(vs ...time.Time) predicate.Server {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Server(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAdded), v...))
	})
}

// AddedGT applies the GT predicate on the "added" field.
func AddedGT(v time.Time) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAdded), v))
	})
}

// AddedGTE applies the GTE predicate on the "added" field.
func AddedGTE(v time.Time) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAdded), v))
	})
}

// AddedLT applies the LT predicate on the "added" field.
func AddedLT(v time.Time) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAdded), v))
	})
}

// AddedLTE applies the LTE predicate on the "added" field.
func AddedLTE(v time.Time) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAdded), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Server) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Server) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
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
func Not(p predicate.Server) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		p(s.Not())
	})
}