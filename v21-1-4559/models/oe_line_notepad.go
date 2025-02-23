package models

import (
	"github.com/uptrace/bun"
	"time"
)

type OeLineNotepad struct {
	bun.BaseModel    `bun:"table:oe_line_notepad"`
	NoteId           float64   `bun:"note_id,type:decimal(19,0),pk"`                                 // Unique ID for this supplier note.
	OrderNo          string    `bun:"order_no,type:varchar(8)"`                                      // Order number associated with this note.
	LineNo           float64   `bun:"line_no,type:decimal(19,0)"`                                    // Line number associated with this note.
	NotepadClassId   *string   `bun:"notepad_class_id,type:varchar(8)"`                              // Class value for this note.
	Topic            string    `bun:"topic,type:varchar(30)"`                                        // The topic of the note for the referenced area.
	Note             *string   `bun:"note,type:text"`                                                // Note text
	ActivationDate   time.Time `bun:"activation_date,type:datetime"`                                 // Date on which the note is activated.
	ExpirationDate   time.Time `bun:"expiration_date,type:datetime"`                                 // Date on which the note expires.
	EntryDate        time.Time `bun:"entry_date,type:datetime"`                                      // date the activity was entered
	Mandatory        string    `bun:"mandatory,type:char(1)"`                                        // Indicates whether the note will automatically present itself.
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`          // User who created the record
	TallyNoteFlag    *string   `bun:"tally_note_flag,type:char(1)"`                                  // Custom, indicate if the note is for tally measurement for tally item.
}
