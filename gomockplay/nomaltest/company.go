package nomaltest

type Company struct {
	//现在感觉这里的接口更像是一个功能的集合 实现这些功能的 你就是我想要的
	Usher Takler
}



func NewCompany(takler Takler)*Company{
	return &Company{
		Usher:takler,
	}
}

func (c *Company)Meeting(guestName string)string{
	return c.Usher.SayHello(guestName)
}