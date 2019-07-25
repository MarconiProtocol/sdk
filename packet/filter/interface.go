package m_packet_filter

import (
	"github.com/MarconiProtocol/gopacket"
)

/*
  Filter function plugins will need to initialize an object named 'Filter' that implements this interface

  Execute will be invoked for every packet going through the Marconi Net codec,
  and expects the result of the filter application to be written to FilterResponse
*/
type Filter interface {
	Execute(gopacket.Packet, *FilterResponse)
}