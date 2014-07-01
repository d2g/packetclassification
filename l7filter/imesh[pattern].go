/*
# iMesh - the native protocol of iMesh, a P2P application - http://imesh.com
# Pattern attributes: ok fast notsofast
# Protocol groups: p2p
# Wiki:   http://protocolinfo.org/wiki/iMesh
# Copyright (C) 2008 Matthew Strait, Ethan Sommer; See ../LICENSE
#
# depending on the version of iMesh (the program), it can also use fasttrack,
# gnutella and edonkey in addition to iMesh (the protocol).
# The first branch matches the login
# The second branch matches the main non-download connection (searches, etc)
# The third branch matches downloads of "premium" content
# The fourth branch matches peer downloads.
*/
package l7filter

import (
	"github.com/d2g/packetclassification"
)

func init() {
	mypattern := Pattern{}
	mypattern.Name = "imesh"
	mypattern.RegularExpresion = "(?i)" + "^(post[\\x09-\\x0d -~]*<PasswordHash>................................</PasswordHash><ClientVer>|\\x34\\x80?\\x0d?\\xfc\\xff\\x04|get[\\x09-\\x0d -~]*Host: imsh\\.download-prod\\.musicnet\\.com|\\x02[\\x01\\x02]\\x83.*\\x02[\\x01\\x02]\\x83)"
	packetclassification.RegisterClassifier(packetclassification.Classifier(mypattern))
}
