/*
# Doom 3 - computer game
# Pattern attributes: good veryfast fast
# Protocol groups: game proprietary
# Wiki: http://www.protocolinfo.org/wiki/Doom
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
#
# Thanks to Clayton Macleod (cherrytwist at gmail.com).
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "doom3"
	mypattern.RegularExpresion = "(?i)" + "^\\xff\\xffchallenge"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
