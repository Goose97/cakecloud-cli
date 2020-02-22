package cmd

import (
  "os"
  "fmt"
  "github.com/spf13/cobra"
  // "github.com/dgrijalva/jwt-go"
)

func init() {
  RootCmd.AddCommand(versionCmd)
  RootCmd.AddCommand(logInCmd)
  RootCmd.AddCommand(sshCmd)
  RootCmd.AddCommand(updateCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of Hugo",
  Long:  `All software has versions. This is Hugo's`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
  },
}

var RootCmd = &cobra.Command{
  Use:   "hugo",
  Short: "Hugo is a very fast static site generator",
  Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
  Run: func(cmd *cobra.Command, args []string) {
    // Do Stuff Here
    fmt.Println("Hello CLI")
  },
}

func Execute() {
  if err := RootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

