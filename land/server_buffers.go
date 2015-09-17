package land

import (
	"github.com/Nyarum/noterius/core"
	"github.com/Nyarum/noterius/pills"

	"bytes"
	"encoding/binary"
	"io"
	"log"
	"net"
)

// Buffers struct for read and write channels
type Buffers struct {
	WriteChannel chan string
	ReadChannel  chan string
}

// NewBuffers method for init Buffers struct
func NewBuffers() *Buffers {
	return &Buffers{
		WriteChannel: make(chan string),
		ReadChannel:  make(chan string),
	}
}

// GetWriteChannel method for get WriteChannel from Buffers struct
func (b *Buffers) GetWriteChannel() chan string {
	return b.WriteChannel
}

// GetReadChannel method for get ReadChannel from Buffers struct
func (b *Buffers) GetReadChannel() chan string {
	return b.ReadChannel
}

// WriteHandler method for write bytes to socket in loop from channel
func (b *Buffers) WriteHandler(c net.Conn) {
	// Write one packet for client with time.Now()
	pill := pills.NewPill()
	c.Write(pill.Encrypt(pill.SetOpcode(940).GetOutcomingCrumb()))

	for v := range b.WriteChannel {
		c.Write([]byte(v))
	}
}

// ReadHandler method for read bytes from socket in loop to channel
func (b *Buffers) ReadHandler(c net.Conn, conf core.Config) {
	var (
		bytesAlloc []byte = make([]byte, conf.Option.LenBuffer)
	)

	buf := bytes.NewBuffer(bytesAlloc)
	for {
		_, err := c.Read(bytesAlloc)
		if err == io.EOF {
			log.Printf("Client [%v] is disconnected\n", c.RemoteAddr())
			return
		} else if err != nil {
			if err.(net.Error).Timeout() {
				log.Printf("Client [%v] is timeout\n", c.RemoteAddr())
				return
			}

			log.Printf("Client [%v] is error read packet, err - %v\n", c.RemoteAddr(), err)
		}

		var lastGotLen int
		readLen := func() bool {
			lastGotLen = int(binary.BigEndian.Uint16(buf.Bytes()[0:2]))
			if lastGotLen == 0 {
				return false
			} else if buf.Len() < lastGotLen {
				return false
			}

			return true
		}

		for readLen() {
			b.ReadChannel <- string(buf.Next(lastGotLen))
		}
	}
}
