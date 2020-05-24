# FUN Cloud API Service
## Description (説明)
### List Service（サービス一覧）
1. Auth: 認証系
2. Container: コンテナ操作（Pod）
3. Data: Pod⇄クライアントのデータ転送
4. Directory: Podに対して，lsするやつ
5. Image: Docker Imageを管理するアプリに対する操作
### Detail Service（サービス詳細）
#### Container
- FUN Cloud Platform（FCP: FUN Cloud Platform）で管理するKubernetes Clusterを操作するアプリです．
##### Endpoints（エンドポイント）
| Method |              URL              | Description |
| :----: | :---------------------------: | :---------: |
|  GET   |        /v1/containers         |   RESTFUL   |
|  POST  |        /v1/containers         |   RESTFUL   |
| DELETE | /v1/containers/{container_id} |   RESTFUL   |
---
### Auth
#### Endpoints
| Method |        URL        |     Description      |
| :----: | :---------------: | :------------------: |
|  GET   | /v1/auth/userinfo | Get User Information |
|  POST  |  /v1/auth/singin  |        SignIn        |
|  POST  | /v1/auth/singout  |       SignOut        |
|  POST  |  /v1/auth/signup  |        SignUp        |
---
### Data
#### Endpoints
| Method |             URL              |         Description          |
| :----: | :--------------------------: | :--------------------------: |
|  POST  |  /v1/containers/file/upload  |   file upload to container   |
|  POST  | /v1/containers/file/download | file download from container |
---
### Directory
#### Endpoints
| Method |                URL                |   Description    |
| :----: | :-------------------------------: | :--------------: |
|  GET   | /v1/directory/ls/{directory_path} | get pod dir path |
---
### Image
#### Endpoints
| Method |         URL          | Description |
| :----: | :------------------: | :---------: |
|  POST  | /v1/container_images |   RESTFUL   |
|  GET   | /v1/container_images |   RESTFUL   |
| DELETE | /v1/container_images |   RESTFUL   |

## Requirement (要件)
### Installation (設置)
#### Tool (ツール)
- Docker
- Docker-Compose
#### KUBECONFIG
- 環境変数の設定: `$KUBECONFIG`にパスを指定
---
### Deployment (起動)
