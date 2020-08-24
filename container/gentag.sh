#!/bin/sh
set -e
cat bazel-out/stable-status.txt bazel-out/volatile-status.txt \
  | grep ^BUILD_SCM_HEAD                                      \
  | awk '{print $2}'
