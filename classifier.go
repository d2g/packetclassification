package packetclassification

import ()

type Classifier interface {
	// Each Type Returns
	// bool 			- I was/wasn't able to classify this packet.
	// Classification 	- The Classification of the packet or nil.
	// error			- Any Errors
	Classify([]byte) (bool, Classification, error)
}
