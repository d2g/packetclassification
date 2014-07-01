/*
# Citrix ICA - proprietary remote desktop application - http://citrix.com
# Pattern attributes: marginal notsofast notsofast
# Protocol groups: remote_access proprietary
# Wiki: http://www.protocolinfo.org/wiki/Citrix
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
# 
# This pattern is UNTESTED.
# This is based on decode_citrix in dsniff 2.4.
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "citrix"
	mypattern.RegularExpresion = "(?i)" + "\\x32\\x26\\x85\\x92\\x58"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
