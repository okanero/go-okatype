/*
Package okatype_nonce (i.e., Okanero type length) provides the okatype_nonce.Type.

In the context of the okatype_datum.Type and okatype_block.Type, the nonce is used to:

• help reduce the impact of certain kinds of attack,

• while at the same time offering a basis for providing (economic) value.

An author wishing to create Okanero Datum (okatype_datum.Type) must discover an appropriate nonce, before it request an
endorser to turn it into a valid Okanero Block (okatype_datum.Block).

Without an appropriate nonce, an Okanero Datum is NOT considered valid.
And endorsers will ignore the author's datum.

Similarly, an endorser wishing to have the Okanero Network accept its Okanero Block into a Okanero Blockchain must discover a separate
appropriate nonce.

Without an appropriate nonce, an Okanero Block is NOT considered valid.
And the Okanero Network will ignore the endorser's block.

MOST LIKELY YOU WOULD NOT CREATE YOUR OWN okatype_nonce.Type, BUT INSTEAD USE IT FROM AN okatype_datum.Type OR A okatype_block.Type.
*/
package okatype_nonce
