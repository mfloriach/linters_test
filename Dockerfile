FROM golang:1.20.1

WORKDIR /usr/src/app

# Hot reload
RUN go install github.com/cosmtrek/air@latest

# acrhitecture compliance
RUN go install -v github.com/fdaines/arch-go@latest

# Coupling metrics
RUN go install github.com/fdaines/spm-go@latest

# golangci-lint
RUN cd /tmp &&\
    git clone https://github.com/golangci/golangci-lint.git &&\
    cd golangci-lint/ &&\ 
    CGO_ENABLED=1 make build &&\
    mv golangci-lint /usr/local/go/bin

COPY . .

# graphical packages dependency 
RUN cd /tmp &&\
    git clone https://github.com/ofabry/go-callvis.git &&\
    cd go-callvis &&\ 
    rm go.sum go.mod &&\ 
    cp /usr/src/app/callvis.mod ./go.mod &&\
    go mod tidy &&\
    go install .

RUN go mod tidy

RUN go build -buildmode=plugin -o ./plugins ./plugins/arch.go