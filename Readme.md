#mercury200

----
Реализация протокола обмена электросчетчика Меркурий 200.02.
Файл commands.go содержит реализованные команды счетчика. Ниже приведен список этих команд с hex-кодами и соответствующими им функциями.

##команды на чтение

Код | Команда | Функция |
:---: | :---: | :---: |
21 | Текущее время | GetCurrentTime
22 | Лимит мощности | GetPowerLimit
23 | Лимит энергии | GetEnergyLimit
24 | Флаг сезонного перевода часов | GetSeasonSwitchFlag
25 | Пределы ручной коррекции | GetManualCorrectionAmount
27 | Энергия от сброса | GetEnergyFromReset
28 | Версия ПО | GetVersion
29 | Напряжение встроенной батареи | GetBatteryVoltage
2A | Отображаемые на дисплее значения | GetTariffsDisplayOptions
2B | Время последнего выключения | GetLastTurnOffTime
2C | Время последнего включения | GetLastTurnOnTime
2D | Режим работы импульсного выхода | GetImpOutputOptions
2E | Количество тарифов | GetTariffsCount
2F | Серийный номер | GetSerial
30 | Выходные дни | GetHolidays
32 | Энергия на начало месяца | GetEnergyAtMonthStart	
61 | Время последнего вскрытия корпуса | GetLastOpenedTime
62 | Время последнего закрытия корпуса | GetLastClosedTime
63 | Мгновенные значения | GetInstants	
66 | Дата выпуска | GetProductionDate
67 | Интервалы отображения значений на дисплее | GetDisplayIntervals

##команды на запись

Код | Команда | Функция |
:---: | :---: | :---: |
2 | Установка времени | SetCurrentTime
3 | Установка ограничения мощности |  SetPowerLimit
4 | Установка ограничения энергии |  SetEnergyLimit
5 | Установка флага сезонного перевода часов |  SetSeasonSwitchFlag
6 | Установка пределов ручной коррекции |  SetManualCorrectionAmount
7 | Выбор режима работы импульсного выхода |  SetImpOutputOptions	
9 | Выбор отображаемых на дисплее значений |  SetTariffsDisplayOptions
0A | Установка количества тарифов |  SetTariffsCount	
0D | Установка интервалов отображения |  SetDisplayIntervals	
10 | Запись выходных дней  |  SetHolidays

###пример использования

*запрос значений накопленной энергии по тарифам на начало 11 месяца*

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
}
```

```shell
$ go run main.go
&{0684.92 0342.65 000.0 000.0}
```