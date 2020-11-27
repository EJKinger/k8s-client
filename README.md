# k8s-client

`GOOS=linux go build -o ./app .`
`docker build . -t 839693007267.dkr.ecr.us-east-2.amazonaws.com/ejkinger:1.0.11`
`docker push 839693007267.dkr.ecr.us-east-2.amazonaws.com/ejkinger:1.0.11`
`kubectl run --rm -i demo --image=839693007267.dkr.ecr.us-east-2.amazonaws.com/ejkinger:1.0.11`
