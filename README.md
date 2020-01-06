# FUN Cloud API Service by Golang
## Overview (概要)
FUN Cloud PlatformのAPIサーバです．

---
## Description (説明)
### Container
FUN Cloud Platformで管理するKubernetes Clusterを操作するAPIです．
#### Endpoints
|Method|URL|Description|
|---|---|---|
|GET|/v1/containers|get my containers|
|POST|/v1/containers|create container|
|DELETE|/v1/containers/:container_id|delete container|
### Auth
#### Endpoints
|Method|URL|Description|
|---|---|---|
|GET|/mock/token/validate|Header: {"Authorization":"random words"}|
---
## Requirement (要件)
### Installation (設置)
#### Tool (ツール)
- [docker, docker-compose]
#### KUBECONFIGの設定
- $KUBECONFIGに，パスを指定して下さい
- ファイルが複数有る場合は，直接パスを指定できます．
---
### Deploy (起動)
#### Container
##### 開発用
- 環境: コンテナ内に，k8sを構築し，APIはrealizeでHotReloadで開発できます．
- コマンド: `make develop-container`
- MongoExpressに接続し，containerというDBの中に，manifestというコレクションを作成して下さい．
#### Auth
##### 開発用
- コマンド: `make up-auth`