/*
# Day of Defeat: Source - game (Half-Life 2 mod) - http://www.valvesoftware.com
# Pattern attributes: good veryfast fast
# Protocol groups: game proprietary
# Wiki: http://www.protocolinfo.org/wiki/Day_of_Defeat:Source
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
#
# By Clayton Macleod <cherry twist at gmail dot com>
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "dayofdefeat-source"
	mypattern.RegularExpresion = "(?i)" + "^\\xff\\xff\\xff\\xff.*dodDay of Defeat"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
