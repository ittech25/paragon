// Code generated by entc, DO NOT EDIT.

package tag

import (
	"github.com/kcarretto/paragon/ent/schema"
)

const (
	// Label holds the string label denoting the tag type in the database.
	Label = "tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"

	// Table holds the table name of the tag in the database.
	Table = "tags"
	// TargetsTable is the table the holds the targets relation/edge. The primary key declared below.
	TargetsTable = "target_tags"
	// TargetsInverseTable is the table name for the Target entity.
	// It exists in this package in order to avoid circular dependency with the "target" package.
	TargetsInverseTable = "targets"
	// TasksTable is the table the holds the tasks relation/edge. The primary key declared below.
	TasksTable = "task_tags"
	// TasksInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TasksInverseTable = "tasks"
	// JobsTable is the table the holds the jobs relation/edge. The primary key declared below.
	JobsTable = "job_tags"
	// JobsInverseTable is the table name for the Job entity.
	// It exists in this package in order to avoid circular dependency with the "job" package.
	JobsInverseTable = "jobs"
)

// Columns holds all SQL columns are tag fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// TargetsPrimaryKey and TargetsColumn2 are the table columns denoting the
	// primary key for the targets relation (M2M).
	TargetsPrimaryKey = []string{"target_id", "tag_id"}
	// TasksPrimaryKey and TasksColumn2 are the table columns denoting the
	// primary key for the tasks relation (M2M).
	TasksPrimaryKey = []string{"task_id", "tag_id"}
	// JobsPrimaryKey and JobsColumn2 are the table columns denoting the
	// primary key for the jobs relation (M2M).
	JobsPrimaryKey = []string{"job_id", "tag_id"}
)

var (
	fields = schema.Tag{}.Fields()

	// descName is the schema descriptor for Name field.
	descName = fields[0].Descriptor()
	// NameValidator is a validator for the "Name" field. It is called by the builders before save.
	NameValidator = descName.Validators[0].(func(string) error)
)