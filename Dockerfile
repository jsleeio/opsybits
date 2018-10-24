# Start by building the application.
FROM golang:1.11 as build

WORKDIR /go/src/github.com/jsleeio/opsybits
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build

# Now copy it into our base image.
FROM scratch
USER 1000
COPY --from=build /go/src/github.com/jsleeio/opsybits/opsybits /opsybits
ENTRYPOINT ["/opsybits"]
