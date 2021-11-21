# Docker環境の使い方

このプロジェクトでは、ローカルマシンにDocker環境を構築して動かすことを前提としています。  
この README.md は、そのDocker環境の起動方法についてのメモです。

## ディレクトリ構成
* `app`: 実際に Go言語 のコードを入れるディレクトリ
* `Docker`: Docker環境構築に必要なファイル一式があるディレクトリ、`docker compose` コマンドもここで実行する

## 起動方法（初回のみ実行が必要）
### `.env` をコピー
```bash
$ cd /path/to/Docker
$ cp .env.develop .env
```
docker-compose は、 .env に記載された COMPOSE_PROJECT_NAME を参照してプロジェクト名を決定する
これが無い場合は、親ディレクトリ名（このプロジェクトならば `Docker`）が使用される

### ビルド
```bash
$ docker compose build --no-cache
[+] Building 3.2s (10/10) FINISHED                                                                                                        
 => [internal] load build definition from Dockerfile                                                                                 0.0s
 => => transferring dockerfile: 32B                                                                                                  0.0s
 => [internal] load .dockerignore                                                                                                    0.0s
 => => transferring context: 2B                                                                                                      0.0s
 => [internal] load metadata for docker.io/library/golang:1.17-alpine                                                                0.8s
 => CACHED [1/5] FROM docker.io/library/golang:1.17-alpine@sha256:55da409cc0fe11df63a7d6962fbefd1321fedc305d9969da636876893e289e2d   0.0s
 => [internal] load build context                                                                                                    0.0s
 => => transferring context: 183B                                                                                                    0.0s
 => [2/5] RUN apk update && apk add git                                                                                              2.0s
 => [3/5] RUN mkdir /go/src/app                                                                                                      0.3s
 => [4/5] WORKDIR /go/src/app                                                                                                        0.0s 
 => [5/5] ADD . /go/src/app                                                                                                          0.0s 
 => exporting to image                                                                                                               0.1s 
 => => exporting layers                                                                                                              0.1s 
 => => writing image sha256:1eeb8b779d4742fc965ce830546bbd0dd6631435c20bcbf825f8ae439940c2e1                                         0.0s 
 => => naming to docker.io/library/play-golang_app                                                                                   0.0s

Use 'docker scan' to run Snyk tests against images to find vulnerabilities and learn how to fix them
```

### イメージの確認

次のようにしているため、 `play-golang_app` という行があればOK
* .env にて `COMPOSE_PROJECT_NAME=play-golang` としている
* docker-compose.yml にて service名を `app` としている

```bash
$ docker images
REPOSITORY                     TAG          IMAGE ID       CREATED          SIZE
play-golang_app                latest       d64127b8edd2   8 seconds ago    328MB
```

## 起動（毎回実行する）
```bash
$ docker-compose up -d
Creating network "play-golang_default" with the default driver
Creating play-golang_app_1 ... done
```

## 実行方法
```bash
$ docker-compose exec app go run main.go
Hello golang from docker!
```

## 参考
このDocker環境は、下記の記事を参考にして構築しました。
* https://zenn.dev/tomi/articles/2020-10-14-go-docker

