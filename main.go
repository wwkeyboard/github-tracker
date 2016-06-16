// You can edit this code!
// Click here and start typing.
package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	conf, err := parseArgs()
	if err != nil {
		log.Panic(err)
	}

	stuff, err := getStuff(conf.Token)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("found%v\n", stuff)
}

func getStuff(token string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.github.com", nil)

	req.Header.Add("Authorization: token", token)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return string(body), err
}

// Config is the application's configuration
type Config struct {
	Token string
}

func parseArgs() (*Config, error) {
	// flag
	token := flag.String("token", "token", "A github oauth token")

	flag.Parse()

	if token == nil {
		return nil, errors.New("You must specify a token")
	}

	return &Config{
		Token: *token,
	}, nil
}
