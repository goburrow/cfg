package ini

import (
	"fmt"
	"os"
)

// FileLoader loads ini from file. It flattens keys into the form: section.name = value.
type FileLoader struct {
	Path     string
	Required bool
}

func (l *FileLoader) Load(keys map[string]string) error {
	f, err := os.Open(l.Path)
	if err != nil {
		if os.IsNotExist(err) && l.Required {
			return err
		}
		return nil
	}
	defer f.Close()
	fn := func(s, n, v []byte) error {
		if len(s) > 0 {
			keys[fmt.Sprintf("%s.%s", s, n)] = string(v)
		} else {
			keys[string(n)] = string(v)
		}
		return nil
	}
	return Walk(f, fn)
}
