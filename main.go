package main

import "net/http"
import "fmt"
import "os"
import "time"

func main() {
	timeout := time.Duration(200 * time.Millisecond)
	client := http.Client{
		    Timeout: timeout,
	    }
	_, err := client.Get(fmt.Sprintf("http://%s:%s/", os.Args[3], os.Args[4]))

	if err != nil {
		os.Exit(1)
	}
}
