# Start by building the application.
FROM golang:1.12-alpine3.9 as build
RUN apk add git
RUN adduser -h /build -s /bin/sh -D build
ADD . /build/src
RUN chown -R build /build
USER build
WORKDIR /build/src
ENV GOCACHE=/build/.cache
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly

# Now copy it into our base image.
FROM scratch
COPY --from=build /build/src/opsybits /opsybits
USER 1000
ENTRYPOINT ["/opsybits"]
