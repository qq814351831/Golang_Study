package utils


//数据表字段
type ExamScore struct {
	Id int `db:"id"`
	Name string `db:"name"`
	Score int `db:"score"`
}
