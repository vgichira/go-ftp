package ftp

import (
	"net"
)

// Conn represents a connection to the FTP Server
type Conn struct {
	conn     net.Conn
	dataType dataType
	dataPort *dataPort
	rootDir  string
	workDir  string
}

// NewConn returns a new FTP Connection
func NewConn(conn net.Conn, rootDir string) *Conn {
	return &Conn{
		conn:    conn,
		rootDir: rootDir,
		workDir: "/",
	}
}
