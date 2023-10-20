package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"github/Markovk1n/ArchiverByGolang/lib/compression"
	"github/Markovk1n/ArchiverByGolang/lib/compression/vlc"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack file",
	Run:   pack,
}

const packedExtension = "vlc"

var ErrEmptyPath = errors.New("path to file is not specified")

func pack(cmd *cobra.Command, args []string) {
	var encoder compression.Encoder
	if len(args) == 0 || args[0] == "" {
		handelErr(ErrEmptyPath)
	}

	method := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		encoder = vlc.New()
	default:
		cmd.PrintErr("unknown method")
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

	packed := encoder.Encode(string(data))

	err = os.WriteFile(packedFileName(filePath), (packed), 0644)
	if err != nil {
		handelErr(err)
	}
}

func packedFileName(path string) string {
	// path = /path/to/file/myFile.txt
	fileName := filepath.Base(path) // myFile.txt
	//ext := filepath.Ext(fileName)					// .txt
	//baseName := strings.TrimSuffix(fileName, ext)	// 'myFile.txt' - '.txt' = 'myFile'
	//return baseName +" "+packedExtension

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtension
}

func init() {
	rootCmd.AddCommand(packCmd)
	packCmd.Flags().StringP("method", "m", "", "compression method: vlc ")
	if err := packCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}
