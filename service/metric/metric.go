package metric

// Metric for service
type Metric struct {
	requestsCount int
}

func (m *Metric) GetRequestsCount() int {
	return m.requestsCount
}

func (m *Metric) IncRequestsCount() {
	m.requestsCount++
}
