package urlshorter

type (
	App interface {
		Start() error
		Stop() error
	}
)
