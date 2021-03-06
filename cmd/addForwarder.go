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
	"github.com/spf13/cobra"
	"log"
	"net/url"
	"os"
)

var addForwarderCmd = &cobra.Command{
	Use:   "add-forwarder [from] [to]",
	Short: "Add an email forwarder",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		domain, _ := cmd.Flags().GetString("domain")
		email := args[0]
		fwdemail := args[1]
		data := url.Values{}
		data.Set("domain", domain)
		data.Set("email", email)
		data.Set("fwdemail", fwdemail)
		data.Set("fwdopt", "fwd")
		if resp, err := Get("Email/add_forwarder", data); err != nil {
			log.Fatal(err)
		} else {
			JsonEncode(resp, os.Stdout)
		}
	},
}

func init() {
	rootCmd.AddCommand(addForwarderCmd)
}
