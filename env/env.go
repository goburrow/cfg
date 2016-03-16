package env

import (
	"os"
	"strings"
)

// WalkFunc is the type of the function called for each key visited by Walk.
type WalkFunc func(name, value string) error

// Walk walks through all environment variables.
func Walk(fn WalkFunc) error {
	envs := os.Environ()
	for _, env := range envs {
		i := strings.Index(env, "=")
		if i >= 0 {
			err := fn(env[:i], env[i+1:])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Loader loads environment variables to given map. Only names starting with Prefix
// are added (without Prefix).
type Loader struct {
	Prefix string
}

func (l *Loader) Load(keys map[string]string) error {
	fn := func(name, value string) error {
		if l.Prefix == "" {
			keys[name] = value
		} else if strings.HasPrefix(name, l.Prefix) {
			keys[name[len(l.Prefix):]] = value
		}
		return nil
	}
	return Walk(fn)
}
