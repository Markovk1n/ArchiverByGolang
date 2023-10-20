package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type BinaryChunks []BinaryChunk

type BinaryChunk string

const chunksSize = 8
const hexChunkSeparator = " "

type encodingTable map[rune]string

func NewBinChunks(data []byte) BinaryChunks {

	res := make(BinaryChunks, 0, len(data))
	for _, code := range data {
		res = append(res, NewBinChunk(code))
	}
	return res
}

func NewBinChunk(code byte) BinaryChunk {
	return BinaryChunk(fmt.Sprintf("%08b", code))
}

func (bcs BinaryChunks) Bytes() []byte {
	res := make([]byte, 0, len(bcs))
	for _, bc := range bcs {
		res = append(res, bc.Byte())
	}
	return res
}

func (bc BinaryChunk) Byte() byte {
	num, err := strconv.ParseUint(string(bc), 2, chunksSize)
	if err != nil {
		panic("can't parse binary chunk: " + err.Error())
	}
	return byte(num)
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
//
//	func (hcs HexChunks) ToBinary() BinaryChunks {
//		res := make(BinaryChunks, 0, len(hcs))
//		for _, chunk := range hcs {
//			binChunk := chunk.ToBinary()
//			res = append(res, binChunk)
//		}
//		return res
//	}
//
// Join joins chunks into one line and returns as string
func (bcs BinaryChunks) Join() string {
	var buf strings.Builder
	for _, bc := range bcs {
		buf.WriteString(string(bc))
	}
	return buf.String()
}

//
//func (hc HexChunk) ToBinary() BinaryChunk {
//	num, err := strconv.ParseUint(string(hc), 16, chunksSize)
//	if err != nil {
//		panic("can't parse hex chunk: " + err.Error())
//	}
//	res := fmt.Sprintf("%08b", num)
//	return BinaryChunk(res)
//}
//
//func (hcs HexChunks) ToString() string {
//
//	switch len(hcs) {
//	case 0:
//		return ""
//	case 1:
//		return string(hcs[0])
//	}
//
//	var buf strings.Builder
//	buf.WriteString(string(hcs[0]))
//
//	for _, hc := range hcs[1:] {
//		buf.WriteString(hexChunkSeparator)
//		buf.WriteString(string(hc))
//
//	}
//	return buf.String()
//}
//
//func (bcs BinaryChunks) ToHex() HexChunks {
//	res := make(HexChunks, 0, len(bcs))
//
//	for _, chunk := range bcs {
//		hexChunk := chunk.ToHex()
//		res = append(res, hexChunk)
//	}
//	return res
//}
//
//func (bc BinaryChunk) ToHex() HexChunk {
//	num, err := strconv.ParseUint(string(bc), 2, chunksSize)
//	if err != nil {
//		panic("can't parse binary chunk: " + err.Error())
//	}
//	res := strings.ToUpper(fmt.Sprintf("%x", num))
//
//	if len(res) == 1 {
//		res = "0" + res
//	}
//
//	return HexChunk(res)
//}

// splitByChunks split binary string by chunks with given size,
// i.g.: `100101011001010110010101` -> `10010101 10010101 10010101
func splitByChunks(bStr string, chunkSize int) BinaryChunks {
	strLen := utf8.RuneCountInString(bStr)
	chunksCount := strLen / chunkSize

	if strLen/chunkSize != 0 {
		chunksCount++
	}

	res := make(BinaryChunks, 0, chunksCount)
	var buf strings.Builder

	for i, ch := range bStr {
		buf.WriteString(string(ch))

		if (i+1)%chunkSize == 0 {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}
	if buf.Len() != 0 {
		lastChunk := buf.String()
		lastChunk += strings.Repeat("0", chunkSize-len(lastChunk))
		res = append(res, BinaryChunk(lastChunk))
	}
	return res
}
