package dtos

type Dto interface {
	ToModel() interface{}
	ToDto(interface{})
}
