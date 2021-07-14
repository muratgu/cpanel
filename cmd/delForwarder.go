/*
Copyright © 2021 muratgu <mgungora@gmail.com>

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
	"log"
	"os"
	"net/url"
	"github.com/spf13/cobra"
)

var delForwarderCmd = &cobra.Command{
	Use:   "del-forwarder [from] [to]",
	Short: "Delete an email forwarder",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		domain, _ := cmd.Flags().GetString("domain")
		email := args[0]
		fwdemail := args[1]
		data := url.Values{}
    	data.Set("domain", domain)
		data.Set("address", email)
		data.Set("forwarder", fwdemail)
		if resp, err := Get("Email/delete_forwarder", data); err != nil {
			log.Fatal(err)
		} else {
			JsonEncode(resp, os.Stdout)
		}
	},
}

func init() {
	rootCmd.AddCommand(delForwarderCmd)
}
