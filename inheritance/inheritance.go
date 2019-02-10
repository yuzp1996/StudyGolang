package inheritance

import "fmt"

type Author struct{
	Firstname string
	Lastname string
	Age int
}

func(a Author)fullName() string{
	fullname := a.Firstname + a.Lastname
	return fullname
}

type Post struct {
	Title string
	Content string
	Author
}

func(p Post)Details(){
	fmt.Printf(" p Title is %s \n",p.Title)
	fmt.Printf("p contetn is %s \n", p.Content)
	fmt.Printf("p author Firstname is %s \n", p.Firstname)
	fmt.Printf("p author full name  is %s \n", p.fullName())
	fmt.Printf("p author full name  is %s \n", p.Author.fullName())
}


type Website struct {
	 Posts []Post
}

func(web Website)Contents(){
	for what ,post := range web.Posts {
		fmt.Printf("waht is %v", what)
		post.Details()
	}
}