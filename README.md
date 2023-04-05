### Iniciar modulo

```
go mod init server
```

#### lib
```
go get -u github.com/....
```

### Daemon
```
fresh
```

### Correr server local
```
go run .
```

### Crear ejecutable
```
go build server
```
### Crear bin
```
go build -o server .
```

### Ejecutar app
```
./server
```

# DOCKER

### Crear imagen
```
docker build -t server-go:latest .
```

### Correr Contendor
```
docker run --name servergo -p 8080:8808 server-go:latest
```
