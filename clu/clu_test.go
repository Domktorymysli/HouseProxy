package clu

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	cluIp := "192.168.0.6"
	cluPort := "1234"
	cluKey := "/qekHD76Z4PNQU57qzgxuA=="
	cluIv := "DSowHWnFXbZISF4fhIuF9A=="
	fromIp := "192.168.0.1"

	clu1 := NewClient(cluIp, cluPort, cluKey, cluIv, fromIp)

	res, err := clu1.CallFunction("dupa", "test1", "test2")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("msg: " + res.Msg)
	fmt.Println("type: " + res.Type)
	fmt.Println("sessionId: " + res.SessionId)
	fmt.Println("ip: " + res.Ip)
}
