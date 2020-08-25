The `cmscal` program generates an iCalendar version of the Clark Middle
School block schedule; it's written in the [Go programming
language](https://golang.org/).

Usage
---

`cmscal` serves HTTP on the port specified by the `PORT` environment
variable, because that's the convention used by the [Google Cloud Run
service](https://cloud.google.com/run). The default is 8080.

Building
---

To build, you only need to install the [Bazel](https://bazel.build/)
build system and run `bazel build`:

```console
$ bazel build -c opt //cmd:cmscal
INFO: Build option --compilation_mode has changed, discarding analysis cache.
INFO: Analyzed target //cmd:cmscal (0 packages loaded, 6935 targets configured).
INFO: Found 1 target...
Target //cmd:cmscal up-to-date:
  bazel-out/k8-opt-ST-5e46445d989a/bin/cmd/cmscal_/cmscal
INFO: Elapsed time: 9.828s, Critical Path: 9.36s
INFO: 5 processes: 5 linux-sandbox.
INFO: Build completed successfully, 8 total actions
```
Container builds are supported. You can build a container tar based on
[Alpine Linux](https://alpinelinux.org/):

```console
$ bazel build -c opt //container:cmscal
```

If you have Docker installed, you can load that tar into your machine's
Docker daemon by using `bazel run` instead of `bazel build`:

```console
$ bazel run -c opt //container:cmscal
$ docker run -d -p 8080 --rm bazel/container:cmscal
```
