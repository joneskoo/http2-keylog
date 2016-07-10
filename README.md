# Go crypto/tls SSL key log Proof of Concept

With [a very simple modification](https://github.com/joneskoo/http2-keylog/commit/d75943121890354e1d9c4eed1bb7281e5eb95761) to crypto/tls we can dump
TLS secrets in a format Wireshark can import. This makes
debugging TLS application issues much easier!

Issue: [golang/go #13057 crypto/tls: support NSS-formatted key log file](https://github.com/golang/go/issues/13057)

```bash
$ make
go run vendor/crypto/tls/generate_cert.go -host localhost
2016/07/10 09:51:45 written cert.pem
2016/07/10 09:51:45 written key.pem
$ make run
go run main.go
2016/07/10 09:51:50 About to listen on 10443. Go to https://[::1]:10443/
```

Meanwhile in another terminal

```bash
$ curl -k 'https://[::1]:10443/'
This is an example server.
```

Now the test server dumps client random and TLS master secret to a file

```bash
2016/07/10 09:51:53 Leaked TLS secrets to  ssl-key-log.txt
^Csignal: interrupt
make: *** [run] Error 1

$ cat ssl-key-log.txt
CLIENT_RANDOM 5781f0898847c3eeea2f5b51b531d9b14f76bc1fd23af2e3896b7871d022e1ad c4dec1b90c263251ae38be20b54e2e7f861c4953042f4fd8a14bcc8c60a86691eb6bb6073e45258e7bfbade1e984987a
```

Boom, TLS decrypted in Wireshark!

![Wireshark showing decrypted TLS](wireshark.png)
