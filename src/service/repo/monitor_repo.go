package repo

type MonitorRepo struct {
}

func (m *MonitorRepo) HealthCheck() string {
	return "Ok"
}
