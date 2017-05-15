# goforeman
Go bindings for Foreman REST API 

Heavily inspired by DigitalOcean GoDo : https://github.com/digitalocean/godo

!! Work In Progress !!

## Usage

```go
package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/amine7536/goforeman"
)

type AuthTransport struct {
	*http.Transport
	Username string
	Password string
}

func (t AuthTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.SetBasicAuth(t.Username, t.Password)
	return t.Transport.RoundTrip(r)
}

func main() {

	foremanURL := os.Getenv("FOREMAN_URL")
	foremanUser := os.Getenv("FOREMAN_USER")
	foremanPassword := os.Getenv("FOREMAN_PASSWORD")

	client := http.Client{
		Transport: AuthTransport{
			&http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			foremanUser,
			foremanPassword,
		},
	}

	foreman := goforeman.New(&client, foremanURL)

	ctx := context.TODO()

	hosts, _, err := foreman.Hosts.List(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, h := range hosts {
		v, _, _ := foreman.Hosts.Get(ctx, h.Name)
		if err != nil {
			log.Fatalln(err.Error())
		}
		fmt.Println(v)
	}

	d, _, err := foreman.Dashboard.Get(ctx)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	fmt.Println(d)

	facts, _, err := foreman.Facts.Get(ctx, "foremanserver.lab.local.dev")
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	fmt.Println(facts)

}
```