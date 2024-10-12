package service

type ServiceRepository interface {
	AddGauge(metricName, value string) error
	AddCounter(metricName, value string) error
	GetMetric(metricType, metricName string) (string, error)
	GetAllMetrics() map[string]string
}

type Service struct {
	serviceRepo ServiceRepository
}

func NewService(s ServiceRepository) *Service {
	return &Service{
		serviceRepo: s,
	}
}

func (s *Service) AddGauge(metricName, value string) error {
	err := s.serviceRepo.AddGauge(metricName, value)
	if err != nil{
		return err
	}

	return nil
}

func (s *Service) AddCounter(metricName, value string) error {
	err := s.serviceRepo.AddCounter(metricName, value)
	if err != nil{
		return err
	}

	return nil
}

func (s *Service) GetMetric(metricType, metricName string) (string, error) {
	value, err := s.serviceRepo.GetMetric(metricType, metricName)
	if err != nil{
		return "", err
	}

	return value, nil
}

func (s *Service) GetAllMetrics() map[string]string {
	return s.serviceRepo.GetAllMetrics()
}