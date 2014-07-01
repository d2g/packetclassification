/*
# 100bao - a Chinese P2P protocol/program - http://www.100bao.com
# Pattern attributes: ok veryfast fast
# Protocol groups: p2p
# Wiki: http://www.protocolinfo.org/wiki/100Bao
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
#
# Pattern written by www.routerclub.com's wsgtrsys.
# The author of this pattern says it works, but this is unconfirmed. 
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "100bao"
	mypattern.RegularExpresion = "(?i)" + "^\\x01\\x01\\x05\\x0a"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
