package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (rot13 *rot13Reader) Read(p []byte) (n int, err error) {
	rot_n, rot_err := rot13.r.Read(p)

	if rot_err != nil {
		return rot_n, rot_err
	}

	for i := 0; i < rot_n; i++ {
		switch {
		case p[i] >= 'A' && p[i] <= 'Z':
			p[i] = 'A' + (p[i] - 'A' + 13) % 26
		case p[i] >= 'a' && p[i] <= 'z':
			p[i] = 'a' + (p[i] - 'a' + 13) % 26
		}
	}

	return rot_n, nil
}

func main() {
    s := strings.NewReader(
        "Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
