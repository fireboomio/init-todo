package embeds

import "embed"

const INTROSPECT_FILE = "introspect.json"

//go:embed introspect.json
var EmbedIntrospect embed.FS
