# hat_push

该项目是一个游戏功能推送服务，主要用于实现游戏内功能的推送和管理。以下是项目的详细说明和使用指南。

## 项目简介

`hat_push` 是一个基于 Go 语言开发的游戏功能推送服务，旨在为游戏提供高效、稳定的功能推送和管理功能。通过该服务，可以实现游戏内各种功能的快速推送和更新，提升游戏的用户体验和运营效率。

## 本地启动参数配置

要本地启动 `hat_push` 服务，需要配置以下参数：

```shell
slot_dev hat_push 5 debug
```

参数说明：

- `slot_dev`：集群名称
- `hat_push`：服务名
- `5`：logicId 节点
- `debug`：调试模式

## Logic API 对外 RPC 接口生成

使用以下命令生成 `logic` API 的对外 RPC 接口：

```shell
protoc --go_out=plugins=kite:../slot_logic/ proto/api.proto
```

## 贡献

如果您有任何改进建议或想要贡献代码，请随时提交 Pull Request 或者在 Issues 中提出。

## 联系方式

如果您在使用过程中遇到任何问题，可以通过以下方式联系开发者：

- GitHub Issues：https://github.com/EddieChan1993/hat_push/issues

感谢您对 `hat_push` 项目的关注和支持！