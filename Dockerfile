FROM golang:1.11-stretch

ENV PROJECT_BUILD="0.0.1"
ENV PROJECT_NAME="duck-tracker"
ENV PROJECT_PATH=/go/src/github.com/bmanth60/DuckTracker

# Load dependencies
RUN go get -u github.com/kardianos/govendor \
    && go get golang.org/x/tools/cmd/goimports \
    && go get -u golang.org/x/lint/golint

COPY . ${PROJECT_PATH}

WORKDIR ${PROJECT_PATH}

RUN make compile

CMD ["make", "serve"]
