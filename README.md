# opsybits

## what is this?

`opsybits` is a collection of small, single-task, cloud-oriented applets in a
single executable. Similar style to `busybox`, but without the hardlinks.

## what applets exist?

### hi-m8 - static message webserver 

`hi-m8` hosts a single, static message on an HTTP endpoint. Optionally, an
artifical delay can be incurred in each HTTP response.

In addition to serving a static message, this applet also serves

* a `/healthz` endpoint, for the usual healthcheck purposes
* a `/metrics` endpoint, for Prometheus integration

### logspew - JSON log generator

`logspew` generates lorem ipsum JSON log data and is intended to assist with
validating ingestion of logs into systems like Graylog or Elasticsearch. Its
capabilities include:

* optionally add emoji to JSON field content (testing UTF8 handling)
* optionally add emoji to JSON field names (testing UTF8 handling)
* specify a quantity of lorem ipsum fields (testing for size limits)
* specify an upper limit on the volume of lorem ipsum per field (testing
  handling of long log lines)

## get a list of applets?
```
$ opsybits --help
```

## get help with a specific applet?

```
$ opsybits APPLET --help
```

## Docker

```
$ docker run gcr.io/jsleeio-containers/opsybits:v0.1.0
```
