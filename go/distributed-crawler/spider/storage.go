package spider

type Storage interface {
	Save(data ...*DataCell) error
}

type DataCell struct {
	Task *Task
	Data map[string]any
}

func (d *DataCell) GetTableName() string {
	return d.Data["Task"].(string)
}

func (d *DataCell) GetTaskName() string {
	return d.Data["Task"].(string)
}
