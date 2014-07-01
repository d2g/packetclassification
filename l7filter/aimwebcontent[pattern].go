/*
# AIM web content - ads/news content downloaded by AOL Instant Messenger
# Pattern attributes: good notsofast notsofast
# Protocol groups: chat document_retrieval proprietary
# Wiki: http://www.protocolinfo.org/wiki/AIM
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
#
# This pattern has been tested and is believed to work well.
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "aimwebcontent"
	mypattern.RegularExpresion = "(?i)" + "user-agent:aim/"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
