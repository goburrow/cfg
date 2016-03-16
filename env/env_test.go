package env

import (
	"os"
	"testing"
)

func TestWalk(t *testing.T) {
	os.Clearenv()
	os.Setenv("name1", "value1")
	os.Setenv("name2", "value2")

	keys := make(map[string]string)
	err := Walk(func(k, v string) error {
		keys[k] = v
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(keys) != 2 || keys["name1"] != "value1" || keys["name2"] != "value2" {
		t.Fatalf("unexpected keys: %+v", keys)
	}
}

func TestLoader(t *testing.T) {
	os.Clearenv()
	os.Setenv("name1", "value1")
	os.Setenv("name2", "value2")

	keys := make(map[string]string)
	l := Loader{}

	err := l.Load(keys)
	if err != nil {
		t.Fatal(err)
	}
	if len(keys) != 2 || keys["name1"] != "value1" || keys["name2"] != "value2" {
		t.Fatalf("unexpected keys: %+v", keys)
	}
}

func TestLoaderWithPrefix(t *testing.T) {
	os.Clearenv()
	os.Setenv("pref_name1", "value1")
	os.Setenv("pre_name2", "value2")
	os.Setenv("pref_name3", "value3")

	keys := make(map[string]string)
	l := Loader{"pref_"}

	err := l.Load(keys)
	if err != nil {
		t.Fatal(err)
	}
	if len(keys) != 2 || keys["name1"] != "value1" || keys["name3"] != "value3" {
		t.Fatalf("unexpected keys: %+v", keys)
	}
}
