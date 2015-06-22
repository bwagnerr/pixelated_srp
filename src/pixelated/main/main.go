package main

import (
  "fmt"
  "flag"
  "pixelated/auth"
)

func initArgs() (string, string, string, string){
  var user, password, provider, cert string
  var skipCertValidation bool

  flag.StringVar(&user, "user", "", "User name")
  flag.StringVar(&password, "password", "", "Password")
  flag.StringVar(&provider, "provider", "", "Server provider")
  flag.StringVar(&cert, "cert", "", "Provider certificate")
  flag.BoolVar(&skipCertValidation, "insecure", false, "Skip certificate validation")
  flag.Parse()

  return user, password, provider, cert
}

func main() {
  user, password, provider, cert := initArgs()

  serverInfo := auth.ServerInfo{provider, cert, 1}
  credentials := auth.Authenticate(user, password, serverInfo)

  fmt.Println(credentials.SessionID)
  fmt.Println(credentials.UserToken)
  fmt.Println(credentials.Username)
  fmt.Println(credentials.UUID)

}
