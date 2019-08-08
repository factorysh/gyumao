package gyumao

import (
	"github.com/factorysh/gyumao/plugin"
)

// Gyumao main object
type Gyumao struct {
	Plugins *plugin.Plugins
}

// New Gyumao instance
func New() *Gyumao {
	return &Gyumao{
		Plugins: plugin.NewPlugins(),
	}
}