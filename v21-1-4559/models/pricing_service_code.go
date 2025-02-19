package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type PricingServiceCode struct {
	bun.BaseModel    `bun:"table:pricing_service_code"`
	CodeId           float64         `bun:"code_id,type:decimal(19,0),pk"`                             // FK to pricing_service_code.code_id
	CodeName         string          `bun:"code_name,type:varchar(10)"`                                // Name of the code
	CodeDesc         *string         `bun:"code_desc,type:varchar(254)"`                               // Description of the code
	DeleteFlag       string          `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated      mssql.DateTime1 `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified mssql.DateTime1 `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy string          `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
