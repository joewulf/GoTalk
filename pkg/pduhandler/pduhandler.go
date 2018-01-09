package pduhandler


import "github.com/GoTalk/pkg/dispatcher"


type PduHandler interface {
	HandlePdu(buf []byte, d *dispatcher.Dispatcher)
}
