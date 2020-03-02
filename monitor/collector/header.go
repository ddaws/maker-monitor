package collector

import (
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	difficultyDesc = prometheus.NewDesc(
		"eth_block_difficulty",
		"Ethereum block difficulty",
		nil, nil,
	)
	gasLimitDesc = prometheus.NewDesc(
		"eth_block_gas_limit",
		"Ethereum block gas limit",
		nil, nil,
	)
	gasUsedDesc = prometheus.NewDesc(
		"eth_block_gas_used",
		"Ethereum block gas used",
		nil, nil,
	)
)

type HeaderCollecter struct {
	mutex *sync.Mutex
	queue []*types.Header
}

func NewHeaderCollector() *HeaderCollecter {
	return &HeaderCollecter{
		mutex: &sync.Mutex{},
		queue: make([]*types.Header, 0),
	}
}

func (col *HeaderCollecter) Describe(ch chan<- *prometheus.Desc) {
	//prometheus.DescribeByCollect(col, ch)
	ch <- gasLimitDesc
	ch <- gasUsedDesc
}

func (col *HeaderCollecter) Collect(ch chan<- prometheus.Metric) {
	log.Printf("Collect(queue=%d)\n", len(col.queue))
	// Take a lock to prevent raise conditions accessing queue
	col.mutex.Lock()
	// Dump accumulated header metrics
	for _, header := range col.queue {
		log.Printf("gasLimit=%f | gasUsed=%f", float64(header.GasLimit), float64(header.GasUsed))
		// TODO: Assume timestamp from the block
		ch <- prometheus.MustNewConstMetric(
			gasLimitDesc,
			prometheus.GaugeValue,
			float64(header.GasLimit),
		)
		ch <- prometheus.MustNewConstMetric(
			gasUsedDesc,
			prometheus.GaugeValue,
			float64(header.GasUsed),
		)
	}
	col.queue = make([]*types.Header, 0)
	// Release lock
	col.mutex.Unlock()
}

func (col *HeaderCollecter) Measure(header *types.Header) {
	log.Println("Measure")
	col.mutex.Lock()
	col.queue = append(col.queue, header)
	col.mutex.Unlock()
}
