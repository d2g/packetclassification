/*
# Dazhihui - stock analysis and trading; Chinese - http://www.gw.com.cn
# Pattern attributes: fast fast ok
# Protocol groups: 
# Wiki: http://www.protocolinfo.org/wiki/Dazhihui
# Copyright (C) 2009 Matthew Strait; See ../LICENSE
# Pattern contributed by liangjun without comment.
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "dazhihui"
	mypattern.RegularExpresion = "(?i)" + "^(longaccoun|qsver2auth|\\x35[57]\\x30|\\+\\x10\\*)"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
