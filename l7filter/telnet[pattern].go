/*
# Telnet - Insecure remote login - RFC 854
# Pattern attributes: good veryfast fast
# Protocol groups: remote_access obsolete ietf_internet_standard
# Wiki: http://www.protocolinfo.org/wiki/Telnet
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
#
# Usually runs on port 23
#
# This pattern is lightly tested.
# Matches at least three IAC (Do|Will|Don't|Won't) commands in a row.  
# My telnet client sends 9 when I connect, so this should be fine.
# This pattern could fail on a unchatty connection or it could be 
# matched by something non-telnet spewing a lot of stuff in the fb-ff range.
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "telnet"
	mypattern.RegularExpresion = "(?i)" + "^\\xff[\\xfb-\\xfe].\\xff[\\xfb-\\xfe].\\xff[\\xfb-\\xfe]"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
