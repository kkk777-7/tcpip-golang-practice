package net

import (
	"fmt"
	"sync"
)

const NET_DEVICE_FLAG_UP uint16 = 0x0001

type HardwareType uint16

const (
	HardwareTypeLoopback = 0x0000
	HardwareTypeEthernet = 0x0001
	HardwareTypeDummy    = 0x9999
)

type LinkDevice interface {
	Type() HardwareType
	Name() string
	// Address() HardwareAddress
	// BroadcastAddress() HardwareAddress
	MTU() int
	// NeedARP() bool
	// Close()
	// Read(data []byte) (int, error)
	// RxHandler(frame []byte, callback LinkDeviceCallbackHandler)
	// Tx(proto EthernetType, data []byte, dst []byte) error
}

type Device struct {
	LinkDevice
	running bool
	mRun    sync.Mutex
}

var devices = map[LinkDevice]*Device{}

func RegisterDevice(link LinkDevice) (*Device, error) {
	if _, ok := devices[link]; ok {
		return nil, fmt.Errorf("link device '%s' is already registerd\n", link.Name())
	}
	dev := &Device{
		LinkDevice: link,
	}
	devices[link] = dev
	return dev, nil
}

func OpenDevice(d *Device) {
	d.mRun.Lock()
	if d.running {
		fmt.Printf("already opened, dev=%s\n", d.Name())
	}
	d.running = true
	d.mRun.Unlock()
}

func CloseDevice(d *Device) {
	d.mRun.Lock()
	if !d.running {
		fmt.Printf("already closed, dev=%s\n", d.Name())
	}
	d.running = false
	d.mRun.Unlock()
}

func InputDevice(d *Device, data uint8, len int) {
	fmt.Printf("dev=%s, type=%16b, len=%d\n", d.Name(), d.Type(), len)
}

func OutputDevice(d *Device, data uint8, len int) error {
	if !d.running {
		return fmt.Errorf("not opened, dev=%s", d.Name())
	}
	if len > d.MTU() {
		return fmt.Errorf("too long, dev=%s, mtu=%d, len=%d\n", d.Name(), d.MTU(), len)
	}
	fmt.Printf("dev=%s, type=%16b, len=%d\n", d.Name(), d.Type(), len)

	// TODO

	return nil
}
