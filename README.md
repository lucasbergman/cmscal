The `cmscal` program generates an iCalendar version of the Clark Middle
School block schedule; it's written in the [Go programming
language](https://golang.org/).

Copying and Distributing
---

Copyright 2020 Lucas Bergman

Licensed under the Apache License, Version 2.0 (the "License"); you may not
use this file except in compliance with the License. You may obtain a copy of
the License at

  https://www.apache.org/licenses/LICENSE-2.0

or in the `LICENSE` file in this program's code.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
License for the specific language governing permissions and limitations under
the License.

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
$ bazel build -c opt //cmd/cmscal
Starting local Bazel server and connecting to it...
INFO: Analyzed target //cmd/cmscal:cmscal (38 packages loaded, 6935 targets configured).
INFO: Found 1 target...
Target //cmd/cmscal:cmscal up-to-date:
  bazel-out/k8-opt-ST-5e46445d989a/bin/cmd/cmscal/cmscal_/cmscal
INFO: Elapsed time: 30.361s, Critical Path: 10.07s
INFO: 7 processes: 7 linux-sandbox.
INFO: Build completed successfully, 11 total actions
```

If you already use Go, you can build with Go:

```console
$ go install -i github.com/lucasbergman/cmscal/cmd/cmscal
$ PORT=12345 ~/go/bin/cmscal
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
