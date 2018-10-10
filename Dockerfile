FROM golang:1.11-stretch

ENV PROJECT_BUILD="0.0.1"
ENV PROJECT_NAME="duck-tracker"
ENV PROJECT_PATH=/go/src/github.com/bmanth60/DuckTracker

# Load dependencies
RUN go get -u github.com/kardianos/govendor \
    && go get golang.org/x/tools/cmd/goimports \
    && go get -u github.com/golang/lint/golint

COPY . ${PROJECT_PATH}

WORKDIR ${PROJECT_PATH}

# Create credentials file and sync
RUN make compile

CMD ["/bin/sh", "-c", "${PROJECT_PATH}/${PROJECT_NAME}"]
