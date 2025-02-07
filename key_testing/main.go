package main

import (
	"fmt"

	"github.com/tillitis/tkeyclient"
	"github.com/tillitis/tkeysign"
	"github.com/tillitis/tkeyutil"
)

func main() {
	tk := tkeyclient.New()
	err := tk.Connect("/dev/ttyACM2")
	if err != nil {
		panic("Couldn't connect to TKey!")
	}

	EnterUSS()

}

func EnterUSS() {
	secret, err := tkeyutil.InputUSS()
	if err != nil {
		panic("Couldn't get USS!")
	}

	fmt.Println("USS phrase entered successfully:", string(secret))
}

func DisplayPubKey(tk *tkeyclient.TillitisKey) {
	signer := tkeysign.New(tk)
	pubkey, err2 := signer.GetPubkey()
	if err2 != nil {
		panic("Couldn't get Public Key!")
	}
	fmt.Printf("Public Key is: %s", pubkey)
}
