package register

type exampleItem interface {
	doing()
}

func exampleLogic(items ...exampleItem) {
	for _, item := range items {
		item.doing()
	}
}

// -----------------------------------------------------------------------------

type exampleA struct{}

func (e exampleA) doing() {
	//TODO implement me
	panic("implement me")
}

type exampleB struct{}

func (e exampleB) doing() {
	//TODO implement me
	panic("implement me")
}

type exampleC struct{}

func (e exampleC) doing() {
	//TODO implement me
	panic("implement me")
}

type exampleD struct{}

func (e exampleD) doing() {
	//TODO implement me
	panic("implement me")
}

type exampleE struct{}

func (e exampleE) doing() {
	//TODO implement me
	panic("implement me")
}
