# ZTable
This repository provides access to a static z-table (standard normal table) through a few lookup functions. The static data is loaded into memory at instantiation using NewTable(). You can access the data through three functions:
* GetFuzzy(p float64) []*ZScore
* GetScore(p float64) float64
* GetProbability(s float64) float64
