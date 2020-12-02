package main

import(
	"fmt"
	"errors"
)

var ErrNoRows = errors.New("err no rows")

func dao() error{
	return errors.Wrap(ErrNoRows,"dao msg")
}

func service() error{
	err := dao()
	if err != nil{
		return errors.WithMessage(err,"service msg")
	}	
	return nil
}

func main(){
	err != service()
	if err != nil{
		fmt.Printf("main: %+v\n",err)
	}
}