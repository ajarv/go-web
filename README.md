# go-web

```bash

docker build -t go-web .

docker run -d -p 7080:8080 --name go-web go-web 

curl localhost:7080/reqinfo
```