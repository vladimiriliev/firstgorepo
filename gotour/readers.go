package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read (x []byte) (int, error){
	for i:=0; i < len(x); i++ {
		x[i]='A'
	}
	return len(x), nil
}

func main() {
	reader.Validate(MyReader{})
}