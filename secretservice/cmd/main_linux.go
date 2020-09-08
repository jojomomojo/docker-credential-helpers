package main

import (
	"github.com/jojomomojo/docker-credential-helpers/credentials"
	"github.com/jojomomojo/docker-credential-helpers/secretservice"
)

func main() {
	credentials.Serve(secretservice.Secretservice{})
}
