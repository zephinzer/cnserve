DOCKER_NAMESPACE=zephinzer
PROJECT_NAME=$(notdir $(CURDIR))

start: build.dev
	-@docker stop $(PROJECT_NAME)-dev
	-@docker rm $(PROJECT_NAME)-dev
	-@$(eval PORT=$(shell cat $(CURDIR)/.env | grep 'PORT' | cut -f 2 -d '='))
	docker run \
		--env-file $(CURDIR)/.env \
		--workdir /go/src/$(PROJECT_NAME) \
		-p ${PORT}:${PORT} \
		-v $(CURDIR):/go/src/$(PROJECT_NAME) \
		-v $(CURDIR)/.cache:/.cache \
		-u $$(id -u) \
		--name $(PROJECT_NAME)-dev \
		$(PROJECT_NAME):dev-latest

start.prd: build.prd
	-@docker stop $(PROJECT_NAME)
	-@docker rm $(PROJECT_NAME)
	-$(eval PORT=$(shell cat $(CURDIR)/.env | grep 'PORT' | cut -f 2 -d '='))
	docker run \
		--env-file $(CURDIR)/.env \
		--workdir / \
		-p ${PORT}:${PORT} \
		-v $(CURDIR)/static:/static \
		-u $$(id -u) \
		--name $(PROJECT_NAME) \
		$(PROJECT_NAME):latest

dep.add: build.dev
	@if [ -z "${DEP}" ]; then \
		echo 'DEP variable not set.'; \
		exit 1; \
	else \
		docker run \
			--workdir /go/src/$(PROJECT_NAME) \
			-v $(CURDIR):/go/src/$(PROJECT_NAME) \
			-v $(CURDIR)/.cache:/.cache \
			-u $$(id -u) \
			--entrypoint=dep \
			$(PROJECT_NAME):dev-latest \
			ensure -add -v ${DEP}; \
	fi

dep.init: build.dev
	-@docker run \
		--workdir /go/src/$(PROJECT_NAME) \
		-v $(CURDIR):/go/src/$(PROJECT_NAME) \
		-v $(CURDIR)/.cache:/.cache \
		-u $$(id -u) \
		--entrypoint=dep \
		$(PROJECT_NAME):dev-latest \
		init
	if [ "$?" != "0" ]; then \
		echo "a"; \
	fi

shell:
	docker exec -it $(PROJECT_NAME)-dev /bin/bash

compile: build.dev
	docker run \
		--env PROJECT_NAME=$(PROJECT_NAME) \
		--env PORT=8080 \
		-p 8080:8080 \
		--env PATH_STATIC=./static \
		--workdir /go/src/$(PROJECT_NAME) \
		-v $(CURDIR):/go/src/$(PROJECT_NAME) \
		-u $$(id -u) \
		--entrypoint=go \
		$(PROJECT_NAME):dev-latest \
		build

publish: build.prd
	docker tag $(PROJECT_NAME):latest $(DOCKER_NAMESPACE)/$(PROJECT_NAME):latest
	docker push $(DOCKER_NAMESPACE)/$(PROJECT_NAME):latest

build.dev:
	mkdir -p $(CURDIR)/.cache/go-build
	docker build \
		--target development \
		--tag $(PROJECT_NAME):dev-latest .

build.prd:
	docker build \
		--target production \
		--tag $(PROJECT_NAME):latest .

build.local:
	GOPATH=$$(pwd) go build
