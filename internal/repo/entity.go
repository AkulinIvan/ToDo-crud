package repo

type StatusType string

const (
	NewStatus        StatusType = "new"
	InProgressStatus StatusType = "in_progress"
	DoneStatus       StatusType = "done"
)

func (s StatusType) String() string {
	switch s {
	case NewStatus:
		return "new"
	case InProgressStatus:
		return "in_progress"
	case DoneStatus:
		return "done"
	default:
		return "new"
	}
}

// Task - структура, соответствующая таблице tasks
type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      StatusType `json:"status"`
}