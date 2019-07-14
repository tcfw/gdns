build:
	docker build -t tcfw/gdns:latest .

.PHONY: push
push:
	docker push tcfw/gdns:latest