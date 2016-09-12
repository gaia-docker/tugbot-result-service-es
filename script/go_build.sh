#!/bin/sh
distdir=.dist

go_build() {
  rm -rf "${distdir}"
  mkdir "${distdir}"
  glide install
  CGO_ENABLED=0 go build -v -o ${distdir}/tugbot-result-service-es
}

go_build
