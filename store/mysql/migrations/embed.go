package migrations

import "embed"

//go:embed *.sql
var LinkGenStoreMySQLMigrations embed.FS
