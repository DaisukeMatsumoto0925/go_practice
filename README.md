tag
`docker tag circleci-ecs:v1 823425155155.dkr.ecr.ap-northeast-1.amazonaws.com/index_indicators:latest`

push
`docker push 823425155155.dkr.ecr.ap-northeast-1.amazonaws.com/index_indicators:latest`

task definition
`aws ecs register-task-definition --cli-input-json file://task-definition.json`

service
`aws ecs create-service --cli-input-json file://ecs-service.json`