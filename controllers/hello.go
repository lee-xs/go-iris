package controllers

type HelloController struct {}

func (h *HelloController) GetHello() string{
	return "hello world"
}

func (h *HelloController) Get() string {
	return "hello"
}
