load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_repositories():
    go_repository(
        name = "com_github_arran4_golang_ical",
        importpath = "github.com/arran4/golang-ical",
        sum = "h1:Yvu285J3YTE3cCCE5rgUGB7sP988cYolf81eP2b3r3U=",
        version = "v0.0.0-20200314043952-7d4940584ecd",
    )
