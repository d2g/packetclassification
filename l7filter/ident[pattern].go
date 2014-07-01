/*
# Ident - Identification Protocol - RFC 1413
# Pattern attributes: good fast fast
# Protocol groups: networking ietf_proposed_standard
# Wiki: http://www.protocolinfo.org/wiki/Ident
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
#
# Usually runs on port 113
#
# This pattern is believed to work.
# "number , numberCRLF" possibly without the CR and/or LF.
# ^$ is appropriate because the first packet should never have anything
# else in it.
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "ident"
	mypattern.RegularExpresion = "(?i)" + "^[1-9][0-9]?[0-9]?[0-9]?[0-9]?[\\x09-\\x0d]*,[\\x09-\\x0d]*[1-9][0-9]?[0-9]?[0-9]?[0-9]?(\\x0d\\x0a|[\\x0d\\x0a])?$"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
