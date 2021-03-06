// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/kcarretto/paragon/ent/job"
	"github.com/kcarretto/paragon/ent/predicate"
	"github.com/kcarretto/paragon/ent/tag"
	"github.com/kcarretto/paragon/ent/target"
	"github.com/kcarretto/paragon/ent/task"
)

// TagUpdate is the builder for updating Tag entities.
type TagUpdate struct {
	config
	Name           *string
	targets        map[int]struct{}
	tasks          map[int]struct{}
	jobs           map[int]struct{}
	removedTargets map[int]struct{}
	removedTasks   map[int]struct{}
	removedJobs    map[int]struct{}
	predicates     []predicate.Tag
}

// Where adds a new predicate for the builder.
func (tu *TagUpdate) Where(ps ...predicate.Tag) *TagUpdate {
	tu.predicates = append(tu.predicates, ps...)
	return tu
}

// SetName sets the Name field.
func (tu *TagUpdate) SetName(s string) *TagUpdate {
	tu.Name = &s
	return tu
}

// AddTargetIDs adds the targets edge to Target by ids.
func (tu *TagUpdate) AddTargetIDs(ids ...int) *TagUpdate {
	if tu.targets == nil {
		tu.targets = make(map[int]struct{})
	}
	for i := range ids {
		tu.targets[ids[i]] = struct{}{}
	}
	return tu
}

// AddTargets adds the targets edges to Target.
func (tu *TagUpdate) AddTargets(t ...*Target) *TagUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddTargetIDs(ids...)
}

// AddTaskIDs adds the tasks edge to Task by ids.
func (tu *TagUpdate) AddTaskIDs(ids ...int) *TagUpdate {
	if tu.tasks == nil {
		tu.tasks = make(map[int]struct{})
	}
	for i := range ids {
		tu.tasks[ids[i]] = struct{}{}
	}
	return tu
}

// AddTasks adds the tasks edges to Task.
func (tu *TagUpdate) AddTasks(t ...*Task) *TagUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddTaskIDs(ids...)
}

// AddJobIDs adds the jobs edge to Job by ids.
func (tu *TagUpdate) AddJobIDs(ids ...int) *TagUpdate {
	if tu.jobs == nil {
		tu.jobs = make(map[int]struct{})
	}
	for i := range ids {
		tu.jobs[ids[i]] = struct{}{}
	}
	return tu
}

// AddJobs adds the jobs edges to Job.
func (tu *TagUpdate) AddJobs(j ...*Job) *TagUpdate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return tu.AddJobIDs(ids...)
}

// RemoveTargetIDs removes the targets edge to Target by ids.
func (tu *TagUpdate) RemoveTargetIDs(ids ...int) *TagUpdate {
	if tu.removedTargets == nil {
		tu.removedTargets = make(map[int]struct{})
	}
	for i := range ids {
		tu.removedTargets[ids[i]] = struct{}{}
	}
	return tu
}

// RemoveTargets removes targets edges to Target.
func (tu *TagUpdate) RemoveTargets(t ...*Target) *TagUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveTargetIDs(ids...)
}

// RemoveTaskIDs removes the tasks edge to Task by ids.
func (tu *TagUpdate) RemoveTaskIDs(ids ...int) *TagUpdate {
	if tu.removedTasks == nil {
		tu.removedTasks = make(map[int]struct{})
	}
	for i := range ids {
		tu.removedTasks[ids[i]] = struct{}{}
	}
	return tu
}

// RemoveTasks removes tasks edges to Task.
func (tu *TagUpdate) RemoveTasks(t ...*Task) *TagUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveTaskIDs(ids...)
}

// RemoveJobIDs removes the jobs edge to Job by ids.
func (tu *TagUpdate) RemoveJobIDs(ids ...int) *TagUpdate {
	if tu.removedJobs == nil {
		tu.removedJobs = make(map[int]struct{})
	}
	for i := range ids {
		tu.removedJobs[ids[i]] = struct{}{}
	}
	return tu
}

// RemoveJobs removes jobs edges to Job.
func (tu *TagUpdate) RemoveJobs(j ...*Job) *TagUpdate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return tu.RemoveJobIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (tu *TagUpdate) Save(ctx context.Context) (int, error) {
	if tu.Name != nil {
		if err := tag.NameValidator(*tu.Name); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"Name\": %v", err)
		}
	}
	return tu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TagUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TagUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TagUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tag.Table,
			Columns: tag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tag.FieldID,
			},
		},
	}
	if ps := tu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := tu.Name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: tag.FieldName,
		})
	}
	if nodes := tu.removedTargets; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.TargetsTable,
			Columns: tag.TargetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: target.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.targets; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.TargetsTable,
			Columns: tag.TargetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: target.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := tu.removedTasks; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.TasksTable,
			Columns: tag.TasksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.tasks; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.TasksTable,
			Columns: tag.TasksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := tu.removedJobs; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.JobsTable,
			Columns: tag.JobsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: job.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.jobs; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.JobsTable,
			Columns: tag.JobsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: job.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// TagUpdateOne is the builder for updating a single Tag entity.
type TagUpdateOne struct {
	config
	id             int
	Name           *string
	targets        map[int]struct{}
	tasks          map[int]struct{}
	jobs           map[int]struct{}
	removedTargets map[int]struct{}
	removedTasks   map[int]struct{}
	removedJobs    map[int]struct{}
}

// SetName sets the Name field.
func (tuo *TagUpdateOne) SetName(s string) *TagUpdateOne {
	tuo.Name = &s
	return tuo
}

// AddTargetIDs adds the targets edge to Target by ids.
func (tuo *TagUpdateOne) AddTargetIDs(ids ...int) *TagUpdateOne {
	if tuo.targets == nil {
		tuo.targets = make(map[int]struct{})
	}
	for i := range ids {
		tuo.targets[ids[i]] = struct{}{}
	}
	return tuo
}

// AddTargets adds the targets edges to Target.
func (tuo *TagUpdateOne) AddTargets(t ...*Target) *TagUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddTargetIDs(ids...)
}

// AddTaskIDs adds the tasks edge to Task by ids.
func (tuo *TagUpdateOne) AddTaskIDs(ids ...int) *TagUpdateOne {
	if tuo.tasks == nil {
		tuo.tasks = make(map[int]struct{})
	}
	for i := range ids {
		tuo.tasks[ids[i]] = struct{}{}
	}
	return tuo
}

// AddTasks adds the tasks edges to Task.
func (tuo *TagUpdateOne) AddTasks(t ...*Task) *TagUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddTaskIDs(ids...)
}

// AddJobIDs adds the jobs edge to Job by ids.
func (tuo *TagUpdateOne) AddJobIDs(ids ...int) *TagUpdateOne {
	if tuo.jobs == nil {
		tuo.jobs = make(map[int]struct{})
	}
	for i := range ids {
		tuo.jobs[ids[i]] = struct{}{}
	}
	return tuo
}

// AddJobs adds the jobs edges to Job.
func (tuo *TagUpdateOne) AddJobs(j ...*Job) *TagUpdateOne {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return tuo.AddJobIDs(ids...)
}

// RemoveTargetIDs removes the targets edge to Target by ids.
func (tuo *TagUpdateOne) RemoveTargetIDs(ids ...int) *TagUpdateOne {
	if tuo.removedTargets == nil {
		tuo.removedTargets = make(map[int]struct{})
	}
	for i := range ids {
		tuo.removedTargets[ids[i]] = struct{}{}
	}
	return tuo
}

// RemoveTargets removes targets edges to Target.
func (tuo *TagUpdateOne) RemoveTargets(t ...*Target) *TagUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveTargetIDs(ids...)
}

// RemoveTaskIDs removes the tasks edge to Task by ids.
func (tuo *TagUpdateOne) RemoveTaskIDs(ids ...int) *TagUpdateOne {
	if tuo.removedTasks == nil {
		tuo.removedTasks = make(map[int]struct{})
	}
	for i := range ids {
		tuo.removedTasks[ids[i]] = struct{}{}
	}
	return tuo
}

// RemoveTasks removes tasks edges to Task.
func (tuo *TagUpdateOne) RemoveTasks(t ...*Task) *TagUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveTaskIDs(ids...)
}

// RemoveJobIDs removes the jobs edge to Job by ids.
func (tuo *TagUpdateOne) RemoveJobIDs(ids ...int) *TagUpdateOne {
	if tuo.removedJobs == nil {
		tuo.removedJobs = make(map[int]struct{})
	}
	for i := range ids {
		tuo.removedJobs[ids[i]] = struct{}{}
	}
	return tuo
}

// RemoveJobs removes jobs edges to Job.
func (tuo *TagUpdateOne) RemoveJobs(j ...*Job) *TagUpdateOne {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return tuo.RemoveJobIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (tuo *TagUpdateOne) Save(ctx context.Context) (*Tag, error) {
	if tuo.Name != nil {
		if err := tag.NameValidator(*tuo.Name); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"Name\": %v", err)
		}
	}
	return tuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TagUpdateOne) SaveX(ctx context.Context) *Tag {
	t, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return t
}

// Exec executes the query on the entity.
func (tuo *TagUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TagUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TagUpdateOne) sqlSave(ctx context.Context) (t *Tag, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tag.Table,
			Columns: tag.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  tuo.id,
				Type:   field.TypeInt,
				Column: tag.FieldID,
			},
		},
	}
	if value := tuo.Name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: tag.FieldName,
		})
	}
	if nodes := tuo.removedTargets; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.TargetsTable,
			Columns: tag.TargetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: target.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.targets; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.TargetsTable,
			Columns: tag.TargetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: target.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := tuo.removedTasks; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.TasksTable,
			Columns: tag.TasksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.tasks; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.TasksTable,
			Columns: tag.TasksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := tuo.removedJobs; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.JobsTable,
			Columns: tag.JobsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: job.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.jobs; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.JobsTable,
			Columns: tag.JobsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: job.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	t = &Tag{config: tuo.config}
	_spec.Assign = t.assignValues
	_spec.ScanValues = t.scanValues()
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return t, nil
}
