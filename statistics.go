package main

func newStatistic() *Statistic {
	return &Statistic{
		Sum:      0.0,
		Avg:      0.0,
		Max:      -MaxFloat64,
		MaxQueue: make([]float64, 0),
		Min:      MaxFloat64,
		MinQueue: make([]float64, 0),
		Count:    0,
	}
}
