############################
# STEP 1 prepare the source
############################
FROM golang:alpine AS source

# Create the user and group files that will be used in the running container to
# run the process as an unprivileged user.
RUN mkdir /user && \
  echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
  echo 'nobody:x:65534:' > /user/group

# Set the environment variables for the go command:
# * CGO_ENABLED=1 to build a non-statically-linked executable
# * GO111MODULE=on Force the go compiler to use modules
# * GOOS=linux to run on linuxos
# * GOARCH=to run on amd64 architecture
ENV CGO_ENABLED=0 GO111MODULE=on GOOS=linux

# Install the Certificate-Authority certificates for the app to be able to make
# calls to HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata
# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Import go mod from the context.
COPY go.mod .

#This is the ‘magic’ step that will download all the dependencies that are specified in 
# the go.mod and go.sum file.
# Because of how the layer caching system works in Docker, the  go mod download 
# command will _ only_ be re-run when the go.mod or go.sum file change 
# (or when we add another docker instruction this line)
RUN go mod download

############################
# STEP 2 build the executable
############################
FROM source AS builder

# Import the code from the context.
COPY . /src

# And compile the project
# RUN go install -a -tags netgo -ldflags '-w -extldflags "-static"' /app
RUN go build \
  -a \
  -installsuffix 'static' \
  -o /app /src/cmd/service

# ############################
# # STEP 3 the running container
# ############################
FROM alpine:latest AS final

ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip
ENV ZONEINFO /zoneinfo.zip

# Copy single timezone
COPY --from=builder /usr/share/zoneinfo/Asia/Bangkok /etc/localtime

# Import the user and group files from the first stage.
COPY --from=builder /user/group /user/passwd /etc/

# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy our static executable.
COPY --from=builder /app .

# Import the code from the context.
COPY ./cmd/service/config.yaml .

# Perform any further action as an unprivileged user.
USER nobody:nobody

# Run the binary.
ENTRYPOINT ["/app"]
