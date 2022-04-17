# go-web-hello-world
hello world sample for docker and kube that created by golang

## prerequistes
Install go-1.18.1 on Ubuntu vm

## Create go project and the Dockerfile
1. create the go-web-hello-world src file
2. run the following cmd:
```
go mod init go-web-hello-world
go mod tidy
```
3. create Dockerfile

## Build docker image
`docker build -t elainwu665/go-web-hello-world:v0.1 .`

## Run the docker image
`docker run -d -p 28082:8081 --name go-test  elainewu665/go-web-hello-world:v0.1`

## Push the image to dockerhub
```
docker login
docker push elainewu665/go-web-hello-world:v0.1
```

## Install kubeadm

