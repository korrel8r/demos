CONFIG=~/src/korrel8r/config/openshift
setup:
	kubectl apply -f bad-deployment.yaml

all: setup
	make -C $(CONFIG) all
