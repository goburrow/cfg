// Package cfg provides configuration file loaders for Go applications.
package cfg

// Loader is the configuration loader.
type Loader interface {
	Load(map[string]string) error
}

type overridingChain struct {
	loaders []Loader
}

func (l *overridingChain) Load(keys map[string]string) error {
	for _, l := range l.loaders {
		err := l.Load(keys)
		if err != nil {
			return err
		}
	}
	return nil
}

// NewOverridingChain returns a Loader which loads all configurations in the chain.
// Keys in later chain will override the same key from previous chain.
func NewOverridingChain(l ...Loader) Loader {
	return &overridingChain{l}
}
