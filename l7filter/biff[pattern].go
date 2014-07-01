/*
# Biff - new mail notification 
# Pattern attributes: good fast fast undermatch overmatch
# Protocol groups: mail
# Wiki: http://www.protocolinfo.org/wiki/Biff
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
#
# Usually runs on port 512
#
# This pattern is completely untested.
# This is a rare case where we will specify a $ (end of line), since
# this is the entirety of the communication.
# something that looks like a username, an @, a number.
# won't catch usernames that have strange characters in them.
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "biff"
	mypattern.RegularExpresion = "(?i)" + "^[a-z][a-z0-9]+@[1-9][0-9]+$"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
