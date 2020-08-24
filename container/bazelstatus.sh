#!/bin/sh
set -e

echo                                                                \
  BUILD_SCM_HEAD                                                    \
  $(TZ=etc/UTC                                                      \
      git show -s --format=%cd-%h --date='format:%Y%m%dT%H%M' HEAD)
