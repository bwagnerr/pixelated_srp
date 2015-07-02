package auth

import "os"

type ServerInfo struct {
	APIUrl     string
	VerifyCert string
	APIVersion int
}

func verifyCertFileExist(cert string) {
	if cert != "" {
		_, err = os.Stat(cert)
		if os.IsNotExist(err) {
			os.Exit(1)
		}
	}
}
