package readers

import (
	"encoding/json"
	"github.com/cloudfoundry/gosigar"
)

func NewLoadAvg() *LoadAvg {
	l := &LoadAvg{}
	l.Data = make(map[string]interface{})
	return l
}

type LoadAvg struct {
	Data map[string]interface{}
}

func (l *LoadAvg) Run() error {
	concreteSigar := sigar.ConcreteSigar{}
	uptime := sigar.Uptime{}
	uptime.Get()

	avg, err := concreteSigar.GetLoadAverage()
	if err != nil {
		return err
	}

	l.Data["LoadAvg1m"] = avg.One
	l.Data["LoadAvg5m"] = avg.Five
	l.Data["LoadAvg15m"] = avg.Fifteen

	return err
}

func (l *LoadAvg) ToJson() ([]byte, error) {
	return json.Marshal(l.Data)
}
