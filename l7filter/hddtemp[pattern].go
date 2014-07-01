/*
# hddtemp - Hard drive temperature reporting
# Pattern attributes: great veryfast fast
# Protocol groups: monitoring open_source
# Wiki: http://www.protocolinfo.org/wiki/HDDtemp
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
# 
# Usually runs on port 7634
#
# You're a silly person if you use this pattern.
#
# This pattern has been tested and is believed to work well.
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "hddtemp"
	mypattern.RegularExpresion = "(?i)" + "^\\|/dev/[a-z][a-z][a-z]\\|[0-9a-z]*\\|[0-9][0-9]\\|[cfk]\\|"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
