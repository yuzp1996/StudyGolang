package inheritance

import "fmt"

//组合取代继承


type Website struct {
	Posts []Post
}

//匿名结构体可以直接进行访问的偶
type Post struct {
	Title string
	Content string
	Author
}

type Author struct{
	Firstname string
	Lastname string
	Age int
}


func(web Website)Contents(){
	for what ,post := range web.Posts {
		fmt.Printf("waht is %v", what)
		post.Details()
	}
}

func(p Post)Details(){
	fmt.Printf(" p Title is %s \n",p.Title)
	fmt.Printf("p contetn is %s \n", p.Content)
	fmt.Printf("p author Firstname is %s \n", p.Firstname)
	fmt.Printf("p author full name  is %s \n", p.fullName())
	fmt.Printf("p author full name  is %s \n", p.Author.fullName())
}

func(a Author)fullName() string{
	fullname := a.Firstname + a.Lastname
	return fullname
}


