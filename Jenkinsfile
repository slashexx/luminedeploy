
pipeline {
    agent { label 'linux' }
    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        stage('Build') {
            steps {
                sh 'go buils'
            }
        }
        stage('Test') {
            steps {
                sh 'go test ./'
            }
        }
    }
    post {
        always {
            echo "Pipeline 'test' completed for branch 'main'"
        }
    }
}
