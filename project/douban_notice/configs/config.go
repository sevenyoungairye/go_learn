package configs

import "embed"

// Config using embed bind all config file.
//go:embed *
var Config embed.FS
