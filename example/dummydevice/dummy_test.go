package dummydevice

import (
	"log"
	"testing"
)

func TestDummy(t *testing.T) {
	dev, err := dummyInit()
	if err != nil {
		log.Fatal(err)
	}
	dummyTransmit(dev)

}
