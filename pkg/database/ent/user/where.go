// Code generated by entc, DO NOT EDIT.

package user

import (
	"llamalauncher/pkg/database/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func IDGT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// AccountID applies equality check predicate on the "account_id" field. It's identical to AccountIDEQ.
func AccountID(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccountID), v))
	})
}

// DisplayName applies equality check predicate on the "display_name" field. It's identical to DisplayNameEQ.
func DisplayName(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDisplayName), v))
	})
}

// AccessToken applies equality check predicate on the "access_token" field. It's identical to AccessTokenEQ.
func AccessToken(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccessToken), v))
	})
}

// AccessTokenExpiresAt applies equality check predicate on the "access_token_expires_at" field. It's identical to AccessTokenExpiresAtEQ.
func AccessTokenExpiresAt(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccessTokenExpiresAt), v))
	})
}

// RefreshToken applies equality check predicate on the "refresh_token" field. It's identical to RefreshTokenEQ.
func RefreshToken(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRefreshToken), v))
	})
}

// RefresTokenExpiresAt applies equality check predicate on the "refres_token_expires_at" field. It's identical to RefresTokenExpiresAtEQ.
func RefresTokenExpiresAt(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRefresTokenExpiresAt), v))
	})
}

// AccountIDEQ applies the EQ predicate on the "account_id" field.
func AccountIDEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccountID), v))
	})
}

// AccountIDNEQ applies the NEQ predicate on the "account_id" field.
func AccountIDNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAccountID), v))
	})
}

// AccountIDIn applies the In predicate on the "account_id" field.
func AccountIDIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAccountID), v...))
	})
}

// AccountIDNotIn applies the NotIn predicate on the "account_id" field.
func AccountIDNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAccountID), v...))
	})
}

// AccountIDGT applies the GT predicate on the "account_id" field.
func AccountIDGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAccountID), v))
	})
}

// AccountIDGTE applies the GTE predicate on the "account_id" field.
func AccountIDGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAccountID), v))
	})
}

// AccountIDLT applies the LT predicate on the "account_id" field.
func AccountIDLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAccountID), v))
	})
}

// AccountIDLTE applies the LTE predicate on the "account_id" field.
func AccountIDLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAccountID), v))
	})
}

// AccountIDContains applies the Contains predicate on the "account_id" field.
func AccountIDContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAccountID), v))
	})
}

// AccountIDHasPrefix applies the HasPrefix predicate on the "account_id" field.
func AccountIDHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAccountID), v))
	})
}

// AccountIDHasSuffix applies the HasSuffix predicate on the "account_id" field.
func AccountIDHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAccountID), v))
	})
}

// AccountIDEqualFold applies the EqualFold predicate on the "account_id" field.
func AccountIDEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAccountID), v))
	})
}

// AccountIDContainsFold applies the ContainsFold predicate on the "account_id" field.
func AccountIDContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAccountID), v))
	})
}

// DisplayNameEQ applies the EQ predicate on the "display_name" field.
func DisplayNameEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDisplayName), v))
	})
}

// DisplayNameNEQ applies the NEQ predicate on the "display_name" field.
func DisplayNameNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDisplayName), v))
	})
}

// DisplayNameIn applies the In predicate on the "display_name" field.
func DisplayNameIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDisplayName), v...))
	})
}

// DisplayNameNotIn applies the NotIn predicate on the "display_name" field.
func DisplayNameNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDisplayName), v...))
	})
}

// DisplayNameGT applies the GT predicate on the "display_name" field.
func DisplayNameGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDisplayName), v))
	})
}

// DisplayNameGTE applies the GTE predicate on the "display_name" field.
func DisplayNameGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDisplayName), v))
	})
}

// DisplayNameLT applies the LT predicate on the "display_name" field.
func DisplayNameLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDisplayName), v))
	})
}

// DisplayNameLTE applies the LTE predicate on the "display_name" field.
func DisplayNameLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDisplayName), v))
	})
}

// DisplayNameContains applies the Contains predicate on the "display_name" field.
func DisplayNameContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDisplayName), v))
	})
}

// DisplayNameHasPrefix applies the HasPrefix predicate on the "display_name" field.
func DisplayNameHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDisplayName), v))
	})
}

// DisplayNameHasSuffix applies the HasSuffix predicate on the "display_name" field.
func DisplayNameHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDisplayName), v))
	})
}

// DisplayNameEqualFold applies the EqualFold predicate on the "display_name" field.
func DisplayNameEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDisplayName), v))
	})
}

// DisplayNameContainsFold applies the ContainsFold predicate on the "display_name" field.
func DisplayNameContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDisplayName), v))
	})
}

// AccessTokenEQ applies the EQ predicate on the "access_token" field.
func AccessTokenEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccessToken), v))
	})
}

// AccessTokenNEQ applies the NEQ predicate on the "access_token" field.
func AccessTokenNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAccessToken), v))
	})
}

// AccessTokenIn applies the In predicate on the "access_token" field.
func AccessTokenIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAccessToken), v...))
	})
}

// AccessTokenNotIn applies the NotIn predicate on the "access_token" field.
func AccessTokenNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAccessToken), v...))
	})
}

// AccessTokenGT applies the GT predicate on the "access_token" field.
func AccessTokenGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAccessToken), v))
	})
}

// AccessTokenGTE applies the GTE predicate on the "access_token" field.
func AccessTokenGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAccessToken), v))
	})
}

// AccessTokenLT applies the LT predicate on the "access_token" field.
func AccessTokenLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAccessToken), v))
	})
}

// AccessTokenLTE applies the LTE predicate on the "access_token" field.
func AccessTokenLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAccessToken), v))
	})
}

// AccessTokenContains applies the Contains predicate on the "access_token" field.
func AccessTokenContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAccessToken), v))
	})
}

// AccessTokenHasPrefix applies the HasPrefix predicate on the "access_token" field.
func AccessTokenHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAccessToken), v))
	})
}

// AccessTokenHasSuffix applies the HasSuffix predicate on the "access_token" field.
func AccessTokenHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAccessToken), v))
	})
}

// AccessTokenEqualFold applies the EqualFold predicate on the "access_token" field.
func AccessTokenEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAccessToken), v))
	})
}

// AccessTokenContainsFold applies the ContainsFold predicate on the "access_token" field.
func AccessTokenContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAccessToken), v))
	})
}

// AccessTokenExpiresAtEQ applies the EQ predicate on the "access_token_expires_at" field.
func AccessTokenExpiresAtEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccessTokenExpiresAt), v))
	})
}

// AccessTokenExpiresAtNEQ applies the NEQ predicate on the "access_token_expires_at" field.
func AccessTokenExpiresAtNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAccessTokenExpiresAt), v))
	})
}

// AccessTokenExpiresAtIn applies the In predicate on the "access_token_expires_at" field.
func AccessTokenExpiresAtIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAccessTokenExpiresAt), v...))
	})
}

// AccessTokenExpiresAtNotIn applies the NotIn predicate on the "access_token_expires_at" field.
func AccessTokenExpiresAtNotIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAccessTokenExpiresAt), v...))
	})
}

// AccessTokenExpiresAtGT applies the GT predicate on the "access_token_expires_at" field.
func AccessTokenExpiresAtGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAccessTokenExpiresAt), v))
	})
}

// AccessTokenExpiresAtGTE applies the GTE predicate on the "access_token_expires_at" field.
func AccessTokenExpiresAtGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAccessTokenExpiresAt), v))
	})
}

// AccessTokenExpiresAtLT applies the LT predicate on the "access_token_expires_at" field.
func AccessTokenExpiresAtLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAccessTokenExpiresAt), v))
	})
}

// AccessTokenExpiresAtLTE applies the LTE predicate on the "access_token_expires_at" field.
func AccessTokenExpiresAtLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAccessTokenExpiresAt), v))
	})
}

// RefreshTokenEQ applies the EQ predicate on the "refresh_token" field.
func RefreshTokenEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenNEQ applies the NEQ predicate on the "refresh_token" field.
func RefreshTokenNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenIn applies the In predicate on the "refresh_token" field.
func RefreshTokenIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRefreshToken), v...))
	})
}

// RefreshTokenNotIn applies the NotIn predicate on the "refresh_token" field.
func RefreshTokenNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRefreshToken), v...))
	})
}

// RefreshTokenGT applies the GT predicate on the "refresh_token" field.
func RefreshTokenGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenGTE applies the GTE predicate on the "refresh_token" field.
func RefreshTokenGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenLT applies the LT predicate on the "refresh_token" field.
func RefreshTokenLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenLTE applies the LTE predicate on the "refresh_token" field.
func RefreshTokenLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenContains applies the Contains predicate on the "refresh_token" field.
func RefreshTokenContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenHasPrefix applies the HasPrefix predicate on the "refresh_token" field.
func RefreshTokenHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenHasSuffix applies the HasSuffix predicate on the "refresh_token" field.
func RefreshTokenHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenEqualFold applies the EqualFold predicate on the "refresh_token" field.
func RefreshTokenEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenContainsFold applies the ContainsFold predicate on the "refresh_token" field.
func RefreshTokenContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRefreshToken), v))
	})
}

// RefresTokenExpiresAtEQ applies the EQ predicate on the "refres_token_expires_at" field.
func RefresTokenExpiresAtEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRefresTokenExpiresAt), v))
	})
}

// RefresTokenExpiresAtNEQ applies the NEQ predicate on the "refres_token_expires_at" field.
func RefresTokenExpiresAtNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRefresTokenExpiresAt), v))
	})
}

// RefresTokenExpiresAtIn applies the In predicate on the "refres_token_expires_at" field.
func RefresTokenExpiresAtIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRefresTokenExpiresAt), v...))
	})
}

// RefresTokenExpiresAtNotIn applies the NotIn predicate on the "refres_token_expires_at" field.
func RefresTokenExpiresAtNotIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRefresTokenExpiresAt), v...))
	})
}

// RefresTokenExpiresAtGT applies the GT predicate on the "refres_token_expires_at" field.
func RefresTokenExpiresAtGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRefresTokenExpiresAt), v))
	})
}

// RefresTokenExpiresAtGTE applies the GTE predicate on the "refres_token_expires_at" field.
func RefresTokenExpiresAtGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRefresTokenExpiresAt), v))
	})
}

// RefresTokenExpiresAtLT applies the LT predicate on the "refres_token_expires_at" field.
func RefresTokenExpiresAtLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRefresTokenExpiresAt), v))
	})
}

// RefresTokenExpiresAtLTE applies the LTE predicate on the "refres_token_expires_at" field.
func RefresTokenExpiresAtLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRefresTokenExpiresAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
