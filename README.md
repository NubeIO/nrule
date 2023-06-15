# nrule

## example scope for the editor
```
// JS code
let devs = R1.Ping("rc", "modbus") // ping devices via tags and return a list of network/devicesUUIDs
let getHistories = R1.Hists(devs, "temp","room","zone", "last15min") // call histories for via UUIDs and tags (any device that ping failed we will not call histories on)
let lessThen18 = R1.DFs(getHistories, "<", "18", "last15min") // dataframe function that can do some analysis
let makeAlerts = R1.MakeAlerts(lessThen18) // make all the alerts on what was returned
R1.PrintJson(makeAlerts) // print some info
```

## use case for an OEM developer 
dev can use the core nube-rules lib, so we can do the core base code with all the functions to get the data from FF and core functions for analysis
```
// JS code
let hist = CoreLib.GetHist() // core lib by nube with all basic functions
NonNubeContrib.MyOEMFunction(hist) // non nube developer can import the nube-rules lib and make and pass in their own functions
```