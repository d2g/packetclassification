/*
# pcAnywhere - Symantec remote access program
# Pattern attributes: marginal veryfast fast
# Protocol groups: remote_access proprietary
# Wiki: http://www.protocolinfo.org/wiki/PcAnywhere
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
# This is completely untested!
# See http://www.unixwiz.net/tools/pcascan.txt
# I think this only matches queries and not the bulk of the traffic!
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "pcanywhere"
	mypattern.RegularExpresion = "(?i)" + "^(nq|st)$"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
