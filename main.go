package main

import (
  "./cmd"
)

func main() {
  err := cmd.RootCmd.GenBashCompletionFile("bin/ddev_completion.sh")
  if err != nil {
    panic(err)
  }
  cmd.Execute()
}