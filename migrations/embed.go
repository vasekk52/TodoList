package migrations

import "embed"

// FS holds the SQL migration files, embedded into the binary so they ship
// inside the Docker image and get applied automatically on startup.
//
//go:embed *.sql
var FS embed.FS
