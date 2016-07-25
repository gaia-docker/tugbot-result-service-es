#!/bin/sh
distdir=.dist

go_build() {
  rm -rf "${distdir}"
  mkdir "${distdir}"
  glide install
  go build -v -o ${distdir}/tugbot-result-service-es
}

go_build