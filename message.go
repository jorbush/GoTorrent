package peertopeerclient

import (
	"encoding/binary"
	"io"
)

type messageID uint8

const (
	MsgChoke         messageID = 0
	MsgUnchoke       messageID = 1
	MsgInterested    messageID = 2
	MsgNotInterested messageID = 3
	MsgHave          messageID = 4
	MsgBitfield      messageID = 5
	MsgRequest       messageID = 6
	MsgPiece         messageID = 7
	MsgCancel        messageID = 8
)

// Message stores ID and payload of a message
type Message struct {
	ID      messageID
	Payload []byte
}

// Serialize serializes a message into a buffer of the form
// <length prefix><message ID><payload>
// Interprets `nil` as a keep-alive message
func (m *Message) Serialize() []byte {
	// If the message is nil, return a keep-alive message,
	// which is simply 4 bytes of zeros
	if m == nil {
		return make([]byte, 4)
	}
	// Calculate the length of the message, which
	// includes the ID byte and optionally the payload
	length := uint32(1 + len(m.Payload)) // +1 for id
	// Create a buffer to store the message, with enough
	// space for the length prefix and message content
	buf := make([]byte, 4+length)
	// Convert the length to big-endian and place it
	// in the first 4 bytes of the buffer
	binary.BigEndian.PutUint32(buf[0:4], length)
	// Put the message ID in the next byte
	buf[4] = byte(m.ID)
	// Copy the payload into the buffer, starting
	// from the fifth byte
	copy(buf[5:], m.Payload)
	// Return the buffer containing the serialized message
	return buf
}

// Read parses a message from a stream. Returns `nil` on keep-alive message
func ReadMessage(r io.Reader) (*Message, error) {
	lengthBuf := make([]byte, 4)
	_, err := io.ReadFull(r, lengthBuf)
	if err != nil {
		return nil, err
	}
	length := binary.BigEndian.Uint32(lengthBuf)

	// keep-alive message
	if length == 0 {
		return nil, nil
	}

	messageBuf := make([]byte, length)
	_, err = io.ReadFull(r, messageBuf)
	if err != nil {
		return nil, err
	}

	m := Message{
		ID:      messageID(messageBuf[0]),
		Payload: messageBuf[1:],
	}

	return &m, nil
}
