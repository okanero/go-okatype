/*
Package okatype_nonce (i.e., Okanero type length) provides the okatype_nonce.Type.

There are 2 different nonce that exist in an Okanero Message:

• #1: there is a nonce that in okatype_datum.Type, and

• #2: there is a nonce that in okatype_block.Type.

Each of these nonce are discovered by different agents:

• #1: the nonce in okatype_datum.Type is discovered by the author (agent), and

• #2: the nonce in okatype_block.Type is discovered by the endorser (agent).

In the context of the okatype_datum.Type and okatype_block.Type, the nonce is used to:

• help reduce the impact of certain kinds of attack,

• while at the same time offering a basis for providing (economic) value.

An author wishing to create Okanero Datum (okatype_datum.Type) must discover an appropriate nonce, before it request an
endorser to turn it into a valid Okanero Block (okatype_datum.Block).

Without an appropriate nonce, an Okanero Datum is NOT considered valid.
And endorsers will ignore the author's datum.

Similarly, an endorser wishing to have the Okanero Network accept its Okanero Block into an Okanero Blockchain must discover a separate
appropriate nonce.

Without an appropriate nonce, an Okanero Block is NOT considered valid.
And the Okanero Network will ignore the endorser's block.

MOST LIKELY YOU WOULD NOT CREATE YOUR OWN okatype_nonce.Type, BUT INSTEAD USE IT FROM AN okatype_datum.Type OR AN okatype_block.Type.


References

• "Hashcash - A Denial of Service Counter-Measure", by Adam Back (2002) http://www.hashcash.org/hashcash.pdf

*/
package okatype_nonce
