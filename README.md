# abredis

[Redis](https://github.com/go-redis/redis) client extensions to use [Redis notifications](https://redis.io/topics/notifications)

## Try it out

docker-compose up
docker exec -ti redis sh
redis-cli
set config:app example
del config:app example

You should notice in the docker-compose output that key modification has been detected.