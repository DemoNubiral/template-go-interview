package source

type helloService struct {
}

func NewHelloService() helloService {
	return helloService{}
}

func (hel *helloService) GetName(parameter1 string, parameter2 string) (string, error) {

	helloRepository := NewHelloRepository()

	name, _ := helloRepository.GetName(parameter1, parameter2)

	return name, nil
}
