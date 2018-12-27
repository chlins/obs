# obs

A tiny poor objects store server.

## Architecture

http_request -> apiserver(support multi) -> dataserver(suport multi)

## Usage

1. `build binary(see build/Makefile)`

2. `STORAGE_ROOT='/tmp' ZOOK_SERVER='127.0.0.1:2181' ./dataserver && ./apiserver`

3. `curl -v -X PUT -d "test content" 127.0.0.1:6868/objects/test`

## Dependency

- [Apache zookeeper](https://zookeeper.apache.org/)

## Progress

- [x] put
- [ ] get
- [ ] meta info

...

## LISCENCE

MIT