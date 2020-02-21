package enthook

import "sync"

var (
	hooksMu sync.RWMutex
	hooks   = map[uintptr]interface{}{}
)

const expectedHooks = 1

// Register makes a database driver available by the provided name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func registerHooks(ptr uintptr, hook interface{}) {
	hooksMu.Lock()
	defer hooksMu.Unlock()
	if hook == nil {
		panic("sql: Register driver is nil")
	}
	if _, ok := hooks[ptr]; ok {
		return
	}
	hooks[ptr] = hook
}
