package perusahaan

type VP struct {
	Subordinate []Employee
}

func (vp VP) GetSalary() int {
	return 20
}

func (vp VP) TotalDivisonSalary() int {
	//0 // TODO: replace this
	var sum int
	for _, v := range vp.Subordinate {
		sum = sum + v.GetSalary()
	}
	return vp.GetSalary() + sum
}
