package okatype_datum

import (
	"github.com/okanero/go-okatype/agent"
	"github.com/okanero/go-okatype/blockhash"
	"github.com/okanero/go-okatype/count"
	"github.com/okanero/go-okatype/interaction"
	"github.com/okanero/go-okatype/nonce"
	"github.com/okanero/go-okatype/payload"
	"github.com/okanero/go-okatype/signature"
	"github.com/okanero/go-okatype/time"
)

type Type struct {
	Genesis okatype_blockhash.Type
	Prev    okatype_blockhash.Type
	Count   okatype_count.Type

	Lateral1 okatype_blockhash.Type
	Lateral2 okatype_blockhash.Type
	Lateral3 okatype_blockhash.Type

	Payload okatype_payload.Type

	Interaction     okatype_interaction.Type
	Time            okatype_time.Type
	Author          okatype_agent.Type
	AuthorSignature okatype_signature.Type
	Nonce           okatype_nonce.Type
}
