setup:
	make -C  ../../hack/openshift operators resources
	kubectl apply -f bad-deployment.yaml
	make -C ../.. deploy
