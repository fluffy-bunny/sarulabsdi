package di

type (
	ICloser interface {
		Close()
	}
	ICtor interface {
		Ctor()
	}
	ICtorAndCloser interface {
		ICtor
		ICloser
	}
)
