package main

import (
	"github.com/jojomomojo/docker-credential-helpers/credentials"
	"github.com/jojomomojo/docker-credential-helpers/wincred"
)

func main() {
	credentials.Serve(wincred.Wincred{})
}
