FROM golang

COPY go.mod .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

WORKDIR /go/app