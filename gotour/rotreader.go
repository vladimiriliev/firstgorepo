package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read (x []byte) (int, error){
	resultSize, ok := reader.r.Read(x)
	if (ok == io.EOF){
		if (resultSize == 0){
			return 0, io.EOF;
		} else {
			for i, value := range x {
				x[i] = rot13convert(value);
			}
			return len(x), io.EOF;
		}
	} else if (ok!=nil){
		return 0, ok;
	} else {
		for i, value := range x {
			x[i] = rot13convert(value);
		}
		return len(x), nil;
	}
}

func rot13convert (c byte) byte{
	if 'a' <= c && c <= 'z'{
		return (((c - 'a') + 13) % 26) + 'a'
	} else if 'A' <= c && c <= 'Z'{
		return (((c - 'A') + 13) % 26) + 'A'
	}
	return c
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}