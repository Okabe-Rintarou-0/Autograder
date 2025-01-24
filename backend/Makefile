PROJECT_DIR := $(shell pwd)
lint:
	goimports -local autograder -w .

gen:
	go run $(PROJECT_DIR)/pkg/gen/main.go

build:
	$(shell) $(PROJECT_DIR)/build.sh

bootstrap:
	$(shell) $(PROJECT_DIR)/bootstrap.sh