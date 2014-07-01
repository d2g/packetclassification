/*
# Tonghuashun - stock analysis and trading; Chinese - http://www.10jqka.com.cn
# Pattern attributes: ok fast fast
# Protocol groups: 
# Wiki: http://www.protocolinfo.org/wiki/Tonghuashun
# Copyright (C) 2009 Matthew Strait; See ../LICENSE
# Pattern contributed by liangjun without comment.
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "tonghuashun"
	mypattern.RegularExpresion = "(?i)" + "^(GET /docookie\\.php\\?uname=|\\xfd\\xfd\\xfd\\xfd\\x30\\x30\\x30\\x30\\x30)"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
