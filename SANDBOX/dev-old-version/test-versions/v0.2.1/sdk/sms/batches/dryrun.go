package batches

func (s *BranchRepo) DryRun() string {
	return " DryRun SMS " + s.C.ClientId
}

