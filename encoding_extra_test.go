package encoding_extra

import (
	"testing"
)

func TestEncodingBitcoinType(t *testing.T) {
	num := 123456789
	expectedEncodedStr := "BukQL"

	bitcoinEncoder := BaseExtraEncoding{
		EncodingCharacters: Base58EncodingTypeBitcoin,
	}
	encodedStr := bitcoinEncoder.Encode(num)

	if encodedStr != expectedEncodedStr {
		t.Fatalf("Bitcoin Type Encoder - got %s, want %s", encodedStr, expectedEncodedStr)
	}
}

func TestDecodingBitcoinType(t *testing.T) {
	encodedStr := "BukQL"
	expectedDecodedInt := 123456789

	bitcoinDecoder := BaseExtraEncoding{
		EncodingCharacters: Base58EncodingTypeBitcoin,
	}
	decodedInt := bitcoinDecoder.Decode(encodedStr)

	if decodedInt != expectedDecodedInt {
		t.Fatalf("Bitcoin Type Encoder - got %d, want %d", decodedInt, expectedDecodedInt)
	}
}
