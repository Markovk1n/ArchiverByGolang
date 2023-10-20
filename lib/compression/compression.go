package compression

type Encoder interface {
	Encode(srt string) []byte
}

type Decoder interface {
	Decode(data []byte) string
}
