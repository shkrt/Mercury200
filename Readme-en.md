# Mercury200

----
Implementation of Mercury-200 power meter's exchange protocol.
The commands.go file contains all implemented commands. Below is the list of commands with hex codes and corresponding functions in code.

## reading commands

Code | Command | Function |
--- | --- | --- |
21 | Current time | GetCurrentTime
22 | Power limit | GetPowerLimit
23 | Energy limit | GetEnergyLimit
24 | Seasonal time shift flag | GetSeasonSwitchFlag
25 | Limits of manual correction | GetManualCorrectionAmount
27 | Energy from last reset | GetEnergyFromReset
28 | Firmware version | GetVersion
29 | Voltage of builtin battery | GetBatteryVoltage
2A | Displayed values | GetTariffsDisplayOptions
2B | Last turnoff time | GetLastTurnOffTime
2C | Last standby time | GetLastTurnOnTime
2D | Impulse output operation mode | GetImpOutputOptions
2E | Number of tariffs | GetTariffsCount
2F | Serial number | GetSerial
30 | Holidays | GetHolidays
32 | Energy at month start time | GetEnergyAtMonthStart	
61 | Last case opening time | GetLastOpenedTime
62 | Last case closing time | GetLastClosedTime
63 | Instant values | GetInstants	
66 | Production date | GetProductionDate
67 | Values displaying intervals | GetDisplayIntervals

## writing commands

Code | Command | Function |
--- | --- | --- |
2 | Set current time | SetCurrentTime
3 | Set power limit |  SetPowerLimit
4 | Set energy limit |  SetEnergyLimit
5 | Set seasonal time shift flag |  SetSeasonSwitchFlag
6 | Set manual correction limits |  SetManualCorrectionAmount
7 | Set impulse output operation mode |  SetImpOutputOptions	
9 | Choose displayed values |  SetTariffsDisplayOptions
0A | Set number of tariffs |  SetTariffsCount	
0D | Set displaying intervals |  SetDisplayIntervals	
10 | Set holidays |  SetHolidays

### usage example

* get values of accumulated energy with tariffs breakdown at the start of 11th month

```go
package main

import (
	"fmt"
	"mercury200/commands"
)

func main() {
	netNum := "266608"
	port := "COM5"
	timeOut := 5
	baudRate := 9600

	result, _ := commands.GetEnergyAtMonthStart(&netNum, &port, &timeOut, &baudRate, 11)
	fmt.Println(result)

	fmt.Printf("Tariff 1: %s kW\n", result.T1)
	fmt.Printf("Tariff 2: %s kW", result.T2)
}
```

```shell
$ go run main.go
&{0684.92 0342.65 000.0 000.0}
Tariff 1: 0684.92 kW
Tariff 2: 0342.65 kW
```

