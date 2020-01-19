#!/usr/bin/env groovy

node(label 'docker && sonar') {
  String applicationName = "cfdns"
  String buildNumber = "${env.GIT_BRANCH}.${env.BUILD_NUMBER}"
  String goPath = "/go/src/github.com/someone-stole-my-name/${applicationName}"

  stage('Checkout from GitHub') {
    checkout scm
  }

  stage('Sonarqube') {
    environment {
      scannerHome = tool 'SonarQubeScanner'
    }
    steps {
      withSonarQubeEnv('sonar') {
          sh "${scannerHome}/bin/sonar-scanner"
      }
      timeout(time: 10, unit: 'MINUTES') {
        waitForQualityGate abortPipeline: true
      }
    }
}

  stage("Create binaries") {
    docker.image("golang:1.13.6-alpine3.11").inside("-u root -v ${pwd()}:${goPath}") {
        sh "apk add dep git && cd ${goPath} && dep ensure"
        sh "cd ${goPath} &&  GOOS=darwin GOARCH=amd64 go build -o binaries/amd64/${buildNumber}/darwin/${applicationName}-${buildNumber}.darwin.amd64"
        sh "cd ${goPath} && GOOS=windows GOARCH=amd64 go build -o binaries/amd64/${buildNumber}/windows/${applicationName}-${buildNumber}.windows.amd64.exe"
        sh "cd ${goPath} && GOOS=linux GOARCH=amd64 go build -o binaries/amd64/${buildNumber}/linux/${applicationName}-${buildNumber}.linux.amd64"
    }
  }

//   stage("Archive artifacts") {
//     // Archive the binary files in Jenkins so we can retrieve them later should we need to audit them
//     archiveArtifacts artifacts: 'binaries/**', fingerprint: true
//   }
}