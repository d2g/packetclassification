package l7filter

import (
	"code.google.com/p/go.net/ipv4"
	"github.com/d2g/packetclassification"
	"github.com/d2g/tcp"
	"regexp"
)

/*
	This Adds support for the L7-Filter Patterns
*/
type Pattern struct {
	Name                     string
	RegularExpresion         string
	compiledRegularExpresion *regexp.Regexp
}

func (t Pattern) Classify(packet []byte) (bool, packetclassification.Classification, error) {

	ipheader, err := ipv4.ParseHeader(packet)
	if err != nil {
		return false, packetclassification.Classification{}, err
	}

	tcpheader, err := tcp.ParseHeader(packet[ipheader.Len:])
	if err != nil {
		return false, packetclassification.Classification{}, err
	}

	packet = packet[(ipheader.Len + (int(tcpheader.DataOffset) * 4)):]

	if t.compiledRegularExpresion == nil {
		t.compiledRegularExpresion = regexp.MustCompile(t.RegularExpresion)
	}

	if t.compiledRegularExpresion.Find(packet) != nil {
		classification := packetclassification.Classification{}
		classification.Protocol = t.Name
		return true, classification, nil
	}

	return false, packetclassification.Classification{}, nil
}
