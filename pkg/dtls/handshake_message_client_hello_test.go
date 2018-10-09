package dtls

import (
	"reflect"
	"testing"
	"time"
)

func TestHandshakeMessageClientHello(t *testing.T) {
	rawClientHello := []byte{0x01, 0x00, 0x00, 0x7b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7b, 0xfe, 0xff, 0xb6, 0x2f,
		0xce, 0x5c, 0x42, 0x54, 0xff, 0x86, 0xe1, 0x24, 0x41, 0x91, 0x42, 0x62, 0x15, 0xad, 0x16, 0xc9,
		0x15, 0x8d, 0x95, 0x71, 0x8a, 0xbb, 0x22, 0xd7, 0x47, 0xec, 0xd8, 0x3d, 0xdc, 0x4b, 0x00, 0x00,
		0x00, 0x04, 0xc0, 0x09, 0xc0, 0x23, 0x01, 0x00, 0x00, 0x2f, 0x00, 0x0b, 0x00, 0x04, 0x03, 0x00, 0x01, 0x02,
		0x00, 0x0a, 0x00, 0x0c, 0x00, 0x0a, 0x00, 0x1d, 0x00, 0x17, 0x00, 0x1e, 0x00, 0x19, 0x00, 0x18,
		0x00, 0x23, 0x00, 0x00, 0x00, 0x0e, 0x00, 0x07, 0x00, 0x04, 0x00, 0x02, 0x00, 0x01, 0x00, 0x00,
		0x16, 0x00, 0x00, 0x00, 0x17, 0x00, 0x00}
	parsedClientHello := &clientHello{
		messageSequence: 0,
		fragmentOffset:  0,
		fragmentLength:  123,
		version:         protocolVersion{0xFE, 0xFF},
		random: handshakeRandom{
			time.Unix(3056539021, 0),
			[28]byte{0x42, 0x54, 0xff, 0x86, 0xe1, 0x24, 0x41, 0x91, 0x42, 0x62, 0x15, 0xad, 0x16, 0xc9, 0x15, 0x8d, 0x95, 0x71, 0x8a, 0xbb, 0x22, 0xd7, 0x47, 0xec, 0xd8, 0x3d, 0xdc, 0x4b},
		},
		cipherSuites: []*cipherSuite{
			cipherSuites[TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA],
			cipherSuites[TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256],
		},
	}

	var c clientHello
	if err := c.unmarshal(rawClientHello); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(c, parsedClientHello) {
		t.Errorf("clientHello unmarshal: got %#v, want %#v", c, parsedClientHello)
	}

	raw, err := c.marshal()
	if err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(raw, rawClientHello) {
		t.Errorf("clientHello marshal: got %#v, want %#v", raw, rawClientHello)
	}
}
