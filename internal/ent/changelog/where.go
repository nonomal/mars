// Code generated by ent, DO NOT EDIT.

package changelog

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/duc-cnzj/mars/v5/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Changelog {
	return predicate.Changelog(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Changelog {
	return predicate.Changelog(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Changelog {
	return predicate.Changelog(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Changelog {
	return predicate.Changelog(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Changelog {
	return predicate.Changelog(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Changelog {
	return predicate.Changelog(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Changelog {
	return predicate.Changelog(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldDeletedAt, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v int) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldVersion, v))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldUsername, v))
}

// Config applies equality check predicate on the "config" field. It's identical to ConfigEQ.
func Config(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldConfig, v))
}

// GitBranch applies equality check predicate on the "git_branch" field. It's identical to GitBranchEQ.
func GitBranch(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldGitBranch, v))
}

// GitCommit applies equality check predicate on the "git_commit" field. It's identical to GitCommitEQ.
func GitCommit(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldGitCommit, v))
}

// GitCommitWebURL applies equality check predicate on the "git_commit_web_url" field. It's identical to GitCommitWebURLEQ.
func GitCommitWebURL(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldGitCommitWebURL, v))
}

// GitCommitTitle applies equality check predicate on the "git_commit_title" field. It's identical to GitCommitTitleEQ.
func GitCommitTitle(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldGitCommitTitle, v))
}

// GitCommitAuthor applies equality check predicate on the "git_commit_author" field. It's identical to GitCommitAuthorEQ.
func GitCommitAuthor(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldGitCommitAuthor, v))
}

// GitCommitDate applies equality check predicate on the "git_commit_date" field. It's identical to GitCommitDateEQ.
func GitCommitDate(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldGitCommitDate, v))
}

// ConfigChanged applies equality check predicate on the "config_changed" field. It's identical to ConfigChangedEQ.
func ConfigChanged(v bool) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldConfigChanged, v))
}

// ProjectID applies equality check predicate on the "project_id" field. It's identical to ProjectIDEQ.
func ProjectID(v int) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldProjectID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Changelog {
	return predicate.Changelog(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Changelog {
	return predicate.Changelog(sql.FieldNotNull(FieldDeletedAt))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v int) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v int) predicate.Changelog {
	return predicate.Changelog(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...int) predicate.Changelog {
	return predicate.Changelog(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...int) predicate.Changelog {
	return predicate.Changelog(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v int) predicate.Changelog {
	return predicate.Changelog(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v int) predicate.Changelog {
	return predicate.Changelog(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v int) predicate.Changelog {
	return predicate.Changelog(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v int) predicate.Changelog {
	return predicate.Changelog(sql.FieldLTE(FieldVersion, v))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.Changelog {
	return predicate.Changelog(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.Changelog {
	return predicate.Changelog(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldContainsFold(FieldUsername, v))
}

// ConfigEQ applies the EQ predicate on the "config" field.
func ConfigEQ(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldConfig, v))
}

// ConfigNEQ applies the NEQ predicate on the "config" field.
func ConfigNEQ(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldNEQ(FieldConfig, v))
}

// ConfigIn applies the In predicate on the "config" field.
func ConfigIn(vs ...string) predicate.Changelog {
	return predicate.Changelog(sql.FieldIn(FieldConfig, vs...))
}

// ConfigNotIn applies the NotIn predicate on the "config" field.
func ConfigNotIn(vs ...string) predicate.Changelog {
	return predicate.Changelog(sql.FieldNotIn(FieldConfig, vs...))
}

// ConfigGT applies the GT predicate on the "config" field.
func ConfigGT(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldGT(FieldConfig, v))
}

// ConfigGTE applies the GTE predicate on the "config" field.
func ConfigGTE(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldGTE(FieldConfig, v))
}

// ConfigLT applies the LT predicate on the "config" field.
func ConfigLT(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldLT(FieldConfig, v))
}

// ConfigLTE applies the LTE predicate on the "config" field.
func ConfigLTE(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldLTE(FieldConfig, v))
}

// ConfigContains applies the Contains predicate on the "config" field.
func ConfigContains(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldContains(FieldConfig, v))
}

// ConfigHasPrefix applies the HasPrefix predicate on the "config" field.
func ConfigHasPrefix(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldHasPrefix(FieldConfig, v))
}

// ConfigHasSuffix applies the HasSuffix predicate on the "config" field.
func ConfigHasSuffix(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldHasSuffix(FieldConfig, v))
}

// ConfigIsNil applies the IsNil predicate on the "config" field.
func ConfigIsNil() predicate.Changelog {
	return predicate.Changelog(sql.FieldIsNull(FieldConfig))
}

// ConfigNotNil applies the NotNil predicate on the "config" field.
func ConfigNotNil() predicate.Changelog {
	return predicate.Changelog(sql.FieldNotNull(FieldConfig))
}

// ConfigEqualFold applies the EqualFold predicate on the "config" field.
func ConfigEqualFold(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldEqualFold(FieldConfig, v))
}

// ConfigContainsFold applies the ContainsFold predicate on the "config" field.
func ConfigContainsFold(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldContainsFold(FieldConfig, v))
}

// GitBranchEQ applies the EQ predicate on the "git_branch" field.
func GitBranchEQ(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldGitBranch, v))
}

// GitBranchNEQ applies the NEQ predicate on the "git_branch" field.
func GitBranchNEQ(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldNEQ(FieldGitBranch, v))
}

// GitBranchIn applies the In predicate on the "git_branch" field.
func GitBranchIn(vs ...string) predicate.Changelog {
	return predicate.Changelog(sql.FieldIn(FieldGitBranch, vs...))
}

// GitBranchNotIn applies the NotIn predicate on the "git_branch" field.
func GitBranchNotIn(vs ...string) predicate.Changelog {
	return predicate.Changelog(sql.FieldNotIn(FieldGitBranch, vs...))
}

// GitBranchGT applies the GT predicate on the "git_branch" field.
func GitBranchGT(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldGT(FieldGitBranch, v))
}

// GitBranchGTE applies the GTE predicate on the "git_branch" field.
func GitBranchGTE(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldGTE(FieldGitBranch, v))
}

// GitBranchLT applies the LT predicate on the "git_branch" field.
func GitBranchLT(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldLT(FieldGitBranch, v))
}

// GitBranchLTE applies the LTE predicate on the "git_branch" field.
func GitBranchLTE(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldLTE(FieldGitBranch, v))
}

// GitBranchContains applies the Contains predicate on the "git_branch" field.
func GitBranchContains(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldContains(FieldGitBranch, v))
}

// GitBranchHasPrefix applies the HasPrefix predicate on the "git_branch" field.
func GitBranchHasPrefix(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldHasPrefix(FieldGitBranch, v))
}

// GitBranchHasSuffix applies the HasSuffix predicate on the "git_branch" field.
func GitBranchHasSuffix(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldHasSuffix(FieldGitBranch, v))
}

// GitBranchIsNil applies the IsNil predicate on the "git_branch" field.
func GitBranchIsNil() predicate.Changelog {
	return predicate.Changelog(sql.FieldIsNull(FieldGitBranch))
}

// GitBranchNotNil applies the NotNil predicate on the "git_branch" field.
func GitBranchNotNil() predicate.Changelog {
	return predicate.Changelog(sql.FieldNotNull(FieldGitBranch))
}

// GitBranchEqualFold applies the EqualFold predicate on the "git_branch" field.
func GitBranchEqualFold(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldEqualFold(FieldGitBranch, v))
}

// GitBranchContainsFold applies the ContainsFold predicate on the "git_branch" field.
func GitBranchContainsFold(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldContainsFold(FieldGitBranch, v))
}

// GitCommitEQ applies the EQ predicate on the "git_commit" field.
func GitCommitEQ(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldGitCommit, v))
}

// GitCommitNEQ applies the NEQ predicate on the "git_commit" field.
func GitCommitNEQ(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldNEQ(FieldGitCommit, v))
}

// GitCommitIn applies the In predicate on the "git_commit" field.
func GitCommitIn(vs ...string) predicate.Changelog {
	return predicate.Changelog(sql.FieldIn(FieldGitCommit, vs...))
}

// GitCommitNotIn applies the NotIn predicate on the "git_commit" field.
func GitCommitNotIn(vs ...string) predicate.Changelog {
	return predicate.Changelog(sql.FieldNotIn(FieldGitCommit, vs...))
}

// GitCommitGT applies the GT predicate on the "git_commit" field.
func GitCommitGT(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldGT(FieldGitCommit, v))
}

// GitCommitGTE applies the GTE predicate on the "git_commit" field.
func GitCommitGTE(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldGTE(FieldGitCommit, v))
}

// GitCommitLT applies the LT predicate on the "git_commit" field.
func GitCommitLT(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldLT(FieldGitCommit, v))
}

// GitCommitLTE applies the LTE predicate on the "git_commit" field.
func GitCommitLTE(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldLTE(FieldGitCommit, v))
}

// GitCommitContains applies the Contains predicate on the "git_commit" field.
func GitCommitContains(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldContains(FieldGitCommit, v))
}

// GitCommitHasPrefix applies the HasPrefix predicate on the "git_commit" field.
func GitCommitHasPrefix(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldHasPrefix(FieldGitCommit, v))
}

// GitCommitHasSuffix applies the HasSuffix predicate on the "git_commit" field.
func GitCommitHasSuffix(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldHasSuffix(FieldGitCommit, v))
}

// GitCommitIsNil applies the IsNil predicate on the "git_commit" field.
func GitCommitIsNil() predicate.Changelog {
	return predicate.Changelog(sql.FieldIsNull(FieldGitCommit))
}

// GitCommitNotNil applies the NotNil predicate on the "git_commit" field.
func GitCommitNotNil() predicate.Changelog {
	return predicate.Changelog(sql.FieldNotNull(FieldGitCommit))
}

// GitCommitEqualFold applies the EqualFold predicate on the "git_commit" field.
func GitCommitEqualFold(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldEqualFold(FieldGitCommit, v))
}

// GitCommitContainsFold applies the ContainsFold predicate on the "git_commit" field.
func GitCommitContainsFold(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldContainsFold(FieldGitCommit, v))
}

// DockerImageIsNil applies the IsNil predicate on the "docker_image" field.
func DockerImageIsNil() predicate.Changelog {
	return predicate.Changelog(sql.FieldIsNull(FieldDockerImage))
}

// DockerImageNotNil applies the NotNil predicate on the "docker_image" field.
func DockerImageNotNil() predicate.Changelog {
	return predicate.Changelog(sql.FieldNotNull(FieldDockerImage))
}

// EnvValuesIsNil applies the IsNil predicate on the "env_values" field.
func EnvValuesIsNil() predicate.Changelog {
	return predicate.Changelog(sql.FieldIsNull(FieldEnvValues))
}

// EnvValuesNotNil applies the NotNil predicate on the "env_values" field.
func EnvValuesNotNil() predicate.Changelog {
	return predicate.Changelog(sql.FieldNotNull(FieldEnvValues))
}

// ExtraValuesIsNil applies the IsNil predicate on the "extra_values" field.
func ExtraValuesIsNil() predicate.Changelog {
	return predicate.Changelog(sql.FieldIsNull(FieldExtraValues))
}

// ExtraValuesNotNil applies the NotNil predicate on the "extra_values" field.
func ExtraValuesNotNil() predicate.Changelog {
	return predicate.Changelog(sql.FieldNotNull(FieldExtraValues))
}

// FinalExtraValuesIsNil applies the IsNil predicate on the "final_extra_values" field.
func FinalExtraValuesIsNil() predicate.Changelog {
	return predicate.Changelog(sql.FieldIsNull(FieldFinalExtraValues))
}

// FinalExtraValuesNotNil applies the NotNil predicate on the "final_extra_values" field.
func FinalExtraValuesNotNil() predicate.Changelog {
	return predicate.Changelog(sql.FieldNotNull(FieldFinalExtraValues))
}

// GitCommitWebURLEQ applies the EQ predicate on the "git_commit_web_url" field.
func GitCommitWebURLEQ(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldGitCommitWebURL, v))
}

// GitCommitWebURLNEQ applies the NEQ predicate on the "git_commit_web_url" field.
func GitCommitWebURLNEQ(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldNEQ(FieldGitCommitWebURL, v))
}

// GitCommitWebURLIn applies the In predicate on the "git_commit_web_url" field.
func GitCommitWebURLIn(vs ...string) predicate.Changelog {
	return predicate.Changelog(sql.FieldIn(FieldGitCommitWebURL, vs...))
}

// GitCommitWebURLNotIn applies the NotIn predicate on the "git_commit_web_url" field.
func GitCommitWebURLNotIn(vs ...string) predicate.Changelog {
	return predicate.Changelog(sql.FieldNotIn(FieldGitCommitWebURL, vs...))
}

// GitCommitWebURLGT applies the GT predicate on the "git_commit_web_url" field.
func GitCommitWebURLGT(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldGT(FieldGitCommitWebURL, v))
}

// GitCommitWebURLGTE applies the GTE predicate on the "git_commit_web_url" field.
func GitCommitWebURLGTE(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldGTE(FieldGitCommitWebURL, v))
}

// GitCommitWebURLLT applies the LT predicate on the "git_commit_web_url" field.
func GitCommitWebURLLT(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldLT(FieldGitCommitWebURL, v))
}

// GitCommitWebURLLTE applies the LTE predicate on the "git_commit_web_url" field.
func GitCommitWebURLLTE(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldLTE(FieldGitCommitWebURL, v))
}

// GitCommitWebURLContains applies the Contains predicate on the "git_commit_web_url" field.
func GitCommitWebURLContains(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldContains(FieldGitCommitWebURL, v))
}

// GitCommitWebURLHasPrefix applies the HasPrefix predicate on the "git_commit_web_url" field.
func GitCommitWebURLHasPrefix(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldHasPrefix(FieldGitCommitWebURL, v))
}

// GitCommitWebURLHasSuffix applies the HasSuffix predicate on the "git_commit_web_url" field.
func GitCommitWebURLHasSuffix(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldHasSuffix(FieldGitCommitWebURL, v))
}

// GitCommitWebURLIsNil applies the IsNil predicate on the "git_commit_web_url" field.
func GitCommitWebURLIsNil() predicate.Changelog {
	return predicate.Changelog(sql.FieldIsNull(FieldGitCommitWebURL))
}

// GitCommitWebURLNotNil applies the NotNil predicate on the "git_commit_web_url" field.
func GitCommitWebURLNotNil() predicate.Changelog {
	return predicate.Changelog(sql.FieldNotNull(FieldGitCommitWebURL))
}

// GitCommitWebURLEqualFold applies the EqualFold predicate on the "git_commit_web_url" field.
func GitCommitWebURLEqualFold(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldEqualFold(FieldGitCommitWebURL, v))
}

// GitCommitWebURLContainsFold applies the ContainsFold predicate on the "git_commit_web_url" field.
func GitCommitWebURLContainsFold(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldContainsFold(FieldGitCommitWebURL, v))
}

// GitCommitTitleEQ applies the EQ predicate on the "git_commit_title" field.
func GitCommitTitleEQ(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldGitCommitTitle, v))
}

// GitCommitTitleNEQ applies the NEQ predicate on the "git_commit_title" field.
func GitCommitTitleNEQ(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldNEQ(FieldGitCommitTitle, v))
}

// GitCommitTitleIn applies the In predicate on the "git_commit_title" field.
func GitCommitTitleIn(vs ...string) predicate.Changelog {
	return predicate.Changelog(sql.FieldIn(FieldGitCommitTitle, vs...))
}

// GitCommitTitleNotIn applies the NotIn predicate on the "git_commit_title" field.
func GitCommitTitleNotIn(vs ...string) predicate.Changelog {
	return predicate.Changelog(sql.FieldNotIn(FieldGitCommitTitle, vs...))
}

// GitCommitTitleGT applies the GT predicate on the "git_commit_title" field.
func GitCommitTitleGT(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldGT(FieldGitCommitTitle, v))
}

// GitCommitTitleGTE applies the GTE predicate on the "git_commit_title" field.
func GitCommitTitleGTE(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldGTE(FieldGitCommitTitle, v))
}

// GitCommitTitleLT applies the LT predicate on the "git_commit_title" field.
func GitCommitTitleLT(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldLT(FieldGitCommitTitle, v))
}

// GitCommitTitleLTE applies the LTE predicate on the "git_commit_title" field.
func GitCommitTitleLTE(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldLTE(FieldGitCommitTitle, v))
}

// GitCommitTitleContains applies the Contains predicate on the "git_commit_title" field.
func GitCommitTitleContains(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldContains(FieldGitCommitTitle, v))
}

// GitCommitTitleHasPrefix applies the HasPrefix predicate on the "git_commit_title" field.
func GitCommitTitleHasPrefix(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldHasPrefix(FieldGitCommitTitle, v))
}

// GitCommitTitleHasSuffix applies the HasSuffix predicate on the "git_commit_title" field.
func GitCommitTitleHasSuffix(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldHasSuffix(FieldGitCommitTitle, v))
}

// GitCommitTitleIsNil applies the IsNil predicate on the "git_commit_title" field.
func GitCommitTitleIsNil() predicate.Changelog {
	return predicate.Changelog(sql.FieldIsNull(FieldGitCommitTitle))
}

// GitCommitTitleNotNil applies the NotNil predicate on the "git_commit_title" field.
func GitCommitTitleNotNil() predicate.Changelog {
	return predicate.Changelog(sql.FieldNotNull(FieldGitCommitTitle))
}

// GitCommitTitleEqualFold applies the EqualFold predicate on the "git_commit_title" field.
func GitCommitTitleEqualFold(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldEqualFold(FieldGitCommitTitle, v))
}

// GitCommitTitleContainsFold applies the ContainsFold predicate on the "git_commit_title" field.
func GitCommitTitleContainsFold(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldContainsFold(FieldGitCommitTitle, v))
}

// GitCommitAuthorEQ applies the EQ predicate on the "git_commit_author" field.
func GitCommitAuthorEQ(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldGitCommitAuthor, v))
}

// GitCommitAuthorNEQ applies the NEQ predicate on the "git_commit_author" field.
func GitCommitAuthorNEQ(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldNEQ(FieldGitCommitAuthor, v))
}

// GitCommitAuthorIn applies the In predicate on the "git_commit_author" field.
func GitCommitAuthorIn(vs ...string) predicate.Changelog {
	return predicate.Changelog(sql.FieldIn(FieldGitCommitAuthor, vs...))
}

// GitCommitAuthorNotIn applies the NotIn predicate on the "git_commit_author" field.
func GitCommitAuthorNotIn(vs ...string) predicate.Changelog {
	return predicate.Changelog(sql.FieldNotIn(FieldGitCommitAuthor, vs...))
}

// GitCommitAuthorGT applies the GT predicate on the "git_commit_author" field.
func GitCommitAuthorGT(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldGT(FieldGitCommitAuthor, v))
}

// GitCommitAuthorGTE applies the GTE predicate on the "git_commit_author" field.
func GitCommitAuthorGTE(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldGTE(FieldGitCommitAuthor, v))
}

// GitCommitAuthorLT applies the LT predicate on the "git_commit_author" field.
func GitCommitAuthorLT(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldLT(FieldGitCommitAuthor, v))
}

// GitCommitAuthorLTE applies the LTE predicate on the "git_commit_author" field.
func GitCommitAuthorLTE(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldLTE(FieldGitCommitAuthor, v))
}

// GitCommitAuthorContains applies the Contains predicate on the "git_commit_author" field.
func GitCommitAuthorContains(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldContains(FieldGitCommitAuthor, v))
}

// GitCommitAuthorHasPrefix applies the HasPrefix predicate on the "git_commit_author" field.
func GitCommitAuthorHasPrefix(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldHasPrefix(FieldGitCommitAuthor, v))
}

// GitCommitAuthorHasSuffix applies the HasSuffix predicate on the "git_commit_author" field.
func GitCommitAuthorHasSuffix(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldHasSuffix(FieldGitCommitAuthor, v))
}

// GitCommitAuthorIsNil applies the IsNil predicate on the "git_commit_author" field.
func GitCommitAuthorIsNil() predicate.Changelog {
	return predicate.Changelog(sql.FieldIsNull(FieldGitCommitAuthor))
}

// GitCommitAuthorNotNil applies the NotNil predicate on the "git_commit_author" field.
func GitCommitAuthorNotNil() predicate.Changelog {
	return predicate.Changelog(sql.FieldNotNull(FieldGitCommitAuthor))
}

// GitCommitAuthorEqualFold applies the EqualFold predicate on the "git_commit_author" field.
func GitCommitAuthorEqualFold(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldEqualFold(FieldGitCommitAuthor, v))
}

// GitCommitAuthorContainsFold applies the ContainsFold predicate on the "git_commit_author" field.
func GitCommitAuthorContainsFold(v string) predicate.Changelog {
	return predicate.Changelog(sql.FieldContainsFold(FieldGitCommitAuthor, v))
}

// GitCommitDateEQ applies the EQ predicate on the "git_commit_date" field.
func GitCommitDateEQ(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldGitCommitDate, v))
}

// GitCommitDateNEQ applies the NEQ predicate on the "git_commit_date" field.
func GitCommitDateNEQ(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldNEQ(FieldGitCommitDate, v))
}

// GitCommitDateIn applies the In predicate on the "git_commit_date" field.
func GitCommitDateIn(vs ...time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldIn(FieldGitCommitDate, vs...))
}

// GitCommitDateNotIn applies the NotIn predicate on the "git_commit_date" field.
func GitCommitDateNotIn(vs ...time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldNotIn(FieldGitCommitDate, vs...))
}

// GitCommitDateGT applies the GT predicate on the "git_commit_date" field.
func GitCommitDateGT(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldGT(FieldGitCommitDate, v))
}

// GitCommitDateGTE applies the GTE predicate on the "git_commit_date" field.
func GitCommitDateGTE(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldGTE(FieldGitCommitDate, v))
}

// GitCommitDateLT applies the LT predicate on the "git_commit_date" field.
func GitCommitDateLT(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldLT(FieldGitCommitDate, v))
}

// GitCommitDateLTE applies the LTE predicate on the "git_commit_date" field.
func GitCommitDateLTE(v time.Time) predicate.Changelog {
	return predicate.Changelog(sql.FieldLTE(FieldGitCommitDate, v))
}

// GitCommitDateIsNil applies the IsNil predicate on the "git_commit_date" field.
func GitCommitDateIsNil() predicate.Changelog {
	return predicate.Changelog(sql.FieldIsNull(FieldGitCommitDate))
}

// GitCommitDateNotNil applies the NotNil predicate on the "git_commit_date" field.
func GitCommitDateNotNil() predicate.Changelog {
	return predicate.Changelog(sql.FieldNotNull(FieldGitCommitDate))
}

// ConfigChangedEQ applies the EQ predicate on the "config_changed" field.
func ConfigChangedEQ(v bool) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldConfigChanged, v))
}

// ConfigChangedNEQ applies the NEQ predicate on the "config_changed" field.
func ConfigChangedNEQ(v bool) predicate.Changelog {
	return predicate.Changelog(sql.FieldNEQ(FieldConfigChanged, v))
}

// ProjectIDEQ applies the EQ predicate on the "project_id" field.
func ProjectIDEQ(v int) predicate.Changelog {
	return predicate.Changelog(sql.FieldEQ(FieldProjectID, v))
}

// ProjectIDNEQ applies the NEQ predicate on the "project_id" field.
func ProjectIDNEQ(v int) predicate.Changelog {
	return predicate.Changelog(sql.FieldNEQ(FieldProjectID, v))
}

// ProjectIDIn applies the In predicate on the "project_id" field.
func ProjectIDIn(vs ...int) predicate.Changelog {
	return predicate.Changelog(sql.FieldIn(FieldProjectID, vs...))
}

// ProjectIDNotIn applies the NotIn predicate on the "project_id" field.
func ProjectIDNotIn(vs ...int) predicate.Changelog {
	return predicate.Changelog(sql.FieldNotIn(FieldProjectID, vs...))
}

// ProjectIDIsNil applies the IsNil predicate on the "project_id" field.
func ProjectIDIsNil() predicate.Changelog {
	return predicate.Changelog(sql.FieldIsNull(FieldProjectID))
}

// ProjectIDNotNil applies the NotNil predicate on the "project_id" field.
func ProjectIDNotNil() predicate.Changelog {
	return predicate.Changelog(sql.FieldNotNull(FieldProjectID))
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.Changelog {
	return predicate.Changelog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectWith applies the HasEdge predicate on the "project" edge with a given conditions (other predicates).
func HasProjectWith(preds ...predicate.Project) predicate.Changelog {
	return predicate.Changelog(func(s *sql.Selector) {
		step := newProjectStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Changelog) predicate.Changelog {
	return predicate.Changelog(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Changelog) predicate.Changelog {
	return predicate.Changelog(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Changelog) predicate.Changelog {
	return predicate.Changelog(sql.NotPredicates(p))
}
