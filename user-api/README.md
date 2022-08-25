

API

```shell
cd user-api

生成api相关文件
选中 user.api，右键 New -  Go Zero - API Quick Code Generattion 

goctl docker -go user.go

goctl kube deploy --name user-api --namespace go-zero-looklook -image user-api:v1.0 -o user.yaml -port 1001 -nodePort 31001

```

RPC

```shell
cd user-rpc

生成rpc相关文件
选中 user.proto，右键 New - Go Zero - ZGPC Quick Code Generattion 

```

DOC

```shell
cd user-api

# 查看一下这个命令怎么使用的
goctl api --help 

# 生成文档

goctl api doc -dir  ./

本级目录下生成 user.md 文档文件


```


Model


先创建数据库 looklook_order

然后创建表 [looklook_order.sql](https://github.com/Mikaelemmmm/go-zero-looklook/blob/main/deploy/sql/looklook_order.sql)

cd 项目根目录

sh genModel/genModel.sh order homestay_order








