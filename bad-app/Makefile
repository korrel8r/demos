IMAGE ?= quay.io/korrel8r/bad-app:latest
NAMESPACE = app-to-infra

.PHONY: all build push deploy undeploy logs

all: deploy

build:
	podman build -t $(IMAGE) .

push: build
	podman push $(IMAGE)

deploy:
	kubectl create namespace $(NAMESPACE) --dry-run=client -o yaml | kubectl apply -f -
	kubectl -n $(NAMESPACE) apply -f deployment.yaml

undeploy:
	kubectl delete namespace $(NAMESPACE) --ignore-not-found

logs:
	kubectl -n $(NAMESPACE) logs -l app=bad-app -f

URL=$(shell oc whoami --show-console)/k8s/ns/app-to-infra/deployments/bad-app
url:
	@echo $(URL)

browse: deploy
	@xdg-open $(URL)
	@sleep 1											# xdg-open bug
