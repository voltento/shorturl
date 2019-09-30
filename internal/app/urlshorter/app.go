package urlshorter

type appImpl struct {
}

func NewApp() (App, error) {
	return &appImpl{}, nil
}

func (app *appImpl) Start() error {
	return nil
}

func (app *appImpl) Stop() error {
	return nil
}
