package models

import "github.com/uptrace/bun"

type Inventoryissuestests struct {
	bun.BaseModel           `bun:"table:InventoryIssuesTests"`
	Testuid                 int32   `bun:"TestUID,type:int"`
	Sql                     string  `bun:"SQL,type:varchar(8000)"`
	Orderby                 *string `bun:"OrderBy,type:varchar(255)"`
	Description             *string `bun:"Description,type:varchar(255)"`
	TotalTestUid            *int32  `bun:"total_test_uid,type:int"`
	InventoryIssuesTestsUid int32   `bun:"inventory_issues_tests_uid,type:int,autoincrement,identity"`
}
