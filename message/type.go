package okatype_message

import (
	"github.com/okanero/go-okatype/block"
	"github.com/okanero/go-okatype/length"
	"github.com/okanero/go-okatype/magic"
	"github.com/okanero/go-okatype/network"
	"github.com/okanero/go-okatype/version"
)

type Type struct {
	Magic   okatype_magic.Type
	Version okatype_version.Type
	Network okatype_network.Type
	Length  okatype_length.Type

	Block   okatype_block.Type
}
