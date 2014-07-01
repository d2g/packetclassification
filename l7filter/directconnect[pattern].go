/*
# Direct Connect - P2P filesharing - http://www.neo-modus.com
# Pattern attributes: good fast fast
# Protocol groups: p2p
# Wiki: http://www.protocolinfo.org/wiki/Direct_Connect
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
#
# Direct Connect "hubs" listen on port 411
# http://www.dcpp.net/wiki/
# I've verified that this pattern can be used to limit direct connect
# bandwidth using DC:PRO 0.2.3.149R11.
# client-to-client handshake|client-to-hub login, hub speaking|client-to-hub login, client speaking
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "directconnect"
	mypattern.RegularExpresion = "(?i)" + "^(\\$mynick |\\$lock |\\$key )"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
