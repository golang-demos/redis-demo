# Fiber + Redis Demo


## Install and run Redis
*Make sure you have docker installed.
```
docker pull redis
docker run --name fiber-redis-demo -p 6379:6379 redis
```

## Run script	
```
go run main.go
```