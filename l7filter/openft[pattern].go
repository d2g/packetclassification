/*
# OpenFT - P2P filesharing (implemented in giFT library)
# Pattern attributes: good notsofast notsofast
# Protocol groups: p2p open_source
# Wiki: http://www.protocolinfo.org/wiki/OpenFT
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
# Ben Efros <ben AT xgendev.com> says:
# "This pattern identifies openFT P2P transfers fine.  openFT is part of giFT
# and is a pretty large p2p network. I would describe this pattern as pretty 
# weak, but it works for the giFT-based clients I've used."
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "openft"
	mypattern.RegularExpresion = "(?i)" + "x-openftalias: [-)(0-9a-z ~.]"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
