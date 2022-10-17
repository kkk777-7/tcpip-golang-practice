package dummydevice

import (
	"github.com/kkk777-7/tcpip-golang-practice/pkg/net"
)

type DummyDevice struct {
	name string
	mtu  int
}

func (dd *DummyDevice) Type() net.HardwareType {
	return net.HardwareTypeDummy
}

func (dd *DummyDevice) Name() string {
	return dd.name
}

func (dd *DummyDevice) MTU() int {
	return dd.mtu
}

func dummyInit() (*net.Device, error) {
	ddev := &DummyDevice{"DUMMY!!", 0}
	dev, err := net.RegisterDevice(ddev)
	if err != nil {
		return nil, err
	}
	return dev, nil
}

func dummyTransmit(dd *net.Device) {}
