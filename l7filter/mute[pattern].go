/*
# MUTE - P2P filesharing - http://mute-net.sourceforge.net
# Pattern attributes: marginal fast fast
# Protocol groups: p2p open_source
# Wiki: http://www.protocolinfo.org/wiki/MUTE
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
#
# This pattern is lightly tested.  I don't know for sure that it will 
# match the actual file transfers.
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "mute"
	mypattern.RegularExpresion = "(?i)" + "^(Public|AES)Key: [0-9a-f]*\\x0aEnd(Public|AES)Key\\x0a$"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
