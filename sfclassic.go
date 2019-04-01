package sfclassic

import (
	"bytes"
	"github.com/richardlehane/siegfried"
	"github.com/richardlehane/siegfried/pkg/core"
	"io"
)

type Siegfried struct {
	*siegfried.Siegfried
}

func New() *Siegfried {
	buf := bytes.NewBuffer(sfcontent)
	sf, _ := siegfried.LoadReader(buf)
	return &Siegfried{sf}
}

func (sf *Siegfried) Identify(rdr io.Reader, name string) ([]core.Identification, error) {
	return sf.Siegfried.Identify(rdr, name, "")
}
