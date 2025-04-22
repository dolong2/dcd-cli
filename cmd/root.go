/*
Copyright © 2024 dolong2

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"errors"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dcd",
	Short: "DCD를 사용하기 위한 CLI",
	Long: `
#####      ####   #####               ####    ###       ##                         ##
## ##    ##  ##   ## ##             ##  ##    ##                                  ##
##  ##  ##        ##  ##           ##         ##      ###      ####    #####     #####
##  ##  ##        ##  ##           ##         ##       ##     ##  ##   ##  ##     ##
##  ##  ##        ##  ##           ##         ##       ##     ######   ##  ##     ##
## ##    ##  ##   ## ##             ##  ##    ##       ##     ##       ##  ##     ## ##
#####      ####   #####               ####    ####     ####     #####   ##  ##      ###

DCD를 사용하기 위한 CLI입니다.
더 많은 정보는 github 링크를 참고해주세요. https://github.com/dolong2/dcd-cli`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		var cmdErr *cmdError.CmdError
		if errors.As(err, &cmdErr) {
			os.Exit(cmdErr.Code)
		} else {
			os.Exit(1)
		}
	}
	os.Exit(0)
}

func init() {
}
