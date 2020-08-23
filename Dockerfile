FROM golang:1.15

RUN mkdir /i-go-go
# Set the Current Working Directory inside the container
WORKDIR /i-go-go

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

RUN ls -la ./
RUN ls -la ./application_layer

RUN go build application_layer/main.go

# This container exposes port 8080 to the outside world
EXPOSE 3000

# Run the executable
CMD ["./main"]