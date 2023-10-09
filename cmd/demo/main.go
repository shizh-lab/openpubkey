package main

import (
	"encoding/pem"
	"os"

	"github.com/openpubkey/openpubkey/parties"
	"github.com/openpubkey/openpubkey/pktoken"
)

func main() {
	signer, err := pktoken.NewSigner("./pktoken.cfg", "ES256", true, map[string]any{"extra": "yes"})
	if err != nil {
		panic(err)
	}
	op, err := parties.NewGithubOpFromEnvironment()
	if err != nil {
		panic(err)
	}
	client := &parties.OpkClient{
		Op:     op,
		Signer: signer,
	}
	pktBytes, err := client.OidcAuth()
	if err != nil {
		panic(err)
	}

	pem.Encode(os.Stdout, &pem.Block{
		Type:  "PK Token",
		Bytes: pktBytes,
	})
}
