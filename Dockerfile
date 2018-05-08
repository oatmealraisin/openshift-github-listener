FROM golang:latest

RUN wget https://github.com/openshift/origin/releases/download/v3.9.0/openshift-origin-client-tools-v3.9.0-191fece-linux-64bit.tar.gz
RUN tar -xvzf openshift-origin-client-tools-v3.9.0-191fece-linux-64bit.tar.gz
RUN cp openshift-origin-client-tools-v3.9.0-191fece-linux-64bit/* /usr/local/bin/

WORKDIR /go/src/github.com/oatmealraisin/openshift-github-listener
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["openshift-github-listener"]
