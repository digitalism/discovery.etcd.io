# discovery.etcd.io

The etcd discovery service, like the one at https://discovery.etcd.io. The API is documented in

https://github.com/coreos/etcd/tree/master/Documentation/cluster-discovery.md
https://github.com/coreos/etcd/tree/master/Documentation/discovery-protocol.md

This repository is forked from https://github.com/mengzhuo/discovery.etcd.io which in turn is forked from https://github.com/coreos/discovery.etcd.io.

## Development

discovery.etcd.io uses devweb for easy development. It is simple to get started:

```
./devweb
curl --verbose -X PUT localhost:8087/new
```

## Run in docker with default settings

```
docker run -d -p 8087:8087 IMAGE_NAME
```

## Run in docker with custom settings

There are three ENVs to control discovery:

1. `DISCOVERY_ROOT_URL` : default https://discovery.etcd.io
2. `DISCOVERY_ORIGIN_ADDR` : default http://127.0.0.1:4001
3. `DISCOVERY_INIT_LEADER` : default 127.0.0.1:4001

Pass them like this: `docker run -d -p 8087:8087 -e DISCOVERY_ROOT_URL=https://mydomain.com IMAGE_NAME`
