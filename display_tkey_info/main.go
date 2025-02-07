package main

import (
	"fmt"

	"github.com/tillitis/tkeyclient"
)

func main() {
	tk := tkeyclient.New()

	ConnectTKey(tk, GetSerialPort(tk))
	DisplayNameVersion(tk)
	DisplayUDI(tk)

}

func DisplayNameVersion(tk *tkeyclient.TillitisKey) {
	nameVer, err := tk.GetNameVersion()
	if err != nil {
		panic("Could not get Name Version")
	}
	fmt.Printf("Firmware name0:'%s' name1:'%s' version:%d\n",
		nameVer.Name0, nameVer.Name1, nameVer.Version)

}

func GetSerialPort(tk *tkeyclient.TillitisKey) string {
	port, err := tkeyclient.DetectSerialPort(false)
	fmt.Printf("Serial Port is: '%s'", port)

	if err != nil {
		panic("Could not detect serial port")
	}
	return port
}

func ConnectTKey(tk *tkeyclient.TillitisKey, port string) {
	err := tk.Connect(port)
	if err != nil {
		panic("Could not connect to TillitisKey")
	}
}

func DisplayUDI(tk *tkeyclient.TillitisKey) {
	udi, err := tk.GetUDI()
	if err != nil {
		panic("Could not get UDI")
	}
	fmt.Printf("Unique Device ID is: '%s'", udi)
}
