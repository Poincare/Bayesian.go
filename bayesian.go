package Bayesian

type Bayesian struct {
	prior []float64 
	likelyhood []float64
	bayesian []float64
	num_classes int
}

func (bay *Bayesian) Initialize(num_classes int) {
	bay.num_classes = num_classes

	bay.prior = make([]float64, num_classes)
	bay.likelyhood = make([]float64, num_classes)
	bay.bayesian = make([]float64, num_classes)
}

func (bay *Bayesian) SetPrior(pr []float64) {
	copy(bay.prior, pr[0:])
}

func (bay *Bayesian) SetLikelyhood(lk []float64) {
	copy(bay.likelyhood, lk[0:]) 
}

func (bay *Bayesian) MaxInSlice(sl []float64) float64 {
	max := 0.0

	max_index := -1.0

	for i := 0; i < len(sl); i++ {
		if sl[i] > max {
			max = sl[i]
			max_index = float64(i)
		}		
	}

	return max_index
}

func (bay *Bayesian) CrunchBayesian() float64 {	
	for i := 0; i < bay.num_classes; i++ {
		var res float64 = bay.prior[i] * bay.likelyhood[i];
		
		append(bay.bayesian, res);
	}

	return bay.MaxInSlice(bay.bayesian)
}

