/*
# Whois - query/response system, usually used for domain name info - RFC 3912
# Pattern attributes: good notsofast notsofast overmatch
# Protocol groups: networking ietf_draft_standard
# Wiki: http://www.protocolinfo.org/wiki/Whois
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
#
# Usually runs on TCP port 43
# 
# This pattern has been tested and is believed to work well.
# Matches the query.  Assumes only that it is printable ASCII without wierd
# whitespace.
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "whois"
	mypattern.RegularExpresion = "(?i)" + "^[ !-~]+\\x0d\\x0a$"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
