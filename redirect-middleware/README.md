This is redirect middleware APP with Docker Solution.


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
local dev env
```
docker run -it -p 3000:3000 redirect
```
deploy to server, then run app 
```
docker run -td -p 3000:3000 mikewolfxyou/redirect
```

## Try use app
In browser call `localhost:3000`, then you will be redirect to other web


