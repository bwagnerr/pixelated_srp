package main

import (
  "fmt"
  "flag"
  "pixelated/auth"
)

func initArgs() (string, string, string, string){
  var user, password, provider, cert string

  flag.StringVar(&user, "user", "", "User name")
  flag.StringVar(&password, "password", "", "Password")
  flag.StringVar(&provider, "provider", "", "Server provider")
  flag.StringVar(&cert, "cert", "", "Provider certificate")
  flag.Parse()

  return user, password, provider, cert
}

func main() {
  user, password, provider, cert := initArgs()

  serverInfo := auth.ServerInfo{"ha", false, 1}
  credentials := auth.Authenticate("username", "password", serverInfo)

  fmt.Println(credentials.SessionID)
  fmt.Println(credentials.UserToken)
  fmt.Println(credentials.Username)
  fmt.Println(credentials.UUID)

}
