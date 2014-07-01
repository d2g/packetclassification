/*
# Subversion - a version control system
# Pattern attributes: ok veryfast fast
# Protocol groups: version_control open_source
# Wiki: http://www.protocolinfo.org/wiki/Subversion
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
#
# This pattern is UNTESTED.  (But it seems straightforward enough...)
#
# Subversion uses TCP port 3690 by default.
# This is not a valid basic GNU regular expression.
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "subversion"
	mypattern.RegularExpresion = "(?i)" + "^\\( success \\( 1 2 \\("
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
