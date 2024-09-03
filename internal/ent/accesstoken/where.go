// Code generated by ent, DO NOT EDIT.

package accesstoken

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/duc-cnzj/mars/v5/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldDeletedAt, v))
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldToken, v))
}

// Usage applies equality check predicate on the "usage" field. It's identical to UsageEQ.
func Usage(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldUsage, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldEmail, v))
}

// ExpiredAt applies equality check predicate on the "expired_at" field. It's identical to ExpiredAtEQ.
func ExpiredAt(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldExpiredAt, v))
}

// LastUsedAt applies equality check predicate on the "last_used_at" field. It's identical to LastUsedAtEQ.
func LastUsedAt(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldLastUsedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotNull(FieldDeletedAt))
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldToken, v))
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldToken, v))
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldToken, vs...))
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldToken, vs...))
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldToken, v))
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldToken, v))
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldToken, v))
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldToken, v))
}

// TokenContains applies the Contains predicate on the "token" field.
func TokenContains(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContains(FieldToken, v))
}

// TokenHasPrefix applies the HasPrefix predicate on the "token" field.
func TokenHasPrefix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasPrefix(FieldToken, v))
}

// TokenHasSuffix applies the HasSuffix predicate on the "token" field.
func TokenHasSuffix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasSuffix(FieldToken, v))
}

// TokenEqualFold applies the EqualFold predicate on the "token" field.
func TokenEqualFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEqualFold(FieldToken, v))
}

// TokenContainsFold applies the ContainsFold predicate on the "token" field.
func TokenContainsFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContainsFold(FieldToken, v))
}

// UsageEQ applies the EQ predicate on the "usage" field.
func UsageEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldUsage, v))
}

// UsageNEQ applies the NEQ predicate on the "usage" field.
func UsageNEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldUsage, v))
}

// UsageIn applies the In predicate on the "usage" field.
func UsageIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldUsage, vs...))
}

// UsageNotIn applies the NotIn predicate on the "usage" field.
func UsageNotIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldUsage, vs...))
}

// UsageGT applies the GT predicate on the "usage" field.
func UsageGT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldUsage, v))
}

// UsageGTE applies the GTE predicate on the "usage" field.
func UsageGTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldUsage, v))
}

// UsageLT applies the LT predicate on the "usage" field.
func UsageLT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldUsage, v))
}

// UsageLTE applies the LTE predicate on the "usage" field.
func UsageLTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldUsage, v))
}

// UsageContains applies the Contains predicate on the "usage" field.
func UsageContains(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContains(FieldUsage, v))
}

// UsageHasPrefix applies the HasPrefix predicate on the "usage" field.
func UsageHasPrefix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasPrefix(FieldUsage, v))
}

// UsageHasSuffix applies the HasSuffix predicate on the "usage" field.
func UsageHasSuffix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasSuffix(FieldUsage, v))
}

// UsageEqualFold applies the EqualFold predicate on the "usage" field.
func UsageEqualFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEqualFold(FieldUsage, v))
}

// UsageContainsFold applies the ContainsFold predicate on the "usage" field.
func UsageContainsFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContainsFold(FieldUsage, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContainsFold(FieldEmail, v))
}

// ExpiredAtEQ applies the EQ predicate on the "expired_at" field.
func ExpiredAtEQ(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldExpiredAt, v))
}

// ExpiredAtNEQ applies the NEQ predicate on the "expired_at" field.
func ExpiredAtNEQ(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldExpiredAt, v))
}

// ExpiredAtIn applies the In predicate on the "expired_at" field.
func ExpiredAtIn(vs ...time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldExpiredAt, vs...))
}

// ExpiredAtNotIn applies the NotIn predicate on the "expired_at" field.
func ExpiredAtNotIn(vs ...time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldExpiredAt, vs...))
}

// ExpiredAtGT applies the GT predicate on the "expired_at" field.
func ExpiredAtGT(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldExpiredAt, v))
}

// ExpiredAtGTE applies the GTE predicate on the "expired_at" field.
func ExpiredAtGTE(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldExpiredAt, v))
}

// ExpiredAtLT applies the LT predicate on the "expired_at" field.
func ExpiredAtLT(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldExpiredAt, v))
}

// ExpiredAtLTE applies the LTE predicate on the "expired_at" field.
func ExpiredAtLTE(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldExpiredAt, v))
}

// ExpiredAtIsNil applies the IsNil predicate on the "expired_at" field.
func ExpiredAtIsNil() predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIsNull(FieldExpiredAt))
}

// ExpiredAtNotNil applies the NotNil predicate on the "expired_at" field.
func ExpiredAtNotNil() predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotNull(FieldExpiredAt))
}

// LastUsedAtEQ applies the EQ predicate on the "last_used_at" field.
func LastUsedAtEQ(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldLastUsedAt, v))
}

// LastUsedAtNEQ applies the NEQ predicate on the "last_used_at" field.
func LastUsedAtNEQ(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldLastUsedAt, v))
}

// LastUsedAtIn applies the In predicate on the "last_used_at" field.
func LastUsedAtIn(vs ...time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldLastUsedAt, vs...))
}

// LastUsedAtNotIn applies the NotIn predicate on the "last_used_at" field.
func LastUsedAtNotIn(vs ...time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldLastUsedAt, vs...))
}

// LastUsedAtGT applies the GT predicate on the "last_used_at" field.
func LastUsedAtGT(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldLastUsedAt, v))
}

// LastUsedAtGTE applies the GTE predicate on the "last_used_at" field.
func LastUsedAtGTE(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldLastUsedAt, v))
}

// LastUsedAtLT applies the LT predicate on the "last_used_at" field.
func LastUsedAtLT(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldLastUsedAt, v))
}

// LastUsedAtLTE applies the LTE predicate on the "last_used_at" field.
func LastUsedAtLTE(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldLastUsedAt, v))
}

// LastUsedAtIsNil applies the IsNil predicate on the "last_used_at" field.
func LastUsedAtIsNil() predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIsNull(FieldLastUsedAt))
}

// LastUsedAtNotNil applies the NotNil predicate on the "last_used_at" field.
func LastUsedAtNotNil() predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotNull(FieldLastUsedAt))
}

// UserInfoIsNil applies the IsNil predicate on the "user_info" field.
func UserInfoIsNil() predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIsNull(FieldUserInfo))
}

// UserInfoNotNil applies the NotNil predicate on the "user_info" field.
func UserInfoNotNil() predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotNull(FieldUserInfo))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AccessToken) predicate.AccessToken {
	return predicate.AccessToken(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AccessToken) predicate.AccessToken {
	return predicate.AccessToken(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AccessToken) predicate.AccessToken {
	return predicate.AccessToken(sql.NotPredicates(p))
}