// +build ignore

package main

import (
	"bytes"
	"fmt"
	sfold "github.com/richardlehane/siegfried"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	sf, _ := sfold.Load("classic/classic.sig")
	var buf bytes.Buffer
	err := sf.SaveWriter(&buf)
	check(err)
	out, _ := os.Create("generated.go")
	fmt.Fprintf(out, "package sfclassic\n\nvar sfcontent = %#v", buf.Bytes())
	out.Close()
}
