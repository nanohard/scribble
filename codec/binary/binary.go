// Package binary contains a codec to encode and decode entities in binary format
package binary

import (
	"github.com/kelindar/binary"
)

const name = "binary"

// Codec that encodes to and decodes from JSON.
var Codec = new(binaryCodec)

type binaryCodec int

func (c binaryCodec) Marshal(v interface{}) ([]byte, error) {
	return binary.Marshal(v)
}

func (c binaryCodec) Unmarshal(b []byte, v interface{}) error {
	return binary.Unmarshal(b, v)
}

func (c binaryCodec) Name() string {
	return name
}
