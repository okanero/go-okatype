package okatype_content

import (
	"github.com/okanero/go-okatype/length"
	"github.com/okanero/go-okatype/pascalstring"

	"io"
)

type Type struct {
	Magic   okatype_contentmagic.Type
	Name    okatype_pascalstring.Type
	Version okatype_pascalstring.Type

	Type    okatype_pascalstring.Type
	Length  okatype_length.Type
	Data    interface{io.WriterTo ; io.ReaderFrom}
}
