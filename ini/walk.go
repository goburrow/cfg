// Package ini provides helpers for reading ini file
package ini

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

// WalkFunc is the type of the function called for each key visited by Walk.
type WalkFunc func(section, name, value []byte) error

// Walk walks all items from r, calling walkFn for each item.
func Walk(r io.Reader, walkFn WalkFunc) error {
	var section []byte
	sep := []byte{'='}

	scanner := bufio.NewScanner(r)
	for lineNo := 1; scanner.Scan(); lineNo++ {
		line := bytes.TrimSpace(scanner.Bytes())
		if len(line) == 0 || line[0] == '#' || line[0] == ';' {
			// Skip comments
			continue
		}
		if line[0] == '[' && line[len(line)-1] == ']' {
			// Section
			section = duplicate(section, bytes.TrimSpace(line[1:len(line)-1]))
			continue
		}
		if line[len(line)-1] == '\\' {
			// Duplicate data from Scanner to append next lines
			line = append([]byte(nil), line[:len(line)-1]...)

			// Continuation lines
			for scanner.Scan() {
				lineNo++
				l := bytes.TrimSpace(scanner.Bytes())
				if l[len(l)-1] == '\\' {
					line = append(line, l[:len(l)-1]...)
				} else {
					line = append(line, l...)
					break
				}
			}
		}
		key := bytes.SplitN(line, sep, 2)
		if len(key) != 2 {
			return fmt.Errorf("missing delimiter (line %d)", lineNo)
		}
		name := bytes.TrimSpace(key[0])
		if len(name) == 0 {
			return fmt.Errorf("empty name (line %d)", lineNo)
		}
		value := bytes.TrimSpace(key[1])
		err := walkFn(section, name, value)
		if err != nil {
			return err
		}
	}
	return nil
}

// duplicate copies data from src to dest. It returns a new slice if dest capacity does not
// fit src length.
func duplicate(dest, src []byte) []byte {
	if cap(dest) < len(src) {
		return append([]byte(nil), src...)
	}
	dest = dest[:len(src)]
	copy(dest, src)
	return dest
}
