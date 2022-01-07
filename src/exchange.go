package messenger

type Exchange struct {
	name string
	strategy string
	autoRemove bool
}

func NewExchange() Exchange {
	return Exchange{strategy: "direct"}
}

func (e Exchange) Named(name string) Exchange {
	e.name = name

	return e
}

func (e Exchange) WithStrategy(strategy string) Exchange {
	e.strategy = strategy

	return e
}

func (e Exchange) ShouldAutoRemove() Exchange {
	e.autoRemove = true

	return e
}
