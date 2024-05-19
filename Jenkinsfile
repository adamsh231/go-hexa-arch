/* groovylint-disable DuplicateMapLiteral, GStringExpressionWithinString, LineLength, NestedBlockDepth, VariableTypeRequired */
pipeline {
    environment {
        credentialId = 'aliregistry'
        url = 'https://registry-intl.ap-southeast-5.aliyuncs.com'
        scannerHome = tool 'Sonarqube'
        serviceDefault = 'go-hexa'
        serviceHttp = 'go-hexa-http'
        serviceConsumer = 'go-hexa-consumer'
        GO111MODULE = 'on'
    }

    agent { node { label 'agent1' } }

    tools {
        go 'go-1.21.4'
    }

    stages {

        stage('Unit Testing') {
            steps {
                script {
                        sh 'go test ./... -coverprofile=coverage.out'
                }
            }
        }

        stage('Sonar Scanner With PR') {
            when {
                branch 'PR-*'
            }
            steps {
                retry(3) {
                    script {
                        withSonarQubeEnv('Sonarqube') {
                            def prKey = "-Dsonar.pullrequest.key=${env.CHANGE_ID}"
                            def prBranch = "-Dsonar.pullrequest.branch=${env.CHANGE_BRANCH}"
                            def prBase = "-Dsonar.pullrequest.base=${env.CHANGE_TARGET}"
                            // Run the scan
                            sh "${scannerHome}/bin/sonar-scanner ${prKey} ${prBranch} ${prBase}"
                        }
                        timeout(time: 10, unit: 'MINUTES') {
                            waitForQualityGate abortPipeline: true
                        }
                    }
                }
            }
            post {
                success {
                    slackSend(color: '#008000', message: "PASSED: Sonarqube PR ${env.JOB_NAME} #${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)")
                }
                failure {
                    slackSend(color: '#FF0000', message: "FAILED: Sonarqube PR ${env.JOB_NAME} #${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)")
                }
            }
        }

        stage('Publish Approval') {
            when {
                tag 'v*'
            }
            steps {
                script {
                    //   input message: "Deploy these changes?", submitter "admin"
                    def userName = input message: 'Deploy these changes?', submitter: 'admin,adam', submitterParameter: 'admin,adam'
                    echo "Accepted by ${userName}"
                    if (!(['admin', 'adam'].contains(userName))) {
                        error('This user is not approved to deploy to PROD.')
                    }
                }
            }
        }

        stage('Sonar Scanner Tag Release Prod') {
            when {
                tag 'v*'
            }
            steps {
                script {
                    withSonarQubeEnv('Sonarqube') {
                            sh "${scannerHome}/bin/sonar-scanner -Dsonar.branch.name=${env.TAG_NAME}"
                    }
                    timeout(time: 10, unit: 'MINUTES') {
                        waitForQualityGate abortPipeline: true
                    }
                }
            }
            post {
                success {
                    slackSend(color: '#008000', message: "PASSED: Sonarqube Prod ${env.JOB_NAME} #${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)")
                }
                failure {
                    slackSend(color: '#FF0000', message: "FAILED: Sonarqube Prod ${env.JOB_NAME} #${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)")
                }
            }
        }

        stage('Build Image Tag Release Prod') {
            when {
                tag 'v*'
            }
            steps {
                slackSend(color: '#FFFF00', message: "STARTED: Build Image Prod ${env.JOB_NAME} #${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)")
                echo 'Build Image Prod'
                sh 'docker build . -t ${serviceDefault}:${TAG_NAME}-${BUILD_NUMBER} --build-arg SSH_PRIVATE_KEY="$(cat /var/lib/jenkins/id_rsa)"'
                sh 'echo ini build image prod'
            }
            post {
                success {
                    slackSend(color: '#008000', message: "SUCCESS: Build Image Prod ${env.JOB_NAME} #${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)")
                }
                failure {
                    slackSend(color: '#FF0000', message: "FAILED: Build Image Prod ${env.JOB_NAME} #${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)")
                }
            }
        }

        stage('Docker Login tag push Tag Release Prod') {
            when {
                tag 'v*'
            }
            steps {
                script {
                    echo 'Push docker image ke docker registry Ali Prod'
                    docker.withRegistry(url, credentialId) {
                        sh 'docker tag ${serviceDefault}:${TAG_NAME}-${BUILD_NUMBER} registry-intl.ap-southeast-5.aliyuncs.com/adam/${serviceDefault}:${TAG_NAME}-${BUILD_NUMBER}'
                        sh 'docker push registry-intl.ap-southeast-5.aliyuncs.com/adam/${serviceDefault}:${TAG_NAME}-${BUILD_NUMBER}'
                        sh 'echo ini docker login tag push prod'
                    }
                }
            }
        }

        stage('Set Image Kubernetes Tag Release Prod') {
            when {
                tag 'v*'
            }
            steps {
                script {
                    sh 'kubectl --kubeconfig="../../kubeconfig-prod-premium-noncore.yaml" set image deployment ${serviceHttp} ${serviceHttp}=registry-intl.ap-southeast-5.aliyuncs.com/adam/${serviceDefault}:${TAG_NAME}-${BUILD_NUMBER} -n=non-core'
                    sh 'kubectl --kubeconfig="../../kubeconfig-prod-premium-noncore.yaml" set image deployment ${serviceConsumer} ${serviceConsumer}=registry-intl.ap-southeast-5.aliyuncs.com/adam/${serviceDefault}:${TAG_NAME}-${BUILD_NUMBER} -n=non-core'
                    sh 'echo set image k8s prod'
                }
            }
            post {
                success {
                    slackSend(color: '#008000', message: "SUCCESS: Deployment Prod ${env.JOB_NAME} #${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)")
                }
                failure {
                    slackSend(color: '#FF0000', message: "FAILED: Deployment Prod ${env.JOB_NAME} #${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)")
                }
            }
        }

        stage('Sonar Scanner Branch Master') {
            when {
                branch 'master'
            }
            steps {
                retry(3) {
                    script {
                        withSonarQubeEnv('Sonarqube') {
                            sh "${scannerHome}/bin/sonar-scanner -Dsonar.branch.name=${env.BRANCH_NAME}"
                        }
                        timeout(time: 10, unit: 'MINUTES') {
                            waitForQualityGate abortPipeline: true
                        }
                    }
                }
            }
            post {
                success {
                    slackSend(color: '#008000', message: "PASSED: Sonarqube Branch Master ${env.BRANCH_NAME} #${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)")
                }
                failure {
                    slackSend(color: '#FF0000', message: "FAILED: Sonarqube Branch Master ${env.BRANCH_NAME} #${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)")
                }
            }
        }

        stage('Build Image Beta Branch Master') {
            when {
                branch 'master'
            }
            steps {
                slackSend(color: '#FFFF00', message: "STARTED: Build Image Beta ${env.JOB_NAME} #${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)")
                echo 'Build Image beta'
                sh 'docker build . -t ${serviceDefault}:${BRANCH_NAME}-${BUILD_NUMBER} --build-arg SSH_PRIVATE_KEY="$(cat /var/lib/jenkins/id_rsa)"'
                sh 'echo ini build image beta'
            }
            post {
                success {
                    slackSend(color: '#008000', message: "SUCCESS: Build Image Beta ${env.JOB_NAME} #${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)")
                }
                failure {
                    slackSend(color: '#FF0000', message: "FAILED: Build Image Beta ${env.JOB_NAME} #${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)")
                }
            }
        }

        stage('Docker Login tag push Beta Branch Master') {
            when {
                branch 'master'
            }
            steps {
                script {
                    echo 'Push docker image ke docker registry Ali Beta'
                    docker.withRegistry(url, credentialId) {
                        sh 'docker tag ${serviceDefault}:${BRANCH_NAME}-${BUILD_NUMBER} registry-intl.ap-southeast-5.aliyuncs.com/adam/${serviceDefault}:${BRANCH_NAME}-${BUILD_NUMBER}'
                        sh 'docker push registry-intl.ap-southeast-5.aliyuncs.com/adam/${serviceDefault}:${BRANCH_NAME}-${BUILD_NUMBER}'
                        sh 'echo ini docker login tag push Beta'
                    }
                }
            }
        }

        stage('Set Image Kubernetes Beta Branch Master') {
            when {
                branch 'master'
            }
            steps {
                script {
                    sh 'kubectl --kubeconfig="../../kubeconfig-beta-premium.yaml" set image deployment ${serviceHttp} ${serviceHttp}=registry-intl.ap-southeast-5.aliyuncs.com/adam/${serviceDefault}:${BRANCH_NAME}-${BUILD_NUMBER} -n=beta'
                    sh 'kubectl --kubeconfig="../../kubeconfig-beta-premium.yaml" set image deployment ${serviceConsumer} ${serviceConsumer}=registry-intl.ap-southeast-5.aliyuncs.com/adam/${serviceDefault}:${BRANCH_NAME}-${BUILD_NUMBER} -n=beta'
                    sh 'echo set image k8s Beta'
                }
            }
            post {
                success {
                    slackSend(color: '#008000', message: "SUCCESS: Deployment Beta ${env.JOB_NAME} #${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)")
                }
                failure {
                    slackSend(color: '#FF0000', message: "FAILED: Deployment Beta ${env.JOB_NAME} #${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)")
                }
            }
        }
    }
    post {
        success {
            slackSend(color: '#008000', message: "SUCCESS: Pipeline ${env.JOB_NAME} #${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)")
        }
        failure {
            slackSend(color: '#FF0000', message: "FAILED: Pipeline ${env.JOB_NAME} #${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)")
        }
    }
}