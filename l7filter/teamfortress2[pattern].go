/*
# Team Fortress 2 - network game - http://www.valvesoftware.com
# Pattern attributes: good veryfast fast
# Protocol groups: game proprietary
# Wiki: http://www.protocolinfo.org/wiki/Team_Fortress
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
#
# Credits: Clayton Macleod <cherry twist at gmail dot com>
#          Jan Engelhardt <jengelh at computergmbh dot de>
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "teamfortress2"
	mypattern.RegularExpresion = "(?i)" + "^\\xff\\xff\\xff\\xff.....*tfTeam Fortress"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
