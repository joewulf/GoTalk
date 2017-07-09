package pdubase

type PduHeader struct {
	length    uint32
	version   uint16
	flag      uint16
	serviceId uint16
	commandId uint16
	seqNum    uint16
	reversed  uint16
}

func (p PduHeader) Version() uint16 {
	return p.version
}

func (p PduHeader) Flag() uint16 {
	return p.flag
}

func (p PduHeader) ServiceId() uint16 {
	return p.serviceId
}

func (p PduHeader) CommandId() uint16 {
	return p.commandId
}

func (p PduHeader) SeqNum() uint16 {
	return p.seqNum
}

func (p PduHeader) Reversed() uint16 {
	return p.reversed
}

func (p *PduHeader) SetVersion(version uint16) {
	p.version = version
}

func (p *PduHeader) SetFlag(flag uint16) {
	p.flag = flag
}

func (p *PduHeader) SetServiceId(serviceId uint16) {
	p.serviceId = serviceId
}

func (p *PduHeader) SetCommandId(commandId uint16){
	p.commandId = commandId
}

func (p *PduHeader) SetSeqNum(seqNum uint16) {
	p.seqNum = seqNum
}

func (p PduHeader) SetReversed(reversed uint16)  {
	p.reversed = reversed
}
