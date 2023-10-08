package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using variable-length code",
	Run:   pack,
}

func init() {
	packCmd.AddCommand(vlcCmd)
}

const packedExtension = "vlc"

var ErrEmptyPath = errors.New("path to file is not specified")

func pack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handelErr(ErrEmptyPath)
	}
	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handelErr(err)
	}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		handelErr(err)
	}

	//packed := Encode(data)
	packed := ""
	fmt.Println(string(data)) // TODO: Remove

	err = ioutil.WriteFile(packedFileName(filePath), []byte(packed), 0644)
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
