package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type NotepadClass struct {
	bun.BaseModel    `bun:"table:notepad_class"`
	NotepadClassId   string          `bun:"notepad_class_id,type:varchar(8),pk"`                       // What is the unique identifier for this notepad cla
	NotepadClassDesc string          `bun:"notepad_class_desc,type:varchar(40)"`                       // What is this notepad class for?
	DeleteFlag       string          `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated      mssql.DateTime1 `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified mssql.DateTime1 `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy string          `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	CreatedBy        string          `bun:"created_by,type:varchar(255),default:(suser_sname())"`      // User who created the record
}
