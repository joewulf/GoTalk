package dispatcher

import (
	"net"
	"fmt"
	"os"

	"github.com/GoTalk/pkg/pduhandler"
)

type Dispatcher struct {
	recvConn []net.Conn
	sendConn []net.Conn
	recvSeq int
	sendSeq int
}

const READ_BUF_SIZE = 2048

func (d *Dispatcher) Start(l net.Listener) {
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		d.recvConn = append(d.recvConn, conn)
		// Handle connections in a new goroutine.
		go d.handleRequest()
	}
}

// Handles incoming requests.
func (d *Dispatcher) handleRequest() {
	// Make a buffer to hold incoming data.
	recvBuf := make([]byte, READ_BUF_SIZE)

	var p pduhandler.PduHandler
	for {
		// Read the incoming connection into the buffer.
		n, err := d.recvConn[d.recvSeq].Read(recvBuf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}

		if n == READ_BUF_SIZE {
			go p.HandlePdu(recvBuf, d)
		}
	}

}

func (d *Dispatcher) handleResponse() {
	sendBuf := make([]byte, READ_BUF_SIZE)

	for {
		n, err := d.sendConn[d.sendSeq].Write(sendBuf)
		if err != nil {
			fmt.Println("Error writing:", err.Error())
		}

		if n == READ_BUF_SIZE {
			break
		}
	}

}

