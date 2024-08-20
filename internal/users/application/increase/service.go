package increase

type UserCounterService struct{}

func NewUserCounterService() UserCounterService {
	return UserCounterService{}
}

func (s UserCounterService) Increase(id string) error {
	return nil
}
