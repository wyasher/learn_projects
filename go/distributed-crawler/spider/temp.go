package spider

type Temp struct {
	data map[string]any
}

func (t *Temp) Get(key string) any {
	return t.data[key]
}

func (t *Temp) Set(key string, value any) error {
	if t.data == nil {
		t.data = make(map[string]any, 8)
	}
	t.data[key] = value
	return nil
}
