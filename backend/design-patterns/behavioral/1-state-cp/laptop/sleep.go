package laptop

type Sleeping struct {
	Laptop *Laptop
}

func (s Sleeping) Press() {
	// TODO: answer here
	s.Laptop.ChangeState(On{s.laptop})
}

func (s Sleeping) CanTurnOnLaptop() bool {
	return true
}

func (s Sleeping) Sleep() {
	// TODO: answer here
	s.Laptop.ChangeCurrentState("Sleeping")
	s.Laptop.ChangeState(On{s.laptop})
}
