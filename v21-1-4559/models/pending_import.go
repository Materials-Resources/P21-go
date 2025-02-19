package models

import (
	"github.com/uptrace/bun"
	"time"
)

type PendingImport struct {
	bun.BaseModel         `bun:"table:pending_import"`
	PendingImportUid      int32     `bun:"pending_import_uid,type:int,pk"`                               // Unique Identifier
	ImpexpSourceId        string    `bun:"impexp_source_id,type:varchar(50)"`                            // Text identifier of the type of import this record is for.  See impexp_source for more.
	TransactionSetId      string    `bun:"transaction_set_id,type:varchar(32)"`                          // Identifies the transaction set. See transaction_set for more.
	DbtableId             string    `bun:"dbtable_id,type:varchar(40)"`                                  // Table that will be updated as part of the import of the transaction type.  See scheduled_import_deffor more.
	PendingImportSetNo    int32     `bun:"pending_import_set_no,type:int"`                               // Identifies the set of pending_import records that go together
	ImportData            string    `bun:"import_data,type:text"`                                        // Import data, can contain multiple rows
	StatusFlag            int32     `bun:"status_flag,type:int"`                                         // Identifies whether import is successful or not when the import was complete.
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	ImportedTransactionId *int32    `bun:"imported_transaction_id,type:int"`                             // transaction/entity ID generated by the import
	SourceReferenceNo     *string   `bun:"source_reference_no,type:varchar(255)"`                        // Value from import source source that identifies import set
	ImportResultInfo      *string   `bun:"import_result_info,type:varchar(8000)"`                        // Information from import to be used in B2B reply
}
