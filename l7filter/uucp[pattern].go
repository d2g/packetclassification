/*
# UUCP - Unix to Unix Copy
# Pattern attributes: ok veryfast fast
# Protocol groups: document_retrieval obsolete
# Wiki: http://www.protocolinfo.org/wiki/UUCP
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
# This is completely untested!  (I don't know how to use UUCP...)
# See http://docs.freebsd.org/info/uucp/uucp.info.The_Initial_Handshake.html
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "uucp"
	mypattern.RegularExpresion = "(?i)" + "^\\x10here="
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
