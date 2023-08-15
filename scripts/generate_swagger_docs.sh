cd cmd && \
	swag init && \
	cd .. &&  \
	mv ./cmd/docs/* ./docs  &&\
	rm -rf ./cmd/docs && \
	cp -f ./docs/swagger.yaml ./api/swagger.yaml