package main

import (
	"github.com/jojomomojo/docker-credential-helpers/credentials"
	"github.com/jojomomojo/docker-credential-helpers/pass"
)

func main() {
	credentials.Serve(pass.Pass{})
}
