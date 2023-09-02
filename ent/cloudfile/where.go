// Code generated by ent, DO NOT EDIT.

package cloudfile

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-file/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldEQ(FieldUpdatedAt, v))
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v bool) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldEQ(FieldState, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldEQ(FieldName, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldEQ(FieldURL, v))
}

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v uint64) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldEQ(FieldSize, v))
}

// FileType applies equality check predicate on the "file_type" field. It's identical to FileTypeEQ.
func FileType(v uint8) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldEQ(FieldFileType, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldEQ(FieldUserID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldLTE(FieldUpdatedAt, v))
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v bool) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldEQ(FieldState, v))
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v bool) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldNEQ(FieldState, v))
}

// StateIsNil applies the IsNil predicate on the "state" field.
func StateIsNil() predicate.CloudFile {
	return predicate.CloudFile(sql.FieldIsNull(FieldState))
}

// StateNotNil applies the NotNil predicate on the "state" field.
func StateNotNil() predicate.CloudFile {
	return predicate.CloudFile(sql.FieldNotNull(FieldState))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldContainsFold(FieldName, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldContainsFold(FieldURL, v))
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v uint64) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldEQ(FieldSize, v))
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v uint64) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldNEQ(FieldSize, v))
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...uint64) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldIn(FieldSize, vs...))
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...uint64) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldNotIn(FieldSize, vs...))
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v uint64) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldGT(FieldSize, v))
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v uint64) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldGTE(FieldSize, v))
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v uint64) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldLT(FieldSize, v))
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v uint64) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldLTE(FieldSize, v))
}

// FileTypeEQ applies the EQ predicate on the "file_type" field.
func FileTypeEQ(v uint8) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldEQ(FieldFileType, v))
}

// FileTypeNEQ applies the NEQ predicate on the "file_type" field.
func FileTypeNEQ(v uint8) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldNEQ(FieldFileType, v))
}

// FileTypeIn applies the In predicate on the "file_type" field.
func FileTypeIn(vs ...uint8) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldIn(FieldFileType, vs...))
}

// FileTypeNotIn applies the NotIn predicate on the "file_type" field.
func FileTypeNotIn(vs ...uint8) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldNotIn(FieldFileType, vs...))
}

// FileTypeGT applies the GT predicate on the "file_type" field.
func FileTypeGT(v uint8) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldGT(FieldFileType, v))
}

// FileTypeGTE applies the GTE predicate on the "file_type" field.
func FileTypeGTE(v uint8) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldGTE(FieldFileType, v))
}

// FileTypeLT applies the LT predicate on the "file_type" field.
func FileTypeLT(v uint8) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldLT(FieldFileType, v))
}

// FileTypeLTE applies the LTE predicate on the "file_type" field.
func FileTypeLTE(v uint8) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldLTE(FieldFileType, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.CloudFile {
	return predicate.CloudFile(sql.FieldContainsFold(FieldUserID, v))
}

// HasStorageProviders applies the HasEdge predicate on the "storage_providers" edge.
func HasStorageProviders() predicate.CloudFile {
	return predicate.CloudFile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, StorageProvidersTable, StorageProvidersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStorageProvidersWith applies the HasEdge predicate on the "storage_providers" edge with a given conditions (other predicates).
func HasStorageProvidersWith(preds ...predicate.StorageProvider) predicate.CloudFile {
	return predicate.CloudFile(func(s *sql.Selector) {
		step := newStorageProvidersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CloudFile) predicate.CloudFile {
	return predicate.CloudFile(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CloudFile) predicate.CloudFile {
	return predicate.CloudFile(func(s *sql.Selector) {
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
func Not(p predicate.CloudFile) predicate.CloudFile {
	return predicate.CloudFile(func(s *sql.Selector) {
		p(s.Not())
	})
}
