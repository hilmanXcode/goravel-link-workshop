package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"goravel/app/facades"
)

type M20261125035649CreateLinksTable struct{}

// Signature The unique signature for the migration.
func (r *M20261125035649CreateLinksTable) Signature() string {
	return "20261125035649_create_links_table"
}

// Up Run the migrations.
func (r *M20261125035649CreateLinksTable) Up() error {
	if !facades.Schema().HasTable("links") {
		return facades.Schema().Create("links", func(table schema.Blueprint) {
			table.ID()
			table.String("code", 10)
			table.Text("url")
			table.UnsignedInteger("clicks").Default(0)
			table.TimestampsTz()
			table.Unique("code")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20261125035649CreateLinksTable) Down() error {
	return facades.Schema().DropIfExists("links")
}
