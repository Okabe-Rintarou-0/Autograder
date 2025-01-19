PROJECT_DIR := $(shell pwd)
lint:
	goimports -local autograder -w .

gen:
	go run $(PROJECT_DIR)/pkg/gen/main.go