package helper

type namingHelper struct {
}

type NamingHelper interface {
	GetPrefix() string
	GetSufix() string
}

func (a namingHelper) GetPrefix() string {
	return "um-prefixo"
}

func (a namingHelper) GetSufix() string {
	return "um-sufixo"
}

func NewNamingHelper() NamingHelper {
	return &namingHelper{}
}
