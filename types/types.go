package types

type DisplayIntervals struct {
	InactiveTEnergy, ActiveTEnergy, Instants, Additionals int
}

type TariffsDisplayOptions struct {
	Date, Time, Power, TSumm, T4, T3, T2, T1 string
}

type Energy struct {
	T1, T2, T3, T4 string
}
