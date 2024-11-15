// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/duc-cnzj/mars/v5/internal/ent/cachelock"
	"github.com/duc-cnzj/mars/v5/internal/ent/predicate"
)

// CacheLockDelete is the builder for deleting a CacheLock entity.
type CacheLockDelete struct {
	config
	hooks    []Hook
	mutation *CacheLockMutation
}

// Where appends a list predicates to the CacheLockDelete builder.
func (cld *CacheLockDelete) Where(ps ...predicate.CacheLock) *CacheLockDelete {
	cld.mutation.Where(ps...)
	return cld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cld *CacheLockDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cld.sqlExec, cld.mutation, cld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cld *CacheLockDelete) ExecX(ctx context.Context) int {
	n, err := cld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cld *CacheLockDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(cachelock.Table, sqlgraph.NewFieldSpec(cachelock.FieldID, field.TypeInt))
	if ps := cld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cld.mutation.done = true
	return affected, err
}

// CacheLockDeleteOne is the builder for deleting a single CacheLock entity.
type CacheLockDeleteOne struct {
	cld *CacheLockDelete
}

// Where appends a list predicates to the CacheLockDelete builder.
func (cldo *CacheLockDeleteOne) Where(ps ...predicate.CacheLock) *CacheLockDeleteOne {
	cldo.cld.mutation.Where(ps...)
	return cldo
}

// Exec executes the deletion query.
func (cldo *CacheLockDeleteOne) Exec(ctx context.Context) error {
	n, err := cldo.cld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{cachelock.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cldo *CacheLockDeleteOne) ExecX(ctx context.Context) {
	if err := cldo.Exec(ctx); err != nil {
		panic(err)
	}
}
