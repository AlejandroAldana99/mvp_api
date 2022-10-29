
##################################
# STEP 1 build executable binary #
##################################
# Get image of golang
FROM golang:1.14 as builder

ENV GO111MODULE=on

# Set workdir
WORKDIR /go/src/api-candidate-data

# Copy all from local to image
ADD . .

# Get libraries
RUN go get -t ./...

# Run test
RUN go test ./...

# Craete binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/app


##############################
# STEP 2 build alpine image #
##############################
# image small of linux
FROM scratch

COPY --from=builder /go/bin/app /go/bin/app

##################### Set environment variables #####################

# Garbage collector conf
ENV GOGC=10

# CORS Configuration (comma separated values)
# e.g GET,OPTIONS,POST,DELETE

ENV CORS_ALLOWED_ORIGINS=*
ENV CORS_ALLOWED_HEADERS=X-Requested-With,Content-Type,Authorization
ENV CORS_ALLOWED_METHODS=GET,OPTIONS,POST,DELETE

# Logger configuration

ENV LOG_LEVEL=1
ENV ENABLE_FILE_LOG=true
ENV ENABLE_CONSOLE_LOGS=true

##################################################

# Run the app binary.
CMD ["/go/bin/app"]

EXPOSE 8151

################################################################
