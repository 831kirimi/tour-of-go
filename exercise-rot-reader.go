package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(w []byte) (int, error){
	n, err := rot.r.Read(w)
	for i := 0; i < len(w); i++{
		if (w[i] >= 'a' && w[i] <= 'm') || (w[i] >= 'A' && w[i] <= 'M'){
		 	w[i] += 13
		}else{
			w[i] -= 13
		}
	}
	return n, err
}




func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

