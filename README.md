# abredis

[Redis](https://github.com/go-redis/redis) client extensions

## Try it out

docker-compose up
docker exec -ti redis sh
redis-cli
set config:app example
del config:app example

You should notice in the docker-compose output that key modification has been detected.