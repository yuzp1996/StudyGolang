package readfromfile

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/prometheus/common/log"
	"io/ioutil"
	"os"
)

func Readfile(){

	fptr := flag.String("fpath", "test.txt","file path to read from")
	flag.Parse()
	fmt.Println("value of fptr is ", *fptr)

	data, err := ioutil.ReadFile(*fptr)
	if err != nil{
		fmt.Println("file read error ", err)
		return
	}
	fmt.Printf("Contents of file: %s", string(data))
}


func Readfrombuffer(){
	fprt := flag.String("fpath","test.txt","the file read from")
	flag.Parse()
	f, err := os.Open(*fprt)
	if err != nil{
		log.Fatal(err)
	}
	defer func(){
		if err = f.Close(); err != nil{
			log.Fatal(err)
		}
	}()
	//r := bufio.NewReader(f)
	//b := make([]byte,3)
	//for{
	//	if _,err := r.Read(b); err != nil{
	//		fmt.Println("Error reading file ", err)
	//		break
	//	}
	//	fmt.Println(string(b))
	//}
	fmt.Println("read form file line by line")
	s := bufio.NewScanner(f)
	for s.Scan(){
		fmt.Println(s.Text())
	}
	if err = s.Err(); err != nil{
		log.Fatal(err)
	}
	return
}