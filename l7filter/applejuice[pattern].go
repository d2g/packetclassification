/*
# Apple Juice - P2P filesharing - http://www.applejuicenet.de
# Pattern attributes: great veryfast fast
# Protocol groups: p2p
# Wiki: http://www.protocolinfo.org/wiki/AppleJuice
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
#
# This pattern has been tested with the Linux version (version
# 0,29,142,229).  It matches search reqests and file transfers.
# this pattern extracted from ipp2p, by Eicke Friedrich.
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "applejuice"
	mypattern.RegularExpresion = "(?i)" + "^ajprot\\x0d\\x0a"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
