#!/usr/bin/env bash

set -ex
set -ev

echo $TRAVIS_PULL_REQUEST

NPROCS=$(getconf _NPROCESSORS_ONLN)

echo "Found $NPROCS processors"
date

# Short-circuit transient 'auto-initialization' builds
git fetch origin master
MASTER=$(git describe --always FETCH_HEAD)
HEAD=$(git describe --always HEAD)
echo $MASTER
echo $HEAD
if [ $HEAD == $MASTER ]
then
    echo "HEAD SHA1 equals master; probably just establishing merge, exiting build early"
    exit 0
fi

