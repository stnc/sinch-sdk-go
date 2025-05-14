package batches

func (s *Batches) DryRun() string {
	return " DryRun SMS " + s.C.ClientId
}

