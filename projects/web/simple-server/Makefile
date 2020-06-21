
.PHONY:  build_keys test

build_keys:
	test -f server.key || openssl req -new -newkey rsa:4096 -days 365 -nodes -x509 \
		-subj "/C=US/ST=Denial/L=Springfield/O=Dis/CN=www.example.com" \
		-keyout server.key  -out server.crt

test: build_keys
	go build main.go
	sudo ./main &
	curl -kL https://127.0.0.1/