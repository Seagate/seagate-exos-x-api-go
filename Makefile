# Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

.PHONY: help clean validator runv generator rung validate generate fmt regen regression run-regression

VALIDATOR_APP := validator
GENERATOR_APP := generator
OPENAPI_YAML := api/mc-openapi.yaml
GOFMT_OPTS := gofmt -w .

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
	rm -f $(APP_NAME)

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
	openapi-generator-cli version
	openapi-generator-cli validate -i $(OPENAPI_YAML)

generate:
	@echo "Generating $(OPENAPI_YAML) using openapi-generator-cli"
	openapi-generator-cli generate -i $(OPENAPI_YAML) -g go -o pkg/client --package-name client --ignore-file-override .openapi-generator-ignore
	@echo "Format files after generation to conform to project standard"
	$(GOFMT_OPTS)

fmt:
	@echo "Format check"
	$(GOFMT_OPTS)

regen: rung validate generate runv
	@echo "Full Regeneration of $(OPENAPI_YAML)"

regression:
	@echo "Build local api-regression..."
	go build -o api-regression cmd/api-regression/main.go

run-regression: regression
	# ./api-regression -debug 4 -config i1.conf --ginkgo.v --ginkgo.fail-fast --ginkgo.focus "v1"
	# ./api-regression -debug 4 -config i1.conf --ginkgo.v --ginkgo.fail-fast --ginkgo.focus "v2"
	./api-regression -debug 4 -config api-regression.conf --ginkgo.v --ginkgo.fail-fast 
