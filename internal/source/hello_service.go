package source

type helloService struct {
}

func NewHelloService() helloService {
	return helloService{}
}

func (hel *helloService) GetName(parameter1 string, parameter2 string) (string, error) {

	helloRepository := NewHelloRepository()

	name, error := helloRepository.GetName(parameter1, parameter2)

	if error != nil {
		return "", error
	}

	return name, nil
}
