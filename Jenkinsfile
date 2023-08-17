pipeline{
    agent any
    environment {
        // Necessary to enable Docker buildkit features such as --ssh
        DOCKER_BUILDKIT = "1"
        ADMINX_IMAGE_NAME = "woocoos/adminx"
        FILES_IMAGE_NAME = "woocoos/files"
        AUTH_IMAGE_NAME = "woocoos/auth"
        VERSION = "v0.0.1"
    }
    stages{
        stage("build") {
            steps {
                script {
                    env.GitCommitID = sh (script: 'git rev-parse --short HEAD', returnStdout: true).trim()
                    try{
                        env.GitTag = sh (script: "git describe --tags ${env.GitCommitID}", returnStdout: true).trim()
                    } catch(exception) {
                        echo 'there is no tag in repo'
                    }
                }
            }
        }
        stage("adminx-build") {
            steps {
                script {
                    def tagName = ""
                    docker.withRegistry("http://${REGISTRY_SERVER}") {
                        if (env.GitTag) {
                            tagName = env.GitTag
                        } else {
                            tagName = "${VERSION}.${env.GitCommitID}"
                        }

                        def image = docker.build("${ADMINX_IMAGE_NAME}:${tagName}","--add-host ${NEXUS_HOST} -f cmd/adminx/Dockerfile .")
                        image.push()
                    }
                    if (tagName) {
                        sh "docker rmi ${REGISTRY_SERVER}/${ADMINX_IMAGE_NAME}:${tagName}"
                    }
                }
            }
        }
        stage("files-build") {
            steps {
                script {
                    def tagName = ""
                    docker.withRegistry("http://${REGISTRY_SERVER}") {
                        if (env.GitTag) {
                            tagName = env.GitTag
                        } else {
                            tagName = "${VERSION}.${env.GitCommitID}"
                        }

                        def image = docker.build("${FILES_IMAGE_NAME}:${tagName}","--add-host ${NEXUS_HOST} -f cmd/files/Dockerfile .")
                        image.push()
                    }
                    if (tagName) {
                        sh "docker rmi ${REGISTRY_SERVER}/${FILES_IMAGE_NAME}:${tagName}"
                    }
                }
            }
        }
        stage("auth-build") {
                    steps {
                        script {
                            def tagName = ""
                            docker.withRegistry("http://${REGISTRY_SERVER}") {
                                if (env.GitTag) {
                                    tagName = env.GitTag
                                } else {
                                    tagName = "${VERSION}.${env.GitCommitID}"
                                }

                                def image = docker.build("${AUTH_IMAGE_NAME}:${tagName}","--add-host ${NEXUS_HOST} -f cmd/auth/Dockerfile .")
                                image.push()
                            }
                            if (tagName) {
                                sh "docker rmi ${REGISTRY_SERVER}/${AUTH_IMAGE_NAME}:${tagName}"
                            }
                        }
                    }
                }
    }
}