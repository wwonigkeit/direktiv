// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// NamespacesColumns holds the columns for the "namespaces" table.
	NamespacesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 64},
		{Name: "created", Type: field.TypeTime},
		{Name: "key", Type: field.TypeBytes},
	}
	// NamespacesTable holds the schema information for the "namespaces" table.
	NamespacesTable = &schema.Table{
		Name:        "namespaces",
		Columns:     NamespacesColumns,
		PrimaryKey:  []*schema.Column{NamespacesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ServersColumns holds the columns for the "servers" table.
	ServersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "ip", Type: field.TypeString, Unique: true},
		{Name: "ext_ip", Type: field.TypeString},
		{Name: "nats_port", Type: field.TypeInt},
		{Name: "member_port", Type: field.TypeInt},
		{Name: "added", Type: field.TypeTime},
	}
	// ServersTable holds the schema information for the "servers" table.
	ServersTable = &schema.Table{
		Name:        "servers",
		Columns:     ServersColumns,
		PrimaryKey:  []*schema.Column{ServersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// SubroutinesColumns holds the columns for the "subroutines" table.
	SubroutinesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "caller_id", Type: field.TypeString, Unique: true},
		{Name: "semaphore", Type: field.TypeInt},
		{Name: "memory", Type: field.TypeString},
		{Name: "subroutine_ids", Type: field.TypeJSON},
		{Name: "subroutine_responses", Type: field.TypeJSON, Nullable: true},
	}
	// SubroutinesTable holds the schema information for the "subroutines" table.
	SubroutinesTable = &schema.Table{
		Name:        "subroutines",
		Columns:     SubroutinesColumns,
		PrimaryKey:  []*schema.Column{SubroutinesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// TimersColumns holds the columns for the "timers" table.
	TimersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "fn", Type: field.TypeString},
		{Name: "cron", Type: field.TypeString, Nullable: true},
		{Name: "one", Type: field.TypeTime, Nullable: true},
		{Name: "data", Type: field.TypeBytes, Nullable: true},
	}
	// TimersTable holds the schema information for the "timers" table.
	TimersTable = &schema.Table{
		Name:        "timers",
		Columns:     TimersColumns,
		PrimaryKey:  []*schema.Column{TimersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// WorkflowsColumns holds the columns for the "workflows" table.
	WorkflowsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "created", Type: field.TypeTime},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 1024, Default: ""},
		{Name: "active", Type: field.TypeBool, Default: true},
		{Name: "revision", Type: field.TypeInt},
		{Name: "workflow", Type: field.TypeBytes},
		{Name: "namespace_workflows", Type: field.TypeString, Nullable: true, Size: 64},
	}
	// WorkflowsTable holds the schema information for the "workflows" table.
	WorkflowsTable = &schema.Table{
		Name:       "workflows",
		Columns:    WorkflowsColumns,
		PrimaryKey: []*schema.Column{WorkflowsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "workflows_namespaces_workflows",
				Columns: []*schema.Column{WorkflowsColumns[7]},

				RefColumns: []*schema.Column{NamespacesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "workflow_name_namespace_workflows",
				Unique:  true,
				Columns: []*schema.Column{WorkflowsColumns[1], WorkflowsColumns[7]},
			},
		},
	}
	// WorkflowEventsColumns holds the columns for the "workflow_events" table.
	WorkflowEventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "events", Type: field.TypeJSON},
		{Name: "correlations", Type: field.TypeJSON},
		{Name: "signature", Type: field.TypeBytes, Nullable: true},
		{Name: "count", Type: field.TypeInt},
		{Name: "workflow_wfevents", Type: field.TypeUUID, Nullable: true},
	}
	// WorkflowEventsTable holds the schema information for the "workflow_events" table.
	WorkflowEventsTable = &schema.Table{
		Name:       "workflow_events",
		Columns:    WorkflowEventsColumns,
		PrimaryKey: []*schema.Column{WorkflowEventsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "workflow_events_workflows_wfevents",
				Columns: []*schema.Column{WorkflowEventsColumns[5]},

				RefColumns: []*schema.Column{WorkflowsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// WorkflowEventsWaitsColumns holds the columns for the "workflow_events_waits" table.
	WorkflowEventsWaitsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "events", Type: field.TypeJSON},
		{Name: "workflow_events_wfeventswait", Type: field.TypeInt, Nullable: true},
	}
	// WorkflowEventsWaitsTable holds the schema information for the "workflow_events_waits" table.
	WorkflowEventsWaitsTable = &schema.Table{
		Name:       "workflow_events_waits",
		Columns:    WorkflowEventsWaitsColumns,
		PrimaryKey: []*schema.Column{WorkflowEventsWaitsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "workflow_events_waits_workflow_events_wfeventswait",
				Columns: []*schema.Column{WorkflowEventsWaitsColumns[2]},

				RefColumns: []*schema.Column{WorkflowEventsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// WorkflowInstancesColumns holds the columns for the "workflow_instances" table.
	WorkflowInstancesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "instance_id", Type: field.TypeString, Unique: true},
		{Name: "invoked_by", Type: field.TypeString},
		{Name: "status", Type: field.TypeString},
		{Name: "revision", Type: field.TypeInt},
		{Name: "begin_time", Type: field.TypeTime},
		{Name: "end_time", Type: field.TypeTime, Nullable: true},
		{Name: "flow", Type: field.TypeJSON, Nullable: true},
		{Name: "input", Type: field.TypeString},
		{Name: "output", Type: field.TypeString, Nullable: true},
		{Name: "state_data", Type: field.TypeString, Nullable: true},
		{Name: "memory", Type: field.TypeString, Nullable: true},
		{Name: "deadline", Type: field.TypeTime, Nullable: true},
		{Name: "attempts", Type: field.TypeInt, Nullable: true},
		{Name: "error_code", Type: field.TypeString, Nullable: true},
		{Name: "error_message", Type: field.TypeString, Nullable: true},
		{Name: "workflow_instances", Type: field.TypeUUID, Nullable: true},
	}
	// WorkflowInstancesTable holds the schema information for the "workflow_instances" table.
	WorkflowInstancesTable = &schema.Table{
		Name:       "workflow_instances",
		Columns:    WorkflowInstancesColumns,
		PrimaryKey: []*schema.Column{WorkflowInstancesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "workflow_instances_workflows_instances",
				Columns: []*schema.Column{WorkflowInstancesColumns[16]},

				RefColumns: []*schema.Column{WorkflowsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		NamespacesTable,
		ServersTable,
		SubroutinesTable,
		TimersTable,
		WorkflowsTable,
		WorkflowEventsTable,
		WorkflowEventsWaitsTable,
		WorkflowInstancesTable,
	}
)

func init() {
	WorkflowsTable.ForeignKeys[0].RefTable = NamespacesTable
	WorkflowEventsTable.ForeignKeys[0].RefTable = WorkflowsTable
	WorkflowEventsWaitsTable.ForeignKeys[0].RefTable = WorkflowEventsTable
	WorkflowInstancesTable.ForeignKeys[0].RefTable = WorkflowsTable
}