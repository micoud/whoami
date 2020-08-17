pipeline {
  agent any

  environment {
    REGISTRY="registry.cta-test.zeuthen.desy.de"
  }

  stages {
    stage('Build container') {
      steps {
        sh "docker build -t ${REGISTRY}/whoami:latest ."
      }
    }

    stage('Push container to registry') {
      steps {
        sh "docker push ${REGISTRY}/whoami:latest"
      }
    }
  }
}
