```zsh
$ mkdir go-kubernetes-bootcamp
$ go mod init github.com/d5avard/go-kubernetes-bootcamp
$ touch main.go
$ touch readme.md
$ code .
$ go build -o main .
$ docker build -t go-kubernetes-bootcamp .
$ docker run -d -p 8080:8080 go-kubernetes-bootcamp
$ docker container ls
$ docker container logs [docker-id]
$ docker container stop [docker-id]
```

Create config map
```zsh
$ kubectl create -f ./k8s/go-kubernetes-bootcamp-cm.yaml
$ kubectl describe cm go-kubernetes-bootcamp
```

Create a deployment
```
$ kubectl create -f ./k8s/go-kubernetes-bootcamp-dp.yaml
$ kubectl get deployment
$ kubectl get rs
$ kubectl get rs/[rs-id]
$ kubectl describe rs/[rs-id]
$ kubectl get pods
```

```main.go
```

Références:
https://www.callicoder.com/docker-golang-image-container-example/
https://golangbyexample.com/basic-http-server-go/