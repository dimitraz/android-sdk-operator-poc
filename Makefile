TAG = 1.0.0
DOCKERORG = dimitraz
IMAGE_NAME = android-sdk-operator
BINARY_NAME = android-sdk-operator
BIN_DIR="tmp/_output/bin"

.phony: build_binary
build_binary:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(BIN_DIR)/$(BINARY_NAME) ./cmd/android-sdk-operator/

.phony: build_image
build: build_binary
	docker build -t $(DOCKERORG)/$(IMAGE_NAME):$(TAG) -f tmp/build/Dockerfile .

.phony: docker_push
push:
	@docker login --username $(DOCKERHUB_USERNAME) --password $(DOCKERHUB_PASSWORD)
	docker push $(DOCKERORG)/$(IMAGE_NAME):$(TAG)

.phony: release
release: build_image docker_push