package pduhandler

type PduHandler interface {
	HandleReq()
	HandleResp()
}
