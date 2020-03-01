package collector

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	difficultyDesc = prometheus.NewDesc(
		"eth_block_difficulty",
		"Ethereum block difficulty",
		[]string{}, nil,
	)
	gasLimitDesc = prometheus.NewDesc(
		"eth_block_gas_limit",
		"Ethereum block gas limit",
		[]string{}, nil,
	)
	gasUsedDesc = prometheus.NewDesc(
		"eth_block_gas_used",
		"Ethereum block gas used",
		[]string{}, nil,
	)
)

type HeaderCollecter struct {
	mutex chan bool
	queue []*types.Header
}

func NewHeaderCollector() *HeaderCollecter {
	return &HeaderCollecter{
		mutex: make(chan bool),
		queue: make([]*types.Header, 0),
	}
}

func (col HeaderCollecter) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(col, ch)
}

func (col HeaderCollecter) Collect(ch chan<- prometheus.Metric) {
	// TODO: Grab lock
	for _, header := range col.queue {
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
	// TODO: Release lock
}

func (col HeaderCollecter) Measure(header *types.Header) {
	// TODO: Implement thread safe locking
	col.queue = append(col.queue, header)
}
