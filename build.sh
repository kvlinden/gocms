#!/bin/bash

GOCMS_VERSION=0.0.$TRAVIS_BUILD_NUMBER
GOCMS_VER_FILE="version.txt"
# for testing. comment out for release.
#TRAVIS_BRANCH=

# grab govendor and sync
echo install govendor
go get -u github.com/kardianos/govendor

echo pulling in deps with govendor
govendor sync

function buildArch() {
    GOOS=$1 GOARCH=$2 go build -o bin/$TRAVIS_BRANCH/$3/$4
    cp -r content bin/$TRAVIS_BRANCH/$3/.
    cp .env_default bin/$TRAVIS_BRANCH/$3/.env
        echo $GOCMS_VERSION > bin/$TRAVIS_BRANCH/$3/$GOCMS_VER_FILE
    pushd bin/$TRAVIS_BRANCH/$3
    rm -rf content/gocms/src
    zip -r gocms.zip * .env
    rm -rf $4 content .env
    popd
}

# build linux64
buildArch linux amd64 linux_64 gocms
buildArch linux 386 linux_32 gocms
buildArch linux arm linux_arm gocms
buildArch darwin amd64 osx_64 gocms
buildArch windows amd64 windows_64 gocms.exe
buildArch windows 386 windows_32 gocms.exe
