This is mini Go Lang APP with Docker Solution.

The idea from https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/

## Build app first

Statically compile app with all libraries built in
```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
```

## Build image
```
docker build -t redirect -f Dockerfile .
```

## Run app
```
docker run -it redirect
```


