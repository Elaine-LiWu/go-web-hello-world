# go-web-hello-world
hello world sample for docker and kube that created by golang

## prerequistes
Install go-1.18.1 on Ubuntu vm (as I'm the root on this vm, there is no sudo in the following cmd)

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
validate the request by `curl http://127.0.0.1:28082/demo/go-web-hello-world` which reponses Go Web Hello World!

## Push the image to dockerhub
```
docker login
docker push elainewu665/go-web-hello-world:v0.1
```

## Install single node cluster by kubeadm
### install kube cluster 
1. run the following cmd to install kubeadm, kubectl, kubelet
```
apt-get update
apt-get install -y kubelet kubeadm kubectl
apt-mark hold kubelet kubeadm kubectl
```
2. install etcd client by
`apt install etcd-client`
and validate the `etcdctl -v`

3. update docker cgroup driver to systemd in /etc/docker/daemon.json
4. install kube control plane by 
```
kubeadm init --pod-network-cidr=10.208.0.0/16
export KUBECONFIG=/etc/kubernetes/admin.conf

```
5. install Calico network component by 
```
curl https://projectcalico.docs.tigera.io/manifests/calico.yaml -O
kubectl apply -f calico.yaml
```
7. validate cluster status by `kubectl get no` and remove the master taint `kubectl taint nodes --all node-role.kubernetes.io/master-`

### deploy the image to kube cluster

8. modify the Dockerfile to expose 8081 port and push v0.2 image to dockerhub
9. deploy the image by `kubectl run go-hello --image=elainewu665/go-web-hello-world:v0.2 --replicas=1`
10. expose the service by NodePort `kubectl expose deploy go-hello --type=NodePort --port=31080 --target-port=8081` and then validate the service by `curl http://10.98.241.198:31080/demo/go-web-hello-world ` which reponses Go Web Hello World!

### deploy kube dashboard
11. deploy dashboard by `kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.5.0/aio/deploy/recommended.yaml`
12. expose dashboard by `kubectl expose deploy kubernetes-dashboard --name=my-kubernetes-dashboard --type=NodePort --port=31081 --target-port=8443 -n kubernetes-dashboard`
13. get the token by `kubectl describe secret kubernetes-dashboard-token-ns2q6 -n kubernetes-dashboard`
