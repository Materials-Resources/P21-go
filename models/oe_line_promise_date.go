package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type OeLinePromiseDate struct {
	bun.BaseModel                   `bun:"table:oe_line_promise_date"`
	OeLineUid                       int32            `bun:"oe_line_uid,type:int"`                                         // Associated oe_line record
	OriginalPromiseDate             *mssql.DateTime1 `bun:"original_promise_date,type:datetime"`                          // Original promise date
	PromiseDate                     mssql.DateTime1  `bun:"promise_date,type:datetime"`                                   // Current promise date
	EditedFlag                      string           `bun:"edited_flag,type:char(1)"`                                     // Indicates whether promise_date was edited
	AroDays                         *int32           `bun:"aro_days,type:int"`                                            // Number of days after receipt of order
	DateCreated                     mssql.DateTime1  `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                       string           `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified                mssql.DateTime1  `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy                string           `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	OeLinePromiseDateUid            int32            `bun:"oe_line_promise_date_uid,type:int,pk"`
	PromiseDateEditedDate           *mssql.DateTime1 `bun:"promise_date_edited_date,type:datetime"` // Date the promise_date field was last edited
	ExtendedDesc                    *string          `bun:"extended_desc,type:varchar(8000)"`       // Extended description for promise dates
	Alert                           *string          `bun:"alert,type:varchar(1)"`                  // alert flag/toggle
	PreviousPromiseDate             *mssql.DateTime1 `bun:"previous_promise_date,type:datetime"`
	InsufficientDataToCalculateFlag *string          `bun:"insufficient_data_to_calculate_flag,type:char(1)"` // flag to know if the promise date has been previously succesfully calculated
	Edi855sResendPromiseDate        *mssql.DateTime1 `bun:"edi855s_resend_promise_date,type:datetime"`        // Custom column to store the promise date at OrderAck EDI855s resend time.
}
