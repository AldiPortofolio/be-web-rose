package dbmodels

type InstructionList struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
	Sequence    int    `json:"sequence"`
}

func (q *InstructionList) TableName() string {
	return "public.instruction_list"
}
