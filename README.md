# INSTRUCTION


sample env file data
```yml
# .env.dev.yml
app:
  server:
    port: :8080
```


to run it on develop env
```sh
go run main.go server dev
```

to run it on staging env
```sh
go run main.go server stage
```
