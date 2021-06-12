package cypher

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncode(t *testing.T) {
	key, _ := base64.StdEncoding.DecodeString("KY1Ajg+pDBQcP2cHnIFNRQ==")
	iv, _ := base64.StdEncoding.DecodeString("/gV+nXMOUlBbuc3uhkk/eA==")

	msg := []byte("req:192.168.1.104:00be11:DOUT_8565:execute(2, 0)\r\n")

	res, _ := Encrypt(key, iv, msg)
	txt := hex.EncodeToString(res)
	fmt.Println(txt)

	assert.Equal(t, txt, "10636185295a6dbd45670d1e05db7d45324422c5548e8fc2014e4da013ba7626390db91181f95e73a2430e446d97d1cc5d6d6e535f96a598e00a0e99f0e51248")
}
