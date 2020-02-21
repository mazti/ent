package hooks

import (
	"reflect"
	_ "unsafe"

	"github.com/facebookincubator/ent/entc/integration/enthook/schema"
)

//go:linkname registerHooks github.com/facebookincubator/ent/entc/integration/enthook.registerHooks
func registerHooks(ptr uintptr, hook interface{})

func init() {
	card := schema.Card{}
	hooks := card.Hooks()
	for _, h := range hooks {
		registerHooks(reflect.ValueOf(h).Pointer(), h)
	}
}
