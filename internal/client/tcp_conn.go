// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package client

import (
	"errors"
	"net"
	"time"

	"github.com/pion/transport/v2"
	"github.com/pion/turn/v2/internal/proto"
)

var (
	errInvalidTURNFrame    = errors.New("data is not a valid TURN frame, no STUN or ChannelData found")
	errIncompleteTURNFrame = errors.New("data contains incomplete STUN or TURN frame")
)

const (
	stunHeaderSize = 20
)

var _ transport.TCPConn = (*TCPConn)(nil) // Includes type check for net.Conn

// TCPConn wraps a net.TCPConn and returns the allocations relayed
// transport address in response to TCPConn.LocalAddress()
type TCPConn struct {
	*net.TCPConn
	remoteAddress  *net.TCPAddr
	allocation     *TCPAllocation
	acceptDeadline time.Duration
	ConnectionID   proto.ConnectionID
}

type ConnectionAttempt struct {
	from *net.TCPAddr
	cid  proto.ConnectionID
}

func (c *TCPConn) LocalAddress() net.Addr {
	return c.allocation.Addr()
}

func (c *TCPConn) RemoteAddress() net.Addr {
	return c.remoteAddress
}