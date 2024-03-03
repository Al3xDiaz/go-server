# go server

documentation [here](https://al3xdiaz.github.io/go-server/)

```bash
# run adminer
NETWORK_NAME=`docker network ls | grep go-server | awk '{print $2}'`
docker run \
  --name=adminer --rm \
  --network=$NETWORK_NAME \
	-p "8080:8080" \
  adminer
```
