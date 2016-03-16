// Package cfg provides configuration file loaders for Go applications.
package cfg

import (
	"log"
	"os/user"
	"path"

	"github.com/goburrow/cfg/env"
	"github.com/goburrow/cfg/ini"
)

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

// DefaultLoader returns a chain of loading ini files from:
// 1. ~/.config/appName/config
// 2. System environments variable with appName_ prefix
func DefaultLoader(appName string) Loader {
	return NewOverridingChain(
		&ini.FileLoader{
			Path:     UserConfigPath(appName, "config"),
			Required: false,
		},
		&env.Loader{
			Prefix: appName + "_",
		},
	)
}

// UserConfigPath returns ~/.config/appName/fileName.
func UserConfigPath(appName, fileName string) string {
	u, err := user.Current()
	if err != nil {
		log.Println("could not get current user:", err)
		return path.Join(".config", appName, fileName)
	}
	return path.Join(u.HomeDir, ".config", appName, fileName)
}
