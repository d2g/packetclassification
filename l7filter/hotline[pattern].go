/*
# Hotline - An old P2P filesharing protocol
# Pattern attributes: marginal fast fast
# Protocol groups: p2p
# Wiki: http://www.protocolinfo.org/wiki/Hotline
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
# 
# This pattern is untested!
#
# This is lifted from http://oofle.com/filesharing.php?app=hotline
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "hotline"
	mypattern.RegularExpresion = "(?i)" + "^....................TRTPHOTL\\x01\\x02"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
