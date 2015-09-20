package types

type DisplayIntervals struct {
	InactiveTEnergy, ActiveTEnergy, Instants, Additionals int
}

type TariffsDisplayOptions struct {
	Date, Time, Power, TSumm, T4, T3, T2, T1 string
}
