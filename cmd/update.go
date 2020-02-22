package cmd

import (
  "fmt"
	"io/ioutil"
	"net/http"
	"github.com/spf13/cobra"
)

type UpdateInstanceResponse struct {
	success bool
	data []byte
	message *string
}

var updateCmd = &cobra.Command{
  Use:   "update",
  Short: "Upate list VM for autocomplete",
  Long:  `Upate list VM for autocomplete`,
  Run: func(cmd *cobra.Command, args []string) {
		access_token := getToken()

		url := fmt.Sprintf("http://10.1.9.49:22:4001/api/instance_name?access_token=%s", access_token)
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		resp_body_map := readRespBody(resp)
		
		if resp_body_map["success"].(bool) == true {
			list_vm_names := resp_body_map["data"].(string)
			list_vm_names_byte := []uint8(list_vm_names)
			err = ioutil.WriteFile("~/.oh-my-zsh/plugins/cakecloud/list_vm", list_vm_names_byte, 0644)
			if err != nil {
				panic(err)
			}
		} else {
      error_message := resp_body_map["message"]
      fmt.Println(error_message)
    }
  },
}