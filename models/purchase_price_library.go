package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type PurchasePriceLibrary struct {
	bun.BaseModel            `bun:"table:purchase_price_library"`
	CompanyId                string          `bun:"company_id,type:varchar(8),pk"`                             // Unique code that identifies a company.
	PurchasePriceLibraryId   string          `bun:"purchase_price_library_id,type:varchar(8),pk"`              // A unique purchase pricing library id
	PurchasePriceLibraryDesc string          `bun:"purchase_price_library_desc,type:varchar(40)"`              // A description of the purchase pricing library
	PurchasePriceLibraryType string          `bun:"purchase_price_library_type,type:varchar(40)"`              // Indicates how to search the pricing books for the correct price, or whether to use the source price & multiplier.
	SourcePrice              *string         `bun:"source_price,type:varchar(40)"`                             // When the library_type is Multiplier, this indicates the base price to use.
	Multiplier               *float64        `bun:"multiplier,type:decimal(19,9)"`                             // When the library_type is Multiplier, this indicates what to multiply the source_price by to get the purchase price.
	Inactive                 string          `bun:"inactive,type:char(1)"`                                     // Indicates if this commission rule is inactive with
	DeleteFlag               string          `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated              mssql.DateTime1 `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified         mssql.DateTime1 `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy         string          `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
