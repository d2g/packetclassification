/*
# RDP - Remote Desktop Protocol (used in Windows Terminal Services)
# Pattern attributes: ok notsofast notsofast
# Protocol groups: remote_access proprietary
# Wiki: http://www.protocolinfo.org/wiki/RDP
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
#
# This pattern was submitted by Michael Leong.  It has been tested under the 
# following conditions: "WinXP Pro with all the patches, rdesktop server 
# running on port 7000 instead of 3389 --> WinXP Pro Remote Desktop Client."
# Also tested is WinXP to Win 2000 Server.
# At least one other person has reported it to work as well.
# Old pattern, submitted by Daniel Weatherford.
# rdpdr.*cliprdp.*rdpsnd 
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "rdp"
	mypattern.RegularExpresion = "(?i)" + "rdpdr.*cliprdr.*rdpsnd"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
