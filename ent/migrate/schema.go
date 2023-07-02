// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// FmsFilesColumns holds the columns for the "fms_files" table.
	FmsFilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Default: 1},
		{Name: "name", Type: field.TypeString},
		{Name: "file_type", Type: field.TypeUint8},
		{Name: "size", Type: field.TypeUint64},
		{Name: "path", Type: field.TypeString},
		{Name: "user_uuid", Type: field.TypeString},
		{Name: "md5", Type: field.TypeString},
	}
	// FmsFilesTable holds the schema information for the "fms_files" table.
	FmsFilesTable = &schema.Table{
		Name:       "fms_files",
		Columns:    FmsFilesColumns,
		PrimaryKey: []*schema.Column{FmsFilesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "file_user_uuid",
				Unique:  false,
				Columns: []*schema.Column{FmsFilesColumns[8]},
			},
			{
				Name:    "file_file_type",
				Unique:  false,
				Columns: []*schema.Column{FmsFilesColumns[5]},
			},
		},
	}
	// FmsTagsColumns holds the columns for the "fms_tags" table.
	FmsTagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Default: 1},
		{Name: "name", Type: field.TypeString, Comment: "Tag's name | 标签名称"},
		{Name: "remark", Type: field.TypeString, Comment: "The remark of tag | 标签的备注"},
	}
	// FmsTagsTable holds the schema information for the "fms_tags" table.
	FmsTagsTable = &schema.Table{
		Name:       "fms_tags",
		Columns:    FmsTagsColumns,
		PrimaryKey: []*schema.Column{FmsTagsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "tag_name",
				Unique:  false,
				Columns: []*schema.Column{FmsTagsColumns[4]},
			},
		},
	}
	// TagFilesColumns holds the columns for the "tag_files" table.
	TagFilesColumns = []*schema.Column{
		{Name: "tag_id", Type: field.TypeUint64},
		{Name: "file_id", Type: field.TypeUUID},
	}
	// TagFilesTable holds the schema information for the "tag_files" table.
	TagFilesTable = &schema.Table{
		Name:       "tag_files",
		Columns:    TagFilesColumns,
		PrimaryKey: []*schema.Column{TagFilesColumns[0], TagFilesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tag_files_tag_id",
				Columns:    []*schema.Column{TagFilesColumns[0]},
				RefColumns: []*schema.Column{FmsTagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "tag_files_file_id",
				Columns:    []*schema.Column{TagFilesColumns[1]},
				RefColumns: []*schema.Column{FmsFilesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		FmsFilesTable,
		FmsTagsTable,
		TagFilesTable,
	}
)

func init() {
	FmsFilesTable.Annotation = &entsql.Annotation{
		Table: "fms_files",
	}
	FmsTagsTable.Annotation = &entsql.Annotation{
		Table: "fms_tags",
	}
	TagFilesTable.ForeignKeys[0].RefTable = FmsTagsTable
	TagFilesTable.ForeignKeys[1].RefTable = FmsFilesTable
}
