package dispatcher

import (
	"net"
	"fmt"
	"os"

	_ "github.com/GoTalk/pkg/pduhandler"
)

type Dispatcher struct {

}


func (d *Dispatcher) Start(l net.Listener) {
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}





