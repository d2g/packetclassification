/*
# POCO and PP365 - Chinese P2P filesharing - http://pp365.com http://poco.cn
# Pattern attributes: ok veryfast fast
# Protocol groups: p2p
# Wiki: http://www.protocolinfo.org/wiki/Poco
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
#
# The author of this pattern says it works, but this is unconfirmed.
# Written by www.routerclub.com wsgtrsys.
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "poco"
	mypattern.RegularExpresion = "(?i)" + "^\\x80\\x94\\x0a\\x01....\\x1f\\x9e"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
