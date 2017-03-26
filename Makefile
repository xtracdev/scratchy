image:
	cp /etc/ssl/certs/ca-certificates.crt .
	docker build -t $(IMAGE_TAG) .
