package dns

type EmptyResolver struct {
	DnsResolver
}

func (slf *EmptyResolver) SetValue(_ string) {}

func (slf *EmptyResolver) Resolve() ([]Resolution, error) {
	return []Resolution{}, nil
}
