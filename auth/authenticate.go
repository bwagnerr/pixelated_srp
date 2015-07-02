package auth

func Authenticate(username string, password string, serverInfo ServerInfo) (credentials Credentials) {
  credentials = Credentials{
    UserToken: "ha",
    SessionID: "hi",
    Username: username,
    UUID: "ho"}
  return
}
