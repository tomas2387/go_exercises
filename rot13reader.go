package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}
func (m rot13Reader) Read(b []byte) (int, error) {
	_, err := m.r.Read(b)
	if err != nil {
		return 0, err
	}

	for i:=0; i < len(b); i++ {
		if isUnknownChar(b[i]) {
			continue
		}

		b[i]+=13
		if isUnknownChar(b[i]) {
			b[i]-=26
		}
	}

	return len(b), nil;
}

func isUnknownChar(x byte) bool {
	return (x < 'A' || x > 'Z') && (x < 'a' || x > 'z')
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

