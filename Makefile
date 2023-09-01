# Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

.PHONY: help clean validator runv generator rung validate generate fmt regen regression run-regression

VALIDATOR_APP := validator
GENERATOR_APP := generator
OPENAPI_YAML := /local/api/mc-openapi.yaml
GOFMT_OPTS := gofmt -w .
REGRESSION_APP := api-regression
GENERATE_USER := $(shell id -u ${USER}):$(shell id -g ${USER})
GENERATOR_VERSION := v6.5.0

help:
	@echo ""
	@echo "-----------------------------------------------------------------------------------"
	@echo "make clean           - remove all"
	@echo "make validator       - build a local validator executable"
	@echo "make runv            - run the local validator executable"
	@echo "make generator       - build a local generator executable"
	@echo "make rung            - run the local generator executable"
	@echo "make validate        - validate the openapi specification using openapi-generator-cli"
	@echo "make generate        - generate go code for the openapi specification using openapi-generator-cli"
	@echo "make fmt             - Run gofmt"	
	@echo "make regen           - Run the generator, validate, generate, and then run the validator"
	@echo "make regression      - Build api-regression"
	@echo "make run-regression  - Run api-regression"
	@echo ""

clean:
	@echo "Clean up..."
	go clean
	rm -f $(APP_NAME) $(VALIDATOR_APP)-app $(GENERATOR_APP)-app $(REGRESSION_APP)

validator: clean
	@echo "Build local '$(VALIDATOR_APP)-app' executable..."
	go build -o $(VALIDATOR_APP)-app -ldflags "-X main.buildTime=`date -u '+%Y-%m-%dT%H:%M:%S'`" ./internal/$(VALIDATOR_APP)/cmd/main.go
	ls -lh ./$(VALIDATOR_APP)-app

runv: validator
	@echo "Running '$(VALIDATOR_APP)-app'..."
	./$(VALIDATOR_APP)-app -config internal/$(VALIDATOR_APP)/$(VALIDATOR_APP).conf	

generator: clean
	@echo "Build local '$(GENERATOR_APP)-app' executable..."
	go build -o $(GENERATOR_APP)-app -ldflags "-X main.buildTime=`date -u '+%Y-%m-%dT%H:%M:%S'`" ./internal/$(GENERATOR_APP)/cmd/main.go
	ls -lh ./$(GENERATOR_APP)-app

rung: generator
	@echo "Running '$(GENERATOR_APP)-app'..."
	./$(GENERATOR_APP)-app -config internal/$(GENERATOR_APP)/$(GENERATOR_APP).conf	

validate:
	@echo "Validating $(OPENAPI_YAML) using openapi-generator-cli"
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:$(GENERATOR_VERSION) version
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:$(GENERATOR_VERSION) validate -i $(OPENAPI_YAML)

generate:
	@echo "Generating $(OPENAPI_YAML) using openapi-generator-cli"
	rm -rf pkg/client
	docker run -u $(GENERATE_USER) --rm -v ${PWD}:/local openapitools/openapi-generator-cli:$(GENERATOR_VERSION) generate -i $(OPENAPI_YAML) -g go -o /local/pkg/client --package-name client --ignore-file-override=/local/.openapi-generator-ignore
	@echo "Format files after generation to conform to project standard"
	$(GOFMT_OPTS)

fmt:
	@echo "Format check"
	$(GOFMT_OPTS)

regen: rung validate generate runv
	@echo "Full Regeneration of $(OPENAPI_YAML)"

regression:
	@echo "Build local $(REGRESSION_APP)..."
	go build -o api-regression cmd/$(REGRESSION_APP)/main.go

run-regression: regression
	# ./$(REGRESSION_APP) -debug 4 -config i1.conf --ginkgo.v --ginkgo.fail-fast --ginkgo.focus "v1"
	# ./$(REGRESSION_APP) -debug 4 -config i1.conf --ginkgo.v --ginkgo.fail-fast --ginkgo.focus "v2"
	./$(REGRESSION_APP) -debug 4 -config $(REGRESSION_APP).conf --ginkgo.v --ginkgo.fail-fast 
