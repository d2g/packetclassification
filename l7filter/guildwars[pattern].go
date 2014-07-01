/*
# Guild Wars - online game - http://guildwars.com
# Pattern attributes: marginal veryfast fast
# Protocol groups: game proprietary
# Wiki: http://www.protocolinfo.org/wiki/Guild_Wars
# Copyright (C) 2008 Matthew Strait; See ../LICENSE
# Contributed on protocolinfo by Greatwolf with the comment, "Guild Wars 
# uses encrypted data on tcp/6112 and may be impossible to match by 
# content. An experimental filter has been written to match Guild Wars 
# packets. More testing is still required to determine the effectiveness 
# of this pattern."
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "guildwars"
	mypattern.RegularExpresion = "(?i)" + "^[\\x04\\x05]\\x0c.i\\x01"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
