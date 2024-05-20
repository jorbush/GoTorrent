package main

// A Bitfield represents the pieces that a peer has
type Bitfield []byte

// HasPiece tells if a bitfield has a particular index set
func (bf Bitfield) HasPiece(index int) bool {
	// Calculate the byte index where the bit resides
	byteIndex := index / 8
	// Calculate the bit's offset within the byte
	offset := index % 8
	// Shift the byte to the right by (7 - offset) bits
	// and check if the least significant bit is set
	return bf[byteIndex]>>(7-offset)&1 != 0
}

// SetPiece sets a bit in the bitfield
func (bf Bitfield) SetPiece(index int) {
	byteIndex := index / 8
	offset := index % 8
	bf[byteIndex] |= 1 << (7 - offset)
}
