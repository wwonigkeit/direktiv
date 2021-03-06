// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BucketSecretsColumns holds the columns for the "bucket_secrets" table.
	BucketSecretsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "ns", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "secret", Type: field.TypeBytes, Size: 65536},
		{Name: "type", Type: field.TypeInt},
	}
	// BucketSecretsTable holds the schema information for the "bucket_secrets" table.
	BucketSecretsTable = &schema.Table{
		Name:        "bucket_secrets",
		Columns:     BucketSecretsColumns,
		PrimaryKey:  []*schema.Column{BucketSecretsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BucketSecretsTable,
	}
)

func init() {
}
