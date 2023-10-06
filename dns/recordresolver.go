package dns

import "fmt"

type RecordResolver struct {
	DnsResolver
	value string
	Child DnsResolver
}

func (slf *RecordResolver) SetValue(v string) {
	slf.value = v
}

func (slf *RecordResolver) Resolve() ([]Resolution, error) {
	slf.Child.SetValue(slf.value)
	result, err := slf.Child.Resolve()
	if err != nil {
		return []Resolution{}, fmt.Errorf("lookup failed for %s. %s", slf.value, err.Error())
	}
	return append([]Resolution{{Type: "", Value: slf.value}}, result...), nil
}
