/*
# NBNS - NetBIOS name service
# Pattern attributes: good slow notsofast
# Protocol groups: networking proprietary
# Wiki: http://www.protocolinfo.org/wiki/NBNS
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
#
# This pattern has been tested and is believed to work well.
#
# name query
# \\x01\\x10 means name query
#
# registration NB
# (\\x10 or )\\x10 means registration
#
# release NB (merged with registration)
# 0\\x10 means release
# This is not a valid basic GNU regular expression.
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "nbns"
	mypattern.RegularExpresion = "(?i)" + "\\x01\\x10\\x01|\\)\\x10\\x01\\x01|0\\x10\\x01"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
