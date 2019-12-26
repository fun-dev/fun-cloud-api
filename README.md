# FUN Cloud API Service by Golang
## Overview (概要)

---
## Description (説明)
### Endpoints
#### Container
|Method|URL|Description|
|---|---|---|
|GET|/v1/containers|get my containers|
|POST|/v1/containers|create container|
|DELETE|/v1/containers/:container_id|delete container|
#### Auth
|Method|URL|Description|
|---|---|---|
|GET|/mock/token/validate|Header: {"Authorization":"random words"}|
---
## Requirement (要件)
### Installation (設置)
#### Tool (ツール)
- docker
- docker-compose
#### put your kubeconfig TODO: fix kubeconfig path
- `copy $HOME/.kube/config project_root/file/kubeconfig`
---
### Activate (起動)
#### Auth Server
##### Specification (仕様)
- Host: `127.0.0.1:3000`
##### Command
- `make up-auth`