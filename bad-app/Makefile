IMAGE ?= quay.io/korrel8r/bad-app:latest
NAMESPACE = default

.PHONY: all build push deploy undeploy logs

all: deploy

image:
	podman build -t $(IMAGE) .
	podman push $(IMAGE)

deploy: image
	kubectl create namespace $(NAMESPACE) --dry-run=client -o yaml | kubectl apply -f -
	kubectl -n $(NAMESPACE) apply -f deployment.yaml

undeploy:
	kubectl delete deployment/bad-app -n $(NAMESPACE) --ignore-not-found

logs:
	kubectl -n $(NAMESPACE) logs -l app=bad-app -f

URL=$(shell oc whoami --show-console)/k8s/ns/default/deployments/bad-app
url:
	@echo $(URL)

browse:
	@xdg-open $(URL)
	@sleep 1											# xdg-open bug
