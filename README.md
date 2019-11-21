# cloud container management service (ccms)
## Overview

## Description
### Endpoints
|Method|URL|Description|
|---|---|---|
|GET|/v1/containers|get my containers|
|POST|/v1/containers|create container|
|DELETE|/v1/containers/:container_id|delete container|
## Requirement
### Tool
- Docker
- Docker-Compose
### put your kubeconfig into ccms/file directory
- `copy $HOME/.kube/config project_root/file/kubeconfig`
## Installation
### Production
- use dockerfile
### Development
- build dockerfile: `make`
- run containers: `make run`