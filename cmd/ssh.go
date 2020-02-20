package cmd

import (
	"os"
  "fmt"
  "net/http"
  "os/exec"
  "github.com/spf13/cobra"
)

var sshCmd = &cobra.Command{
  Use:   "ssh",
  Short: "SSH into VM using VM's name",
  Long:  `SSH into VM using VM's name`,
  Run: func(cmd *cobra.Command, args []string) {
    vm_name := args[0]
    access_token := getToken()
    url := fmt.Sprintf("http://localhost:4003/api/instance_ip?name=%s&access_token=%s", vm_name, access_token)
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
    resp_body_map := readRespBody(resp)

    if resp_body_map["success"].(bool) {
      ip_address := resp_body_map["data"]
      remote_host := fmt.Sprintf("debian@%s", ip_address)

      //fetch default shell
      sh := os.Getenv("SHELL")
      cmd_str := fmt.Sprintf(`echo "%s" > /tmp/cakecloud_script`, generateAppleScript(remote_host))
      cmd := exec.Command(sh, "-c", cmd_str)
      cmd.Run()
      cmd = exec.Command("osascript", "/tmp/cakecloud_script")
      cmd.Run()
      cmd = exec.Command("rm", "/tmp/cakecloud_script")
      cmd.Run()
    } else {
      error_message := resp_body_map["message"]
      fmt.Println(error_message)
    }
  },
}