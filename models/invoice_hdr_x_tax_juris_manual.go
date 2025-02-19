package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type InvoiceHdrXTaxJurisManual struct {
	bun.BaseModel    `bun:"table:invoice_hdr_x_tax_juris_manual"`
	InvoiceNo        string          `bun:"invoice_no,type:varchar(10),pk"`                            // Invoice being paid
	JurisdictionId   string          `bun:"jurisdiction_id,type:varchar(10),pk"`                       // What is the unique identifier for this jurisdiction?
	Taxable          string          `bun:"taxable,type:char(1)"`                                      // Is this line item taxable?
	TaxPercentage    float64         `bun:"tax_percentage,type:decimal(19,8)"`                         // The tax percent applied to a particular invoice
	DateCreated      mssql.DateTime1 `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified mssql.DateTime1 `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy string          `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
