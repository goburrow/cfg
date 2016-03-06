package ini

import (
	"fmt"
	"os"
)

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
		var name string
		if len(s) > 0 {
			name = fmt.Sprintf("%s.%s", s, n)
		} else {
			name = string(n)
		}
		keys[name] = string(v)
		return nil
	}
	return Walk(f, fn)
}
