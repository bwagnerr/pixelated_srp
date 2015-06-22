package main

import (
  "fmt"
  "pixelated/auth"
)

func main() {
  serverInfo := auth.ServerInfo{"ha", false, 1}
  credentials := auth.Authenticate("username", "password", serverInfo)

  fmt.Println(credentials.SessionID)
  fmt.Println(credentials.UserToken)
  fmt.Println(credentials.Username)
  fmt.Println(credentials.UUID)

}
