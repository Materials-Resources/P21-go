package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type Journal struct {
	bun.BaseModel      `bun:"table:journal"`
	JournalId          string          `bun:"journal_id,type:varchar(2),pk"`                             // What is the journal to be used for this repeating journal entry?
	JournalDescription string          `bun:"journal_description,type:varchar(20)"`                      // How would you describe this journal?
	DeleteFlag         string          `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated        mssql.DateTime1 `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified   mssql.DateTime1 `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy   string          `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	JournalType        int32           `bun:"journal_type,type:int"`                                     // Indicates the type  - from code_p21 table (User Defined, System Defined, Roll Up)
}
