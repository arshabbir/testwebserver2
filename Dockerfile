
FROM golang
#FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
        GOOS=linux \
        DBHOST=192.168.0.1 \
		LOGPATH=./logs \
	    GOARCH=amd64

	    # Move to working directory /build
	    WORKDIR /build

	    # Copy and download dependency using go mod
	    COPY ./src/go.mod .
	    COPY ./src/go.sum .
	    RUN go mod download

	    # Copy the code into the container
	    COPY ./src. .

		COPY ./src/love.gif .

	    # Build the application
	    RUN go build -o main .

	    # Move to /dist directory as the place for resulting binary folder
	    WORKDIR /dist

	    # Copy binary from build to main folder
	    RUN cp /build/main .

		RUN cp /build/love.gif .


	    # Export necessary port
	    EXPOSE 8080

	    # Command to run when starting the container
	    CMD ["/dist/main"]
