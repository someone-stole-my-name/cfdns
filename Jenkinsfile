#!/usr/bin/env groovy

def applicationName = "cfdns"
def binaries = [
  [GOOS: "linux", GOARCH: "amd64", SUFFIX: ""],
  [GOOS: "linux", GOARCH: "arm", SUFFIX: ""],
  [GOOS: "windows", GOARCH: "amd64", SUFFIX: ".exe"]
]

node('master') {
  def root = tool name: 'Go 1.14.2', type: 'go'
  stage('Checkout') {
    checkout scm
  }

  stage("Create binaries") {
    commitId = sh(returnStdout: true, script: 'git rev-parse --short HEAD | tr -d "\n\r"')
    build = "${BRANCH_NAME}_${commitId}".replace("/", "-")

    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin:${HOME}/go/bin", "GOPATH=${WORKSPACE}"]) {
      sh 'go get -u github.com/golang/dep/cmd/dep'
      dir("${GOPATH}/src/cfdns") {
        sh "${GOPATH}/bin/dep ensure"
        for (item in binaries) {
          GOOS = item['GOOS']
          GOARCH = item['GOARCH']
          SUFFIX = item['SUFFIX']
          sh "GOOS=${GOOS} GOARCH=${GOARCH} go build -o \
            ${WORKSPACE}/artifacts/${applicationName}_${GOOS}_${GOARCH}_${build}${SUFFIX}"
        }
      }
    }
  }
  stage("Upload Artifacts") {
    step([$class: 'MinioUploader', sourceFile: "artifacts/*", bucketName: "artifacts"])
  }
}
