# Build
FROM golang:alpine AS build-env
ADD ./src /src
RUN cd /src && go build -o hello-world && go test

# Immutable Artifact
FROM alpine
WORKDIR /app
COPY --from=build-env /src/hello-world /app/
ENTRYPOINT ./hello-world
