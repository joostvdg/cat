/**
 * ORIGINAL EXAMPLE FROM HERE: https://github.com/jenkinsci/kubernetes-plugin/blob/master/examples/kaniko-declarative.groovy
 * This pipeline will build and deploy a Docker image with Kaniko
 * https://github.com/GoogleContainerTools/kaniko
 * without needing a Docker host
 *
 * You need to create a jenkins-docker-cfg secret with your docker config
 * as described in
 * https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/#create-a-secret-in-the-cluster-that-holds-your-authorization-token
 *
 * ie.
 * kubectl create secret docker-registry regcred --docker-server=index.docker.io --docker-username=csanchez --docker-password=mypassword --docker-email=john@doe.com
 */

pipeline {
    agent {
        kubernetes {
            //cloud 'kubernetes'
            label 'kaniko'
            yaml """
kind: Pod
metadata:
  name: kaniko
spec:
  containers:
  - name: golang
    image: golang:1.11
    command:
    - cat
    tty: true
  - name: kaniko
    image: gcr.io/kaniko-project/executor:debug
    imagePullPolicy: Always
    command:
    - /busybox/cat
    tty: true
    volumeMounts:
      - name: jenkins-docker-cfg
        mountPath: /root
  volumes:
  - name: jenkins-docker-cfg
    projected:
      sources:
      - secret:
          name: regcred
          items:
            - key: .dockerconfigjson
              path: .docker/config.json
"""
        }
    }
    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/joostvdg/cat.git'
            }
        }
        stage('GoLinter') {
            steps {
                container('golang') {
                    // TODO: install linter
                    sh 'curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $GOPATH/bin v1.10.2'
                    // execute linter
                    sh 'golangci-lint --version'
                    sh 'go get -v'
                    sh 'golangci-lint run --issues-exit-code=0'
                }

            }
        }
        stage('Build') {
            steps {
                container('golang') {
                    sh './build-go-bin.sh'
                }
            }
        }
        stage('Make Image') {
            environment {
                PATH = "/busybox:$PATH"
            }
            steps {
                container(name: 'kaniko', shell: '/busybox/sh') {
                    sh '''#!/busybox/sh
                    /kaniko/executor -f `pwd`/Dockerfile.run -c `pwd` --cache=true --destination=index.docker.io/caladreas/cat
                    '''
                }
            }
        }
    }
}
