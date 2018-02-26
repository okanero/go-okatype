package okatype_block

import (
	"github.com/okanero/go-okatype/datum"
	"github.com/okanero/go-okatype/nonce"
	"github.com/okanero/go-okatype/trace"
)

type Type struct {
	Datum  okatype_datum.Type
	
	Trace  okatype_trace.Type
	Time   okatype_time.Type
	Nonce  okatype_nonce.Type
}