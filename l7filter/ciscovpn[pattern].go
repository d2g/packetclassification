/*
# Cisco VPN - VPN client software to a Cisco VPN server
# Pattern attributes: ok veryfast fast
# Protocol groups: remote_access proprietary
# Wiki: http://www.protocolinfo.org/wiki/Cisco_VPN
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
# 
# This pattern contributed by Myles Uyema <myles AT uyema.net>
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "ciscovpn"
	mypattern.RegularExpresion = "(?i)" + "^\\x01\\xf4\\x01\\xf4"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
