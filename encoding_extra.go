package encoding_extra

import (
	"math"
	"slices"
	"strings"
)

const (
	Base58EncodingTypeBitcoin = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	Base58EncodingTypeFlickr  = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
)

type BaseExtraEncoding struct {
	EncodingCharacters string
}

func (b *BaseExtraEncoding) GetEncodingBaseLength() int {
	return len(b.EncodingCharacters)
}

func (b *BaseExtraEncoding) Encode(integer int) string {
	var encodedStr string

	for integer > 0 {
		remainder := integer % b.GetEncodingBaseLength()
		encodedStr = strings.Split(b.EncodingCharacters, "")[remainder] + encodedStr

		integer /= b.GetEncodingBaseLength()
	}

	return encodedStr
}

func (b *BaseExtraEncoding) Decode(encodedStr string) int {
	var decodedInt int
	var indexes []int

	for _, j := range encodedStr {
		indexes = append(indexes, strings.Index(b.EncodingCharacters, string(j)))
	}

	slices.Reverse(indexes)

	for i, j := range indexes {
		decodedInt += j * int(math.Pow(float64(b.GetEncodingBaseLength()), float64(i)))
	}

	return decodedInt
}
