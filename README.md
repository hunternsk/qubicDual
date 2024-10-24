## HiveOS qubic dual mining autopilot

### Command line args:

-fs Show flight sheets in farm

### Configuration:

.env file

    ACCESSTOKEN=YourAccessToken
    FARMID=YourFarmId
    QUBICFSID=QubicFS
    IDLEFSID=WhatToStartWhenQubicIsIdle
    EXCLUDEWORKERS=*
    INCLUDEWORKERS=CaseSensitiveToo

### Building

    go build