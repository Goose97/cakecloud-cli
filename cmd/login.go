package cmd

import (
  "fmt"
  "syscall"
  "net/http"
  "io/ioutil"
  "bytes"
	"encoding/json"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var logInCmd = &cobra.Command{
  Use:   "login",
  Short: "Provide authentication - use cakecloud username and password",
  Long:  `Provide authentication - use cakecloud username and password`,
  Run: func(cmd *cobra.Command, args []string) {
    var username string
    var password string

    if isAlreadyLogin() {
      return
    }

    fmt.Print("Username: ")
    fmt.Scanln(&username)

    fmt.Print("Password: ")
    bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
    if err != nil {
      fmt.Println(err)
    }
    password = string(bytePassword)

    values := map[string]string{"username": username, "password": password}
    jsonValue, _ := json.Marshal(values)
		resp, err := http.Post("http://10.1.9.49:22:4001/api/login", "application/json", bytes.NewBuffer(jsonValue))
		
		if err != nil {
			fmt.Println(err)
		} else {
			resp_body_map := readRespBody(resp)
			if error, ok := resp_body_map["error"]; ok {
				fmt.Print("\n")
				fmt.Println(error)
			} else {
				accessToken := resp_body_map["jwt"].(string)
				accessTokenInByte := []uint8(accessToken)
				err = ioutil.WriteFile("/tmp/cakecloud.token", accessTokenInByte, 0644)
				if err != nil {
					panic(err)
				}
				fmt.Print("\n")
				fmt.Println(fmt.Sprintf("Log in success as %s", username))
			}
		}
  },
}