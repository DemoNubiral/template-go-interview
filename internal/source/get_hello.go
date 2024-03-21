package source

type GetHello struct {
}

type GetHelloCMD struct {
	Parameter1 string
	Parameter2 string
}

type GetHelloResponse struct {
	Data string `json:"data"`
}

func NewGetHello() GetHello {
	return GetHello{}
}

func Execute(cmd GetHelloCMD) (GetHelloResponse, error) {

	helloService := NewHelloService()

	name, _ := helloService.GetName(cmd.Parameter1, cmd.Parameter2)

	result := GetHelloResponse{Data: name}

	return result, nil
}
