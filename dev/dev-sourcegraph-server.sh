#!/usr/bin/env bash

cd $(dirname "${BASH_SOURCE[0]}")/..
set -ex

# Build a Sourcegraph server docker image to run for development purposes. Note
# that this image is not exactly identical to the published sourcegraph/server
# images, as those include Sourcegraph Enterprise features.
time cmd/server/pre-build.sh
IMAGE=sourcegraph/server-oss:3.15.1 VERSION=3.15.1 time cmd/server/build.sh

IMAGE=sourcegraph/server-oss:3.15.1 dev/run-server-image.sh
