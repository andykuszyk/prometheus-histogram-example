FROM golang
WORKDIR /go/src/github.com/andykuszyk/prometheus-histogram-example/application
COPY ./main.go main.go
COPY ./go.mod go.mod
RUN go build main.go
CMD ./main 
