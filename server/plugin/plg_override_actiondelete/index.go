package plg_override_actiondelete

import (
	"embed"

	. "github.com/mickael-kerjean/filestash/server/common"
)

//go:embed assets/*
var STATIC embed.FS

func init() {
	Hooks.Register.StaticPatch(STATIC)
}
