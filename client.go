package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var sslKeyLogfile = flag.String("keylog", "ssl-keylog.txt", "File name to write NSS key log format log of TLS keys")

func main() {
	flag.Parse()
	if len(os.Args) != 2 {
		log.Fatalln("Usage: client <url>")
	}

	f, err := os.OpenFile(*sslKeyLogfile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(f, "# SSL/TLS secrets log file, generated by go\n")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{KeyLogWriter: f},
	}
	client := &http.Client{Transport: tr}

	res, err := client.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}
