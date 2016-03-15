package cfg

import (
	"errors"
	"path"
	"strings"
	"testing"
)

type staticLoader struct {
	name  string
	value string
}

func (s *staticLoader) Load(keys map[string]string) error {
	if s.name == "" {
		return errors.New("empty name")
	}
	keys[s.name] = s.value
	return nil
}

func TestOverridingChain(t *testing.T) {
	l := NewOverridingChain(
		&staticLoader{"k1", "v1"},
		&staticLoader{"k2", "v22"},
		&staticLoader{"k3", "v3"},
		&staticLoader{"k2", "v2"},
	)
	keys := make(map[string]string)
	err := l.Load(keys)
	if err != nil {
		t.Fatal(err)
	}
	if len(keys) != 3 {
		t.Fatalf("invalid length: %d", len(keys))
	}
	if keys["k1"] != "v1" || keys["k2"] != "v2" || keys["k3"] != "v3" {
		t.Fatalf("invalid keys: %+v", keys)
	}
}

func TestUserConfigPath(t *testing.T) {
	p := UserConfigPath("abc", "def")
	expected := path.Join(".config", "abc", "def")
	if !strings.HasSuffix(p, expected) {
		t.Fatalf("unexpected user config path: %s", p)
	}
}
