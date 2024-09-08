package config

type ErrorManager struct{}

func NewErrorManager() *ErrorManager {
	return &ErrorManager{}
}

func (e *ErrorManager) CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
