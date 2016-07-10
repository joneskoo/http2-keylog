cert.pem:
	go run vendor/crypto/tls/generate_cert.go -host localhost

run: cert.pem
	go run main.go
