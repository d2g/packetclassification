/*
# Gkrellm - a system monitor - http://gkrellm.net
# Pattern attributes: great veryfast fast
# Protocol groups: monitoring open_source
# Wiki: http://www.protocolinfo.org/wiki/Gkrellm
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
# 
# This pattern has been tested and is believed to work well.
# Since this is not anything resembling a published protocol, it may change without
# warning in new versions of gkrellm.
# tested with gkrellm 2.2.7
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "gkrellm"
	mypattern.RegularExpresion = "(?i)" + "^gkrellm [23].[0-9].[0-9]\\x0a$"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
