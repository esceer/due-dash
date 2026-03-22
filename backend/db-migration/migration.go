//go:build linux || darwin
// +build linux darwin

package dbmigration

import "embed"

//go:embed db/migrations/*.sql
var MigrationsFs embed.FS
