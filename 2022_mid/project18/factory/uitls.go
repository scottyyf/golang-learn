package factory

type student struct {
	Name string
	score float32
}

func (stu *student) GetScore() float32 {
	return stu.score
}

func NewStudent(name string, score float32) *student{
	return &student{
		Name: name,
		score: score,
	}
}