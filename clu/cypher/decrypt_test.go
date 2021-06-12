package cypher

import (
	"encoding/base64"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecode1(t *testing.T) {
	/*
	   0000   10 63 61 85 29 5a 6d bd 45 67 0d 1e 05 db 7d 45  .ca.)Zm.Eg....}E
	   0010   32 44 22 c5 54 8e 8f c2 01 4e 4d a0 13 ba 76 26  2D".T....NM...v&
	   0020   39 0d b9 11 81 f9 5e 73 a2 43 0e 44 6d 97 d1 cc  9.....^s.C.Dm...
	   0030   5d 6d 6e 53 5f 96 a5 98 e0 0a 0e 99 f0 e5 12 48  ]mnS_..........H
	*/
	tmp, _ := hex.DecodeString("10636185295a6dbd45670d1e05db7d45324422c5548e8fc2014e4da013ba7626390db91181f95e73a2430e446d97d1cc5d6d6e535f96a598e00a0e99f0e51248")

	key, _ := base64.StdEncoding.DecodeString("KY1Ajg+pDBQcP2cHnIFNRQ==")
	iv, _ := base64.StdEncoding.DecodeString("/gV+nXMOUlBbuc3uhkk/eA==")

	res, _ := Decrypt(key, iv, tmp)

	assert.Equal(t, string(res[:]), "req:192.168.1.104:00be11:DOUT_8565:execute(2, 0)\r\n")
}
