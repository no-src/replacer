package replacer

type record struct {
	Old string `json:"old"`
	New string `json:"new"`
}

func newRecord(old, new string) record {
	return record{
		Old: old,
		New: new,
	}
}
