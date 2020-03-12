package dtls

import (
	"crypto/tls"
	"reflect"
	"testing"

	"golang.org/x/xerrors"
)

func TestParseSignatureSchemes(t *testing.T) {
	cases := map[string]struct {
		input    []tls.SignatureScheme
		expected []signatureHashAlgorithm
		err      error
	}{
		"Translate": {
			input: []tls.SignatureScheme{
				tls.ECDSAWithP256AndSHA256,
				tls.ECDSAWithP384AndSHA384,
				tls.ECDSAWithP521AndSHA512,
				tls.PKCS1WithSHA256,
				tls.PKCS1WithSHA384,
				tls.PKCS1WithSHA512,
			},
			expected: []signatureHashAlgorithm{
				{hashAlgorithmSHA256, signatureAlgorithmECDSA},
				{hashAlgorithmSHA384, signatureAlgorithmECDSA},
				{hashAlgorithmSHA512, signatureAlgorithmECDSA},
				{hashAlgorithmSHA256, signatureAlgorithmRSA},
				{hashAlgorithmSHA384, signatureAlgorithmRSA},
				{hashAlgorithmSHA512, signatureAlgorithmRSA},
			},
			err: nil,
		},
		"InvalidSignatureAlgorithm": {
			input: []tls.SignatureScheme{
				tls.ECDSAWithP256AndSHA256, // Valid
				0x04FF,                     // Invalid: unknown signature with SHA-256
			},
			expected: nil,
			err:      errInvalidSignatureAlgorithm,
		},
		"InvalidHashAlgorithm": {
			input: []tls.SignatureScheme{
				tls.ECDSAWithP256AndSHA256, // Valid
				0x0003,                     // Invalid: ECDSA with MD2
			},
			expected: nil,
			err:      errInvalidHashAlgorithm,
		},
	}

	for name, testCase := range cases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			output, err := parseSignatureSchemes(testCase.input)
			if testCase.err != nil && !xerrors.Is(err, testCase.err) {
				t.Fatalf("Expected error: %v, got: %v", testCase.err, err)
			}
			if !reflect.DeepEqual(testCase.expected, output) {
				t.Errorf("Expected signatureHashAlgorithm:\n%+v\ngot:\n%+v", testCase.expected, output)
			}
		})
	}
}
