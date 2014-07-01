/*
# TeamSpeak - VoIP application - http://goteamspeak.com
# Pattern attributes: good veryfast fast
# Protocol groups: voip proprietary
# Wiki: http://www.protocolinfo.org/wiki/TeamSpeak
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
#
# This pattern has been tested by Matthew Strait and verified by packet
# traces by at least two other people.  The meaning of f4b303 is not
# known, but it seems to appear in all first packets.  This pattern only
# matches the actual UDP voice traffic, not the TeamSpeak web interface
# or "TCP query".
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "teamspeak"
	mypattern.RegularExpresion = "(?i)" + "^\\xf4\\xbe\\x03.*teamspeak"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
