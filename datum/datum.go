package okatype_datum

import (
	"github.com/okanero/go-okatype/length"
	"github.com/okanero/go-okatype/magic"
	"github.com/okanero/go-okatype/version"
)

type Type struct {
	Magic        okatype_magic.Type       //  8 bytes \
	Version      okatype_version.Type     //  8 bytes |
	Length       okatype_length.Type      //  8 bytes |_ cache line
	                                      //  8 bytes |
	                                      // 32 bytes /
}
