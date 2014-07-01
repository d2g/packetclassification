/*
# Armagetron Advanced - open source Tron/snake based multiplayer game
# Pattern attributes: good slow notsofast
# Protocol groups: open_source game
# Wiki: http://protocolinfo.org/wiki/Armagetron
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
# Contributed to protocolinfo.org, possibly by joda.bot, who says "The 
# filter matches the initial transfer of configuration data. Very early 
# versions might not transfer the CYCLE_ Settings (before 0.2.5.x)."
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "armagetron"
	mypattern.RegularExpresion = "(?i)" + "YCLC_E|CYEL"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
