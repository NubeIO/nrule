# nrule

to run
```
go mod tidy
go run main.go 
```

ping server
```
http://0.0.0.0:1666/api/ping
```


```js
let a = 22
let out = {}

out.uuid = `math: ${a*10} call RQL func date: ${RQL.TimeDate()}`

RQL.Return = Client.ToString(out)
```


```js
let body = {
    host_uuid: "hos_c4be792c63c74454",
    type: "ping",
    entity_type: "network"
}

let out = {}

let result = RQL.AddAlert(body.host_uuid, body)


out.uuid = `added new alert date: ${RQL.TimeDate()} id: ${result.Result.UUID}`

RQL.Return = Client.ToString(out)
```



