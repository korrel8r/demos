SRC=$(error Set SRC to korrel8r repository directory.)
setup:
	make -C  $(SRC)/hack/openshift operators resources
	kubectl apply -f bad-deployment.yaml
	make -C $(SRC) deploy
