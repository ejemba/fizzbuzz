package api

import (
	"sort"
	"expvar"
)

type RequestStatItem struct {
	ParameterRepresentation string
	RequestHits int
}

// Inc increments the number of hits RequestHits
func (r *RequestStatItem) Inc()  {
	r.RequestHits = r.RequestHits + 1
}

func RegisterMetrics() {
	expvar.Publish("top.requests", expvar.Func(AllTopRequests))
}

var RootTopRequests *TopRequests = NewTopRequests()

func AllTopRequests() interface{} {
	return RootTopRequests.Stats()
}

func NewTopRequests() *TopRequests {
	return &TopRequests{requestMap: make(map[string]*RequestStatItem)}
}

type TopRequests struct {
	requestMap map[string]*RequestStatItem
}

// AddStat increment the statistics 
func (tr *TopRequests) AddStat(request  *FizzBuzzRequest)  {
	_ , present := tr.requestMap[request.String()]
	if present {
		tr.requestMap[request.String()].Inc() 
	} else {
		tr.requestMap[request.String()] = &RequestStatItem{RequestHits: 1,ParameterRepresentation: request.String()}
	}
	
}

func (tr *TopRequests) Stats() interface{}  {	
	slices := make([]RequestStatItem,0)
	
	for r := range tr.requestMap {
		if r != "" {			
			slices = append(slices,*tr.requestMap[r])
		}
	}
	
	sort.Slice(slices, func(i,j int) bool {
		return slices[i].RequestHits > slices[j].RequestHits
	})	

	return slices
}
