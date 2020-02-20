package cmd

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"golang.org/x/crypto/ssh"
)

func isAlreadyLogin() bool {
  _, err := ioutil.ReadFile(getTokenFilePath())
  // var verify_token_func *jwt.Keyfunc = func(token []byte) {
  //   fmt.Println(token)
  //   return token
  // }
  
  
  if err == nil { 
    // fmt.Println(data)
    // string_data := string([]byte(data[:]))
    // parsed, err := jwt.Parse(string_data, verify_token_func)
    fmt.Println("You already logged in")
    return true 
  } else {
    return false
  }
}

func getTokenFilePath() string {  
  return "/tmp/cakecloud.token"
}

func getToken() string {  
  data, err := ioutil.ReadFile(getTokenFilePath())
  if err != nil {
    panic(err)
  } else {
    return string([]byte(data[:])) 
  }
}

func readRespBody(resp *http.Response) map[string]interface{} {
  defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
  var m map[string]interface{}
  if err := json.Unmarshal(body, &m); err != nil {
    panic(err)
	}
	
  return m
}

func generateAppleScript(remote_host string) string {
  return fmt.Sprintf(
    `tell application \"iTerm2\"
      tell first session of current tab of current window
        write text \"ssh %s\"
      end tell
    end tell
    `, remote_host)
}

func publicKey(path string) ssh.AuthMethod {
  key, err := ioutil.ReadFile(path)
  if err != nil {
   panic(err)
  }

  signer, err := ssh.ParsePrivateKey(key)
  if err != nil {
   panic(err)
  }

  return ssh.PublicKeys(signer)
}