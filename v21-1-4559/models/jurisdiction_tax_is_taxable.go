package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type JurisdictionTaxIsTaxable struct {
	bun.BaseModel        `bun:"table:jurisdiction_tax_is_taxable"`
	JurisdictionId       string          `bun:"jurisdiction_id,type:varchar(10),pk"`                       // Identifier of the jurisdication (ie county) in which to pay tax.
	TaxingJurisdictionId string          `bun:"taxing_jurisdiction_id,type:varchar(10),pk"`                // Identifier of the jurisdication (ie state) that the jurisdication_id is part of.
	DeleteFlag           string          `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated          mssql.DateTime1 `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified     mssql.DateTime1 `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy     string          `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
