/*
# Half-Life 2 Deathmatch - popular computer game
# Pattern attributes: good veryfast fast
# Protocol groups: game proprietary
# Wiki: http://www.protocolinfo.org/wiki/Half-Life
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
# 
# By Clayton Macleod <cherrytwist TA gmail.com>
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "halflife2-deathmatch"
	mypattern.RegularExpresion = "(?i)" + "^\\xff\\xff\\xff\\xff.*hl2mpDeathmatch"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
