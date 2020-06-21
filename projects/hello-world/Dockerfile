# Build
FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o hello-world && go test

# Immutable Artifact
FROM scratch
WORKDIR /app
COPY --from=build-env /src/hello-world /go/bin/hello
ENTRYPOINT ["/go/bin/hello"]
