/*
# Battlefield 1942 - An EA game
# Pattern attributes: ok veryfast fast
# Protocol groups: game proprietary
# Wiki: http://www.protocolinfo.org/wiki/Battlefield_1942
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
#
# Contributed by Myles Uyema <mylesuyema AT gmail.com>
#
# This pattern has only been tested by one person.
# tested on two original EA battlefield 1942 servers
# matches the first two packets of joining a server
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "battlefield1942"
	mypattern.RegularExpresion = "(?i)" + "^\\x01\\x11\\x10\\|\\xf8\\x02\\x10\\x40\\x06"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
