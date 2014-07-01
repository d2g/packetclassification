/*
# IP printing - a new standard for UNIX printing - RFC 2911
# Pattern attributes: good notsofast notsofast
# Protocol groups: printer ietf_proposed_standard
# Wiki: http://www.protocolinfo.org/wiki/IPP
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
#
# This pattern has been tested and is believed to work well.
# It's unlikely that anything else has this string, but I think we could
# do a bit better...
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "ipp"
	mypattern.RegularExpresion = "(?i)" + "ipp://"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
