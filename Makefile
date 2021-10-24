.PHONY: build push

COMMIT_SHA=$(shell git rev-parse --short HEAD)

build:
	docker build -t rg.fr-par.scw.cloud/evilmeow/owl-of-athena:$(COMMIT_SHA) .
	docker tag rg.fr-par.scw.cloud/evilmeow/owl-of-athena:$(COMMIT_SHA) rg.fr-par.scw.cloud/evilmeow/owl-of-athena:latest

push: build
	docker push rg.fr-par.scw.cloud/evilmeow/owl-of-athena:$(COMMIT_SHA)
	docker push rg.fr-par.scw.cloud/evilmeow/owl-of-athena:latest
