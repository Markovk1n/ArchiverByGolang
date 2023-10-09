package cmd

import (
	"github.com/spf13/cobra"
	"github/Markovk1n/ArchiverByGolang/lib/vlc"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var vlcUnpackCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Unpack file using variable-length code",
	Run:   unpack,
}

// TODO: take extension from file
const unpackedExtension = "txt"

func unpack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handelErr(ErrEmptyPath)
	}
	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handelErr(err)
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		handelErr(err)
	}

	packed := vlc.Decode(string(data))

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
	unpackCmd.AddCommand(vlcUnpackCmd)
}
