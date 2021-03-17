package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/oauth2/google"
)

func main() {
	params := &params{}
	params.parse()

	p, err := os.ReadFile(params.jwtfile)
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()

	cfg, err := google.JWTConfigFromJSON(p, params.scopes...)
	if err != nil {
		log.Fatalln(err)
	}
	cfg.Subject = params.acct

	client := cfg.Client(ctx)

	resp, err := client.Get(params.endpoint)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Response status code:", resp.Status, resp.StatusCode)
	fmt.Println("Body:", string(body))
}

type params struct {
	jwtfile  string
	acct     string
	scopes   []string
	endpoint string
}

func (p *params) parse() {
	flag.StringVar(&p.jwtfile, "jwtfile", "", "The credentials JWT file.")
	flag.StringVar(&p.acct, "delegated_account", "", "The delegated account.")
	scopes := flag.String("scopes", "", "A comma-sepparated list of scopes.")
	flag.StringVar(&p.endpoint, "endpoint", "", "The endpoint to test. (Only GET endpoints)")

	flag.Parse()

	fillFromEnvIfEmpty(&p.jwtfile, "JWT_FILE")
	fillFromEnvIfEmpty(&p.acct, "DELEGATED_ACCOUNT")
	fillFromEnvIfEmpty(scopes, "SCOPES")
	fillFromEnvIfEmpty(&p.endpoint, "ENDPOINT")

	p.scopes = strings.Split(*scopes, ",")
}

func fillFromEnvIfEmpty(v *string, envkey string) {
	if v == nil {
		return
	}

	if *v == "" {
		*v = os.Getenv(envkey)
	}
}
