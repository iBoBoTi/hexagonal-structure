pipeline {
    agent any
    tools {
            go 'go1.17'
        }
    environment {
        GO117MODULE = 'on'
        ARTIFACT_ID = 'resource-service'
		SERVICE_PROFILE = 'dev,elk'
		DO_REINSTALL = 1
		BINARY_NAME = 'resource'
    }
    stages {

        stage('Build') {
            steps {
                       echo 'Compiling and building'
                       sh 'go build .'
                  }
        }

        stage('Deploy') {
            steps {
				sh 'sudo slingshot.directories.go ${ARTIFACT_ID}'
				sh 'sudo slingshot.ftp.go ${ARTIFACT_ID} ${BINARY_NAME} ${ARTIFACT_ID}.env'
				sh 'sudo slingshot.systemd.go ${ARTIFACT_ID} ${BINARY_NAME} ${SERVICE_PROFILE} ${DO_REINSTALL}'
            }
        }
    }
}
