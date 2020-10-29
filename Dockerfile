FROM golang:1.14


WORKDIR /go/src/testapp
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
COPY . .
RUN make testapp 

ENTRYPOINT ["/go/src/testapp/testapp"]
EXPOSE 8080
