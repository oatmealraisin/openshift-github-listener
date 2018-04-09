FROM golang:latest

WORKDIR /go/src/github.com/oatmealraisin/openshift-github-listener
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["openshift-github-listener"]
