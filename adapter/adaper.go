package adapter

type Adapter struct {
	oldInterface *OldInterface
}

func (a *Adapter) DoNewSomething(text string) {
	a.oldInterface.DoOldSomething(text)
}
