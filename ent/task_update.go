// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/kcarretto/paragon/ent/job"
	"github.com/kcarretto/paragon/ent/predicate"
	"github.com/kcarretto/paragon/ent/target"
	"github.com/kcarretto/paragon/ent/task"
)

// TaskUpdate is the builder for updating Task entities.
type TaskUpdate struct {
	config
	QueueTime          *time.Time
	LastChangedTime    *time.Time
	ClaimTime          *time.Time
	clearClaimTime     bool
	ExecStartTime      *time.Time
	clearExecStartTime bool
	ExecStopTime       *time.Time
	clearExecStopTime  bool
	Content            *string
	Output             *string
	clearOutput        bool
	Error              *string
	clearError         bool
	SessionID          *string
	clearSessionID     bool
	tags               map[int]struct{}
	job                map[int]struct{}
	target             map[int]struct{}
	removedTags        map[int]struct{}
	clearedJob         bool
	clearedTarget      bool
	predicates         []predicate.Task
}

// Where adds a new predicate for the builder.
func (tu *TaskUpdate) Where(ps ...predicate.Task) *TaskUpdate {
	tu.predicates = append(tu.predicates, ps...)
	return tu
}

// SetQueueTime sets the QueueTime field.
func (tu *TaskUpdate) SetQueueTime(t time.Time) *TaskUpdate {
	tu.QueueTime = &t
	return tu
}

// SetNillableQueueTime sets the QueueTime field if the given value is not nil.
func (tu *TaskUpdate) SetNillableQueueTime(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetQueueTime(*t)
	}
	return tu
}

// SetLastChangedTime sets the LastChangedTime field.
func (tu *TaskUpdate) SetLastChangedTime(t time.Time) *TaskUpdate {
	tu.LastChangedTime = &t
	return tu
}

// SetClaimTime sets the ClaimTime field.
func (tu *TaskUpdate) SetClaimTime(t time.Time) *TaskUpdate {
	tu.ClaimTime = &t
	return tu
}

// SetNillableClaimTime sets the ClaimTime field if the given value is not nil.
func (tu *TaskUpdate) SetNillableClaimTime(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetClaimTime(*t)
	}
	return tu
}

// ClearClaimTime clears the value of ClaimTime.
func (tu *TaskUpdate) ClearClaimTime() *TaskUpdate {
	tu.ClaimTime = nil
	tu.clearClaimTime = true
	return tu
}

// SetExecStartTime sets the ExecStartTime field.
func (tu *TaskUpdate) SetExecStartTime(t time.Time) *TaskUpdate {
	tu.ExecStartTime = &t
	return tu
}

// SetNillableExecStartTime sets the ExecStartTime field if the given value is not nil.
func (tu *TaskUpdate) SetNillableExecStartTime(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetExecStartTime(*t)
	}
	return tu
}

// ClearExecStartTime clears the value of ExecStartTime.
func (tu *TaskUpdate) ClearExecStartTime() *TaskUpdate {
	tu.ExecStartTime = nil
	tu.clearExecStartTime = true
	return tu
}

// SetExecStopTime sets the ExecStopTime field.
func (tu *TaskUpdate) SetExecStopTime(t time.Time) *TaskUpdate {
	tu.ExecStopTime = &t
	return tu
}

// SetNillableExecStopTime sets the ExecStopTime field if the given value is not nil.
func (tu *TaskUpdate) SetNillableExecStopTime(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetExecStopTime(*t)
	}
	return tu
}

// ClearExecStopTime clears the value of ExecStopTime.
func (tu *TaskUpdate) ClearExecStopTime() *TaskUpdate {
	tu.ExecStopTime = nil
	tu.clearExecStopTime = true
	return tu
}

// SetContent sets the Content field.
func (tu *TaskUpdate) SetContent(s string) *TaskUpdate {
	tu.Content = &s
	return tu
}

// SetOutput sets the Output field.
func (tu *TaskUpdate) SetOutput(s string) *TaskUpdate {
	tu.Output = &s
	return tu
}

// SetNillableOutput sets the Output field if the given value is not nil.
func (tu *TaskUpdate) SetNillableOutput(s *string) *TaskUpdate {
	if s != nil {
		tu.SetOutput(*s)
	}
	return tu
}

// ClearOutput clears the value of Output.
func (tu *TaskUpdate) ClearOutput() *TaskUpdate {
	tu.Output = nil
	tu.clearOutput = true
	return tu
}

// SetError sets the Error field.
func (tu *TaskUpdate) SetError(s string) *TaskUpdate {
	tu.Error = &s
	return tu
}

// SetNillableError sets the Error field if the given value is not nil.
func (tu *TaskUpdate) SetNillableError(s *string) *TaskUpdate {
	if s != nil {
		tu.SetError(*s)
	}
	return tu
}

// ClearError clears the value of Error.
func (tu *TaskUpdate) ClearError() *TaskUpdate {
	tu.Error = nil
	tu.clearError = true
	return tu
}

// SetSessionID sets the SessionID field.
func (tu *TaskUpdate) SetSessionID(s string) *TaskUpdate {
	tu.SessionID = &s
	return tu
}

// SetNillableSessionID sets the SessionID field if the given value is not nil.
func (tu *TaskUpdate) SetNillableSessionID(s *string) *TaskUpdate {
	if s != nil {
		tu.SetSessionID(*s)
	}
	return tu
}

// ClearSessionID clears the value of SessionID.
func (tu *TaskUpdate) ClearSessionID() *TaskUpdate {
	tu.SessionID = nil
	tu.clearSessionID = true
	return tu
}

// AddTagIDs adds the tags edge to Tag by ids.
func (tu *TaskUpdate) AddTagIDs(ids ...int) *TaskUpdate {
	if tu.tags == nil {
		tu.tags = make(map[int]struct{})
	}
	for i := range ids {
		tu.tags[ids[i]] = struct{}{}
	}
	return tu
}

// AddTags adds the tags edges to Tag.
func (tu *TaskUpdate) AddTags(t ...*Tag) *TaskUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddTagIDs(ids...)
}

// SetJobID sets the job edge to Job by id.
func (tu *TaskUpdate) SetJobID(id int) *TaskUpdate {
	if tu.job == nil {
		tu.job = make(map[int]struct{})
	}
	tu.job[id] = struct{}{}
	return tu
}

// SetJob sets the job edge to Job.
func (tu *TaskUpdate) SetJob(j *Job) *TaskUpdate {
	return tu.SetJobID(j.ID)
}

// SetTargetID sets the target edge to Target by id.
func (tu *TaskUpdate) SetTargetID(id int) *TaskUpdate {
	if tu.target == nil {
		tu.target = make(map[int]struct{})
	}
	tu.target[id] = struct{}{}
	return tu
}

// SetNillableTargetID sets the target edge to Target by id if the given value is not nil.
func (tu *TaskUpdate) SetNillableTargetID(id *int) *TaskUpdate {
	if id != nil {
		tu = tu.SetTargetID(*id)
	}
	return tu
}

// SetTarget sets the target edge to Target.
func (tu *TaskUpdate) SetTarget(t *Target) *TaskUpdate {
	return tu.SetTargetID(t.ID)
}

// RemoveTagIDs removes the tags edge to Tag by ids.
func (tu *TaskUpdate) RemoveTagIDs(ids ...int) *TaskUpdate {
	if tu.removedTags == nil {
		tu.removedTags = make(map[int]struct{})
	}
	for i := range ids {
		tu.removedTags[ids[i]] = struct{}{}
	}
	return tu
}

// RemoveTags removes tags edges to Tag.
func (tu *TaskUpdate) RemoveTags(t ...*Tag) *TaskUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveTagIDs(ids...)
}

// ClearJob clears the job edge to Job.
func (tu *TaskUpdate) ClearJob() *TaskUpdate {
	tu.clearedJob = true
	return tu
}

// ClearTarget clears the target edge to Target.
func (tu *TaskUpdate) ClearTarget() *TaskUpdate {
	tu.clearedTarget = true
	return tu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (tu *TaskUpdate) Save(ctx context.Context) (int, error) {
	if tu.Content != nil {
		if err := task.ContentValidator(*tu.Content); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"Content\": %v", err)
		}
	}
	if len(tu.job) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"job\"")
	}
	if tu.clearedJob && tu.job == nil {
		return 0, errors.New("ent: clearing a unique edge \"job\"")
	}
	if len(tu.target) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"target\"")
	}
	return tu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TaskUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TaskUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TaskUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	selector := sql.Select(task.FieldID).From(sql.Table(task.Table))
	for _, p := range tu.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = tu.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return 0, fmt.Errorf("ent: failed reading id: %v", err)
		}
		ids = append(ids, id)
	}
	if len(ids) == 0 {
		return 0, nil
	}

	tx, err := tu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		res     sql.Result
		builder = sql.Update(task.Table).Where(sql.InInts(task.FieldID, ids...))
	)
	if value := tu.QueueTime; value != nil {
		builder.Set(task.FieldQueueTime, *value)
	}
	if value := tu.LastChangedTime; value != nil {
		builder.Set(task.FieldLastChangedTime, *value)
	}
	if value := tu.ClaimTime; value != nil {
		builder.Set(task.FieldClaimTime, *value)
	}
	if tu.clearClaimTime {
		builder.SetNull(task.FieldClaimTime)
	}
	if value := tu.ExecStartTime; value != nil {
		builder.Set(task.FieldExecStartTime, *value)
	}
	if tu.clearExecStartTime {
		builder.SetNull(task.FieldExecStartTime)
	}
	if value := tu.ExecStopTime; value != nil {
		builder.Set(task.FieldExecStopTime, *value)
	}
	if tu.clearExecStopTime {
		builder.SetNull(task.FieldExecStopTime)
	}
	if value := tu.Content; value != nil {
		builder.Set(task.FieldContent, *value)
	}
	if value := tu.Output; value != nil {
		builder.Set(task.FieldOutput, *value)
	}
	if tu.clearOutput {
		builder.SetNull(task.FieldOutput)
	}
	if value := tu.Error; value != nil {
		builder.Set(task.FieldError, *value)
	}
	if tu.clearError {
		builder.SetNull(task.FieldError)
	}
	if value := tu.SessionID; value != nil {
		builder.Set(task.FieldSessionID, *value)
	}
	if tu.clearSessionID {
		builder.SetNull(task.FieldSessionID)
	}
	if !builder.Empty() {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(tu.removedTags) > 0 {
		eids := make([]int, len(tu.removedTags))
		for eid := range tu.removedTags {
			eids = append(eids, eid)
		}
		query, args := sql.Delete(task.TagsTable).
			Where(sql.InInts(task.TagsPrimaryKey[0], ids...)).
			Where(sql.InInts(task.TagsPrimaryKey[1], eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(tu.tags) > 0 {
		values := make([][]int, 0, len(ids))
		for _, id := range ids {
			for eid := range tu.tags {
				values = append(values, []int{id, eid})
			}
		}
		builder := sql.Insert(task.TagsTable).
			Columns(task.TagsPrimaryKey[0], task.TagsPrimaryKey[1])
		for _, v := range values {
			builder.Values(v[0], v[1])
		}
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if tu.clearedJob {
		query, args := sql.Update(task.JobTable).
			SetNull(task.JobColumn).
			Where(sql.InInts(job.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(tu.job) > 0 {
		for eid := range tu.job {
			query, args := sql.Update(task.JobTable).
				Set(task.JobColumn, eid).
				Where(sql.InInts(task.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
		}
	}
	if tu.clearedTarget {
		query, args := sql.Update(task.TargetTable).
			SetNull(task.TargetColumn).
			Where(sql.InInts(target.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(tu.target) > 0 {
		for eid := range tu.target {
			query, args := sql.Update(task.TargetTable).
				Set(task.TargetColumn, eid).
				Where(sql.InInts(task.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

// TaskUpdateOne is the builder for updating a single Task entity.
type TaskUpdateOne struct {
	config
	id                 int
	QueueTime          *time.Time
	LastChangedTime    *time.Time
	ClaimTime          *time.Time
	clearClaimTime     bool
	ExecStartTime      *time.Time
	clearExecStartTime bool
	ExecStopTime       *time.Time
	clearExecStopTime  bool
	Content            *string
	Output             *string
	clearOutput        bool
	Error              *string
	clearError         bool
	SessionID          *string
	clearSessionID     bool
	tags               map[int]struct{}
	job                map[int]struct{}
	target             map[int]struct{}
	removedTags        map[int]struct{}
	clearedJob         bool
	clearedTarget      bool
}

// SetQueueTime sets the QueueTime field.
func (tuo *TaskUpdateOne) SetQueueTime(t time.Time) *TaskUpdateOne {
	tuo.QueueTime = &t
	return tuo
}

// SetNillableQueueTime sets the QueueTime field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableQueueTime(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetQueueTime(*t)
	}
	return tuo
}

// SetLastChangedTime sets the LastChangedTime field.
func (tuo *TaskUpdateOne) SetLastChangedTime(t time.Time) *TaskUpdateOne {
	tuo.LastChangedTime = &t
	return tuo
}

// SetClaimTime sets the ClaimTime field.
func (tuo *TaskUpdateOne) SetClaimTime(t time.Time) *TaskUpdateOne {
	tuo.ClaimTime = &t
	return tuo
}

// SetNillableClaimTime sets the ClaimTime field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableClaimTime(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetClaimTime(*t)
	}
	return tuo
}

// ClearClaimTime clears the value of ClaimTime.
func (tuo *TaskUpdateOne) ClearClaimTime() *TaskUpdateOne {
	tuo.ClaimTime = nil
	tuo.clearClaimTime = true
	return tuo
}

// SetExecStartTime sets the ExecStartTime field.
func (tuo *TaskUpdateOne) SetExecStartTime(t time.Time) *TaskUpdateOne {
	tuo.ExecStartTime = &t
	return tuo
}

// SetNillableExecStartTime sets the ExecStartTime field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableExecStartTime(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetExecStartTime(*t)
	}
	return tuo
}

// ClearExecStartTime clears the value of ExecStartTime.
func (tuo *TaskUpdateOne) ClearExecStartTime() *TaskUpdateOne {
	tuo.ExecStartTime = nil
	tuo.clearExecStartTime = true
	return tuo
}

// SetExecStopTime sets the ExecStopTime field.
func (tuo *TaskUpdateOne) SetExecStopTime(t time.Time) *TaskUpdateOne {
	tuo.ExecStopTime = &t
	return tuo
}

// SetNillableExecStopTime sets the ExecStopTime field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableExecStopTime(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetExecStopTime(*t)
	}
	return tuo
}

// ClearExecStopTime clears the value of ExecStopTime.
func (tuo *TaskUpdateOne) ClearExecStopTime() *TaskUpdateOne {
	tuo.ExecStopTime = nil
	tuo.clearExecStopTime = true
	return tuo
}

// SetContent sets the Content field.
func (tuo *TaskUpdateOne) SetContent(s string) *TaskUpdateOne {
	tuo.Content = &s
	return tuo
}

// SetOutput sets the Output field.
func (tuo *TaskUpdateOne) SetOutput(s string) *TaskUpdateOne {
	tuo.Output = &s
	return tuo
}

// SetNillableOutput sets the Output field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableOutput(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetOutput(*s)
	}
	return tuo
}

// ClearOutput clears the value of Output.
func (tuo *TaskUpdateOne) ClearOutput() *TaskUpdateOne {
	tuo.Output = nil
	tuo.clearOutput = true
	return tuo
}

// SetError sets the Error field.
func (tuo *TaskUpdateOne) SetError(s string) *TaskUpdateOne {
	tuo.Error = &s
	return tuo
}

// SetNillableError sets the Error field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableError(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetError(*s)
	}
	return tuo
}

// ClearError clears the value of Error.
func (tuo *TaskUpdateOne) ClearError() *TaskUpdateOne {
	tuo.Error = nil
	tuo.clearError = true
	return tuo
}

// SetSessionID sets the SessionID field.
func (tuo *TaskUpdateOne) SetSessionID(s string) *TaskUpdateOne {
	tuo.SessionID = &s
	return tuo
}

// SetNillableSessionID sets the SessionID field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableSessionID(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetSessionID(*s)
	}
	return tuo
}

// ClearSessionID clears the value of SessionID.
func (tuo *TaskUpdateOne) ClearSessionID() *TaskUpdateOne {
	tuo.SessionID = nil
	tuo.clearSessionID = true
	return tuo
}

// AddTagIDs adds the tags edge to Tag by ids.
func (tuo *TaskUpdateOne) AddTagIDs(ids ...int) *TaskUpdateOne {
	if tuo.tags == nil {
		tuo.tags = make(map[int]struct{})
	}
	for i := range ids {
		tuo.tags[ids[i]] = struct{}{}
	}
	return tuo
}

// AddTags adds the tags edges to Tag.
func (tuo *TaskUpdateOne) AddTags(t ...*Tag) *TaskUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddTagIDs(ids...)
}

// SetJobID sets the job edge to Job by id.
func (tuo *TaskUpdateOne) SetJobID(id int) *TaskUpdateOne {
	if tuo.job == nil {
		tuo.job = make(map[int]struct{})
	}
	tuo.job[id] = struct{}{}
	return tuo
}

// SetJob sets the job edge to Job.
func (tuo *TaskUpdateOne) SetJob(j *Job) *TaskUpdateOne {
	return tuo.SetJobID(j.ID)
}

// SetTargetID sets the target edge to Target by id.
func (tuo *TaskUpdateOne) SetTargetID(id int) *TaskUpdateOne {
	if tuo.target == nil {
		tuo.target = make(map[int]struct{})
	}
	tuo.target[id] = struct{}{}
	return tuo
}

// SetNillableTargetID sets the target edge to Target by id if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableTargetID(id *int) *TaskUpdateOne {
	if id != nil {
		tuo = tuo.SetTargetID(*id)
	}
	return tuo
}

// SetTarget sets the target edge to Target.
func (tuo *TaskUpdateOne) SetTarget(t *Target) *TaskUpdateOne {
	return tuo.SetTargetID(t.ID)
}

// RemoveTagIDs removes the tags edge to Tag by ids.
func (tuo *TaskUpdateOne) RemoveTagIDs(ids ...int) *TaskUpdateOne {
	if tuo.removedTags == nil {
		tuo.removedTags = make(map[int]struct{})
	}
	for i := range ids {
		tuo.removedTags[ids[i]] = struct{}{}
	}
	return tuo
}

// RemoveTags removes tags edges to Tag.
func (tuo *TaskUpdateOne) RemoveTags(t ...*Tag) *TaskUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveTagIDs(ids...)
}

// ClearJob clears the job edge to Job.
func (tuo *TaskUpdateOne) ClearJob() *TaskUpdateOne {
	tuo.clearedJob = true
	return tuo
}

// ClearTarget clears the target edge to Target.
func (tuo *TaskUpdateOne) ClearTarget() *TaskUpdateOne {
	tuo.clearedTarget = true
	return tuo
}

// Save executes the query and returns the updated entity.
func (tuo *TaskUpdateOne) Save(ctx context.Context) (*Task, error) {
	if tuo.Content != nil {
		if err := task.ContentValidator(*tuo.Content); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"Content\": %v", err)
		}
	}
	if len(tuo.job) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"job\"")
	}
	if tuo.clearedJob && tuo.job == nil {
		return nil, errors.New("ent: clearing a unique edge \"job\"")
	}
	if len(tuo.target) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"target\"")
	}
	return tuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TaskUpdateOne) SaveX(ctx context.Context) *Task {
	t, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return t
}

// Exec executes the query on the entity.
func (tuo *TaskUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TaskUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TaskUpdateOne) sqlSave(ctx context.Context) (t *Task, err error) {
	selector := sql.Select(task.Columns...).From(sql.Table(task.Table))
	task.ID(tuo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = tuo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		t = &Task{config: tuo.config}
		if err := t.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into Task: %v", err)
		}
		id = t.ID
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, &ErrNotFound{fmt.Sprintf("Task with id: %v", tuo.id)}
	case n > 1:
		return nil, fmt.Errorf("ent: more than one Task with the same id: %v", tuo.id)
	}

	tx, err := tuo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		res     sql.Result
		builder = sql.Update(task.Table).Where(sql.InInts(task.FieldID, ids...))
	)
	if value := tuo.QueueTime; value != nil {
		builder.Set(task.FieldQueueTime, *value)
		t.QueueTime = *value
	}
	if value := tuo.LastChangedTime; value != nil {
		builder.Set(task.FieldLastChangedTime, *value)
		t.LastChangedTime = *value
	}
	if value := tuo.ClaimTime; value != nil {
		builder.Set(task.FieldClaimTime, *value)
		t.ClaimTime = *value
	}
	if tuo.clearClaimTime {
		var value time.Time
		t.ClaimTime = value
		builder.SetNull(task.FieldClaimTime)
	}
	if value := tuo.ExecStartTime; value != nil {
		builder.Set(task.FieldExecStartTime, *value)
		t.ExecStartTime = *value
	}
	if tuo.clearExecStartTime {
		var value time.Time
		t.ExecStartTime = value
		builder.SetNull(task.FieldExecStartTime)
	}
	if value := tuo.ExecStopTime; value != nil {
		builder.Set(task.FieldExecStopTime, *value)
		t.ExecStopTime = *value
	}
	if tuo.clearExecStopTime {
		var value time.Time
		t.ExecStopTime = value
		builder.SetNull(task.FieldExecStopTime)
	}
	if value := tuo.Content; value != nil {
		builder.Set(task.FieldContent, *value)
		t.Content = *value
	}
	if value := tuo.Output; value != nil {
		builder.Set(task.FieldOutput, *value)
		t.Output = *value
	}
	if tuo.clearOutput {
		var value string
		t.Output = value
		builder.SetNull(task.FieldOutput)
	}
	if value := tuo.Error; value != nil {
		builder.Set(task.FieldError, *value)
		t.Error = *value
	}
	if tuo.clearError {
		var value string
		t.Error = value
		builder.SetNull(task.FieldError)
	}
	if value := tuo.SessionID; value != nil {
		builder.Set(task.FieldSessionID, *value)
		t.SessionID = *value
	}
	if tuo.clearSessionID {
		var value string
		t.SessionID = value
		builder.SetNull(task.FieldSessionID)
	}
	if !builder.Empty() {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(tuo.removedTags) > 0 {
		eids := make([]int, len(tuo.removedTags))
		for eid := range tuo.removedTags {
			eids = append(eids, eid)
		}
		query, args := sql.Delete(task.TagsTable).
			Where(sql.InInts(task.TagsPrimaryKey[0], ids...)).
			Where(sql.InInts(task.TagsPrimaryKey[1], eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(tuo.tags) > 0 {
		values := make([][]int, 0, len(ids))
		for _, id := range ids {
			for eid := range tuo.tags {
				values = append(values, []int{id, eid})
			}
		}
		builder := sql.Insert(task.TagsTable).
			Columns(task.TagsPrimaryKey[0], task.TagsPrimaryKey[1])
		for _, v := range values {
			builder.Values(v[0], v[1])
		}
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if tuo.clearedJob {
		query, args := sql.Update(task.JobTable).
			SetNull(task.JobColumn).
			Where(sql.InInts(job.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(tuo.job) > 0 {
		for eid := range tuo.job {
			query, args := sql.Update(task.JobTable).
				Set(task.JobColumn, eid).
				Where(sql.InInts(task.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if tuo.clearedTarget {
		query, args := sql.Update(task.TargetTable).
			SetNull(task.TargetColumn).
			Where(sql.InInts(target.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(tuo.target) > 0 {
		for eid := range tuo.target {
			query, args := sql.Update(task.TargetTable).
				Set(task.TargetColumn, eid).
				Where(sql.InInts(task.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return t, nil
}
