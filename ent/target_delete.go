// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/kcarretto/paragon/ent/predicate"
	"github.com/kcarretto/paragon/ent/target"
)

// TargetDelete is the builder for deleting a Target entity.
type TargetDelete struct {
	config
	predicates []predicate.Target
}

// Where adds a new predicate to the delete builder.
func (td *TargetDelete) Where(ps ...predicate.Target) *TargetDelete {
	td.predicates = append(td.predicates, ps...)
	return td
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (td *TargetDelete) Exec(ctx context.Context) (int, error) {
	return td.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (td *TargetDelete) ExecX(ctx context.Context) int {
	n, err := td.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (td *TargetDelete) sqlExec(ctx context.Context) (int, error) {
	var res sql.Result
	selector := sql.Select().From(sql.Table(target.Table))
	for _, p := range td.predicates {
		p(selector)
	}
	query, args := sql.Delete(target.Table).FromSelect(selector).Query()
	if err := td.driver.Exec(ctx, query, args, &res); err != nil {
		return 0, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(affected), nil
}

// TargetDeleteOne is the builder for deleting a single Target entity.
type TargetDeleteOne struct {
	td *TargetDelete
}

// Exec executes the deletion query.
func (tdo *TargetDeleteOne) Exec(ctx context.Context) error {
	n, err := tdo.td.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{target.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tdo *TargetDeleteOne) ExecX(ctx context.Context) {
	tdo.td.ExecX(ctx)
}