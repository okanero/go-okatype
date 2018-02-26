package okatype_datum

import (
	"github.com/okanero/go-okatype/agent"
	"github.com/okanero/go-okatype/count"
	"github.com/okanero/go-okatype/hash"
	"github.com/okanero/go-okatype/interaction"
	"github.com/okanero/go-okatype/length"
	"github.com/okanero/go-okatype/magic"
	"github.com/okanero/go-okatype/nonce"
	"github.com/okanero/go-okatype/payload"
	"github.com/okanero/go-okatype/signature"
	"github.com/okanero/go-okatype/time"
	"github.com/okanero/go-okatype/version"
)

type Type struct {
	Genesis okatype_hash.Type
	Prev    okatype_hash.Type
	Count   okatype_count.Type

	Lateral1 okatype_hash.Type
	Lateral2 okatype_hash.Type
	Lateral3 okatype_hash.Type

	Payload okatype_payload.Type

	Interaction     okatype_interaction.Type
	Time            okatype_time.Type
	Author          okatype_agent.Type
	AuthorSignature okatype_signature.Type
	Nonce           okatype_nonce.Type
}
