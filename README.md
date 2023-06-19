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
function timestamp(){
  function pad(n) {return n<10 ? "0"+n : n}
  d=new Date()
  dash="-"
  colon=":"
  return d.getFullYear()+dash+
  pad(d.getMonth()+1)+dash+
  pad(d.getDate())+" "+
  pad(d.getHours())+colon+
  pad(d.getMinutes())+colon+
  pad(d.getSeconds())
}


let body = {
    host_uuid: "hos_c4be792c63c74454",
    type: "ping",
    entity_type: "network"
}

let out = {}

let result = Client.AddAlert(body.host_uuid, body)


out.uuid = `added new alert date: ${timestamp()} id: ${result.Result.UUID}`

Client.Result = Client.ToString(out)
```



