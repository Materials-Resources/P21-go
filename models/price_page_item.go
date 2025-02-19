package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type PricePageItem struct {
	bun.BaseModel    `bun:"table:price_page_item"`
	PricePageItemUid int32           `bun:"price_page_item_uid,type:int,autoincrement,identity,pk"` // Unique ID for this price_page_item record.
	PricePageUid     int32           `bun:"price_page_uid,type:int"`                                // Unique ID for Price Pages.
	InvMastUid       int32           `bun:"inv_mast_uid,type:int"`                                  // Unique identifier for the item id.
	EffectiveDate    mssql.DateTime1 `bun:"effective_date,type:datetime"`                           // Starting date on which pricing page will be used in pricing calculations.
	ExpirationDate   mssql.DateTime1 `bun:"expiration_date,type:datetime"`                          // Date on which pricing page is no longer used in pricing calculations.
	DateCreated      mssql.DateTime1 `bun:"date_created,type:datetime"`                             // Indicates the date/time this record was created.
	DateLastModified mssql.DateTime1 `bun:"date_last_modified,type:datetime"`                       // Indicates the date/time this record was last modified.
	LastMaintainedBy string          `bun:"last_maintained_by,type:varchar(30)"`                    // ID of the user who last maintained this record
}
