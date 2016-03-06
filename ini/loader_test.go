package ini

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileLoader(t *testing.T) {
	content := []byte(`a=1
# X
[x]
b=2
# Y
[y]
c=3`)
	tmpf, err := ioutil.TempFile("", "cfg-ini")
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tmpf.Name()) // clean up

	_, err = tmpf.Write(content)
	if err != nil {
		t.Fatal(err)
	}
	err = tmpf.Close()
	if err != nil {
		t.Fatal(err)
	}

	l := FileLoader{tmpf.Name(), false}
	keys := make(map[string]string)
	err = l.Load(keys)
	if err != nil {
		t.Fatal(err)
	}
	if len(keys) != 3 {
		t.Fatalf("unexpected length: %d", len(keys))
	}
	if keys["a"] != "1" || keys["x.b"] != "2" || keys["y.c"] != "3" {
		t.Fatalf("unexpected keys: %+v", keys)
	}
}

func TestFileLoaderRequired(t *testing.T) {
	l := FileLoader{"/not/exist", false}
	keys := make(map[string]string)
	err := l.Load(keys)
	if err != nil {
		t.Fatal(err)
	}
	if len(keys) != 0 {
		t.Fatalf("unexpected length: %d", len(keys))
	}
	l.Required = true
	err = l.Load(keys)
	if !os.IsNotExist(err) {
		t.Fatalf("unexpected error: %#v", err)
	}
}
