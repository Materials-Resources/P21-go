package models

import (
	"github.com/uptrace/bun"
	"time"
)

type SupplierTrade struct {
	bun.BaseModel    `bun:"table:supplier_trade"`
	SupplierTradeUid int32     `bun:"supplier_trade_uid,type:int,autoincrement,identity,pk"`        // Unique identifier for record.
	SupplierId       float64   `bun:"supplier_id,type:decimal(19,0),unique"`                        // Supplier ID for which this record pertains.
	NaftaTaxId       *string   `bun:"nafta_tax_id,type:varchar(255)"`                               // Pedimento/NAFTA Tax ID for this supplier.
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
