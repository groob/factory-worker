// see http://timeclock.sourceforge.net/
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	var (
		serverURL = flag.String("server-url", "", "url for timeclock")
		name      = flag.String("name", "", "your name(first, last)")
		password  = flag.String("password", "", "your password")
	)
	flag.Parse()

	var v url.Values
	{
		v.Set("left_displayname", *name)
		v.Set("employee_passwd", *password)
		v.Set("left_inout", "in")
		v.Set("left_notes", "")
	}

	client := &http.Client{}
	if err := submit(client, *serverURL, v); err != nil {
		log.Fatal(err)
	}

}

func submit(client *http.Client, serverURL string, v url.Values) error {
	if client == nil {
		client = &http.Client{}
	}
	resp, err := client.PostForm(serverURL, v)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("could not sign you in to %q, server responded with %q\n", serverURL, resp.Status)
	}
	return nil
}
