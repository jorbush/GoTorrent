package core

import (
	"bytes"
	"fmt"
	"gotorrent/bittorrent/network"
	"gotorrent/bittorrent/utils"
	"log"
	"net"
	"time"
)

// A Client is a TCP connection with a peer
type Client struct {
	Conn     net.Conn
	Choked   bool
	Bitfield utils.Bitfield
	peer     network.Peer
	infoHash [20]byte
	peerID   [20]byte
}

func completeHandshake(conn net.Conn, infohash, peerID [20]byte) (*network.Handshake, error) {
	if err := conn.SetDeadline(time.Now().Add(3 * time.Second)); err != nil {
		log.Printf("Failed to set deadline: %v", err)
	}

	defer func() {
		if err := conn.SetDeadline(time.Time{}); err != nil {
			log.Printf("Failed to reset deadline: %v", err)
		}
	}()

	req := network.NewHandshake(infohash, peerID)
	_, err := conn.Write(req.Serialize())
	if err != nil {
		return nil, err
	}

	res, err := network.ReadHandshake(conn)
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(res.InfoHash[:], infohash[:]) {
		return nil, fmt.Errorf("expected infohash %x but got %x", res.InfoHash, infohash)
	}
	return res, nil
}

func recvBitfield(conn net.Conn) (utils.Bitfield, error) {
	if err := conn.SetDeadline(time.Now().Add(5 * time.Second)); err != nil {
		log.Printf("Failed to set deadline: %v", err)
	}
	defer func() {
		if err := conn.SetDeadline(time.Time{}); err != nil {
			log.Printf("Failed to reset deadline: %v", err)
		}
	}() // Disable the deadline

	msg, err := network.ReadMessage(conn)
	if err != nil {
		return nil, err
	}
	if msg == nil {
		err := fmt.Errorf("expected bitfield but got %v", msg)
		return nil, err
	}
	if msg.ID != network.MsgBitfield {
		err := fmt.Errorf("expected bitfield but got ID %d", msg.ID)
		return nil, err
	}

	return msg.Payload, nil
}

// New connects with a peer, completes a handshake, and receives a handshake
// returns an err if any of those fail.
func NewClient(peer network.Peer, peerID, infoHash [20]byte) (*Client, error) {
	conn, err := net.DialTimeout("tcp", peer.String(), 3*time.Second)
	if err != nil {
		return nil, err
	}

	_, err = completeHandshake(conn, infoHash, peerID)
	if err != nil {
		conn.Close()
		return nil, err
	}

	bf, err := recvBitfield(conn)
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &Client{
		Conn:     conn,
		Choked:   true,
		Bitfield: bf,
		peer:     peer,
		infoHash: infoHash,
		peerID:   peerID,
	}, nil
}

// Read reads and consumes a message from the connection
func (c *Client) Read() (*network.Message, error) {
	msg, err := network.ReadMessage(c.Conn)
	return msg, err
}

// SendRequest sends a Request message to the peer
func (c *Client) SendRequest(index, begin, length int) error {
	req := network.FormatRequest(index, begin, length)
	_, err := c.Conn.Write(req.Serialize())
	return err
}

// SendInterested sends an Interested message to the peer
func (c *Client) SendInterested() error {
	msg := network.Message{ID: network.MsgInterested}
	_, err := c.Conn.Write(msg.Serialize())
	return err
}

// SendNotInterested sends a NotInterested message to the peer
func (c *Client) SendNotInterested() error {
	msg := network.Message{ID: network.MsgNotInterested}
	_, err := c.Conn.Write(msg.Serialize())
	return err
}

// SendUnchoke sends an Unchoke message to the peer
func (c *Client) SendUnchoke() error {
	msg := network.Message{ID: network.MsgUnchoke}
	_, err := c.Conn.Write(msg.Serialize())
	return err
}

// SendHave sends a Have message to the peer
func (c *Client) SendHave(index int) error {
	msg := network.FormatHave(index)
	_, err := c.Conn.Write(msg.Serialize())
	return err
}
