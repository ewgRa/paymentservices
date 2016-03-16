package metric

// Metric for service
type Metric struct {
	requestsCount int
}

// GetRequestsCount return count of requests that service handled during his life
func (m *Metric) GetRequestsCount() int {
	return m.requestsCount
}

// IncRequestsCount increment requests count
func (m *Metric) IncRequestsCount() {
	m.requestsCount++
}
