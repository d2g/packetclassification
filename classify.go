package packetclassification

import (
//"log"
)

/*
	Classifiers Don't have a priority. Although this means that we have to
	optimise the classifiers by the ammount they are used.
	TODO: Optimise.
*/
var packetclassifiers []Classifier

func RegisterClassifier(registratant Classifier) {
	packetclassifiers = append(packetclassifiers, registratant)
}

func Classify(packet []byte) (classified bool, classification Classification, err error) {

	for i := 0; i < len(packetclassifiers); i++ {
		classified, classification, err = packetclassifiers[i].Classify(packet)
		if classified {
			return
		}
	}

	return
}
