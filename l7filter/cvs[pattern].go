/*
# CVS - Concurrent Versions System
# Pattern attributes: good veryfast fast
# Protocol groups: version_control open_source
# Wiki: http://www.protocolinfo.org/wiki/CVS
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
# Matches pserver login.  AUTH is for actually starting the protocol
# VERIFICATION is for authenticating without starting the protocols
# and GSSAPI is for using security services such as kerberos.
# http://www.loria.fr/~molli/cvs/doc/cvsclient_3.html
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "cvs"
	mypattern.RegularExpresion = "(?i)" + "^BEGIN (AUTH|VERIFICATION|GSSAPI) REQUEST\\x0a"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
