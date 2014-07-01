/*
# Freenet - Anonymous information retrieval - http://freenetproject.org
# Pattern attributes: poor veryfast fast
# Protocol groups: p2p document_retrieval open_source
# Wiki: http://www.protocolinfo.org/wiki/Freenet
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
# Freenet is intentionally hard to identify...
# This is empirical, only tested on one computer, and unlikely to work anymore.
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "freenet"
	mypattern.RegularExpresion = "(?i)" + "^\\x01[\\x08\\x09][\\x03\\x04]"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
