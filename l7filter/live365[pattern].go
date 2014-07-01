/*
# live365 - An Internet radio site - http://live365.com
# Pattern attributes: marginal notsofast notsofast
# Protocol groups: streaming_audio
# Wiki: http://www.protocolinfo.org/wiki/Live365
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
#
# This pattern was "contributed" (taken with permission) by the bandwidth 
# arbitrator project (www.bandwidtharbitrator.com).
#
# This pattern is unconfirmed.
# FIXME: what's going on here?
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "live365"
	mypattern.RegularExpresion = "(?i)" + "membername.*session.*player"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
