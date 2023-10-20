package cmd

import (
	"github.com/spf13/cobra"
	"github/Markovk1n/ArchiverByGolang/lib/compression"
	"github/Markovk1n/ArchiverByGolang/lib/compression/vlc"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Unpack file",
	Run:   unpack,
}

// TODO: take extension from file
const unpackedExtension = "txt"

func unpack(cmd *cobra.Command, args []string) {
	var decoder compression.Decoder
	if len(args) == 0 || args[0] == "" {
		handelErr(ErrEmptyPath)
	}
	filePath := args[0]
	method := cmd.Flag("method").Value.String()
	switch method {
	case "vlc":
		decoder = vlc.New()
	default:
		cmd.PrintErr("unknown method")
	}

	r, err := os.Open(filePath)
	if err != nil {
		handelErr(err)
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		handelErr(err)
	}

	packed := decoder.Decode(data)

	err = os.WriteFile(unpackedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handelErr(err)
	}
}

// TODO: refactor this
func unpackedFileName(path string) string {
	// path = /path/to/file/myFile.txt
	fileName := filepath.Base(path) // myFile.txt
	//ext := filepath.Ext(fileName)					// .txt
	//baseName := strings.TrimSuffix(fileName, ext)	// 'myFile.txt' - '.txt' = 'myFile'
	//return baseName +" "+packedExtension

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + unpackedExtension
}

func init() {
	rootCmd.AddCommand(unpackCmd)

	unpackCmd.Flags().StringP("method", "m", "", "decompression method: vlc ")
	if err := packCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}
