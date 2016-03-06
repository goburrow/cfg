package ini

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestWalk(t *testing.T) {
	data := `
k1=v1
# Comment 1
 k2 = v2 

; Comment 2
[s1]
k3  =v3
    
[ s22 ]
k4=  v4=

[s3]
k5=v5`
	var results bytes.Buffer
	err := Walk(strings.NewReader(data), func(s, k, v []byte) error {
		fmt.Fprintf(&results, "%s.%s=%s\n", s, k, v)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	expected := ".k1=v1\n.k2=v2\ns1.k3=v3\n s22 .k4=v4=\ns3.k5=v5\n"
	actual := results.String()
	if expected != actual {
		t.Fatalf("unexpected results: %s", actual)
	}
}

func BenchmarkWalk(b *testing.B) {
	data := `
[s1]
k1=v1
[s22]
k22 = v22
[s333]
k333=v333
# Comment
[s4444]
k4444 = v4444
[s5]
k5 = v5
`
	r := strings.NewReader(data)
	f := func(s, k, v []byte) error {
		return nil
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Walk(r, f)
		r.Seek(0, 0)
	}
}
