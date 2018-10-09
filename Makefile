# Makefile variables
PROJECT_NAME=duck-tracker

# List of packages 1 package per line relative to current location
PKG_ML = $(shell go list ./... | sed "s%_$$(pwd)%\.%g" | grep -v -e "vendor*")
# List of packages space delimited
PKG = $(shell echo ${PKG_ML} | tr "\n" " ")
# All .go files, excluding the vendors
GOFILES = $(shell find . -type f -name '*.go' -not -path "./vendor*/*")

# Environment variables that should be set
DOCKER_USER ?=
DOCKER_PASS ?=
GIT_USER ?=
GIT_PASS ?=

# Run on docker
RUN = docker-compose run ${OPTS} --name ${PROJECT_NAME}_$(shell od -An -N1 -i /dev/random | tr -dc '0-9') --rm ${PROJECT_NAME}

#
## Basic commands with docker
all: setup
# Setup the project
setup:
	docker-compose build
	${RUN} make compile
# Install the image and run the project
install:
	${RUN} bash -c "go build -o dist/${PROJECT_NAME} && ./dist/${PROJECT_NAME}"
clean:
	rm -rf reports/ dist/
	docker-compose down -v --rmi all

# Run command defined by ARGS variable
run:
	${RUN} bash -c "${ARGS}"

#
## Go

# go build
compile: dep
	mkdir -p dist
	go build -o dist/${PROJECT_NAME}

# go dep
dep:
	govendor sync
depinfo:
	govendor list && govendor status

# go testing
test:
	go test ${PKG}
testv:
	go test -v ${PKG}
testx: pre check depinfo testv

# go special tests: race/cover/benchmark
race:
	go test -race
cover:
	mkdir -p reports
	echo "mode: count" > reports/coverage-all.out
	$(foreach pkg,$(PKG_ML), \
		go test -coverprofile=reports/coverage.out -covermode=count $(pkg); \
		tail -n +2 reports/coverage.out >> reports/coverage-all.out; \
	)
	go tool cover -html=reports/coverage-all.out -o reports/coverage.html
benchmark:
	go test -run=XXX -benchtime=5s -bench=. 2>/dev/null

# go linting
pre:
	goimports -w ${GOFILES} && gofmt -l -s -w ${GOFILES}
check:
	$(foreach pkg,$(PKG_ML), \
		echo $(pkg) && go vet $(pkg) && golint $(pkg); \
	)
