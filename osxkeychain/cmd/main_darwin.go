package main

import (
	"github.com/jojomomojo/docker-credential-helpers/credentials"
	"github.com/jojomomojo/docker-credential-helpers/osxkeychain"
)

func main() {
	credentials.Serve(osxkeychain.Osxkeychain{})
}
