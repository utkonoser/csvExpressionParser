package implementQueue

type Queue struct {
	Fields []string
	Rows   []string
}

func (q *Queue) InsertField(item string) {
	q.Fields = append(q.Fields, item)
}

func (q *Queue) InsertRow(item string) {
	q.Rows = append(q.Rows, item)
}

func (q *Queue) RemoveField() string {
	returnField := q.Fields[0]
	q.Fields = q.Fields[1:]
	return returnField
}

func (q *Queue) RemoveRow() string {
	returnRow := q.Rows[0]
	q.Rows = q.Rows[1:]
	return returnRow
}

func (q *Queue) InRows(row string) bool {
	for _, r := range q.Rows {
		if r == row {
			return true
		}
	}
	return false
}

func (q *Queue) InFields(field string) bool {
	for _, f := range q.Fields {
		if f == field {
			return true
		}
	}
	return false
}
