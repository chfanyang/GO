package main

import "fmt"

func main() {
	for i:=1;i<10;i++{
		for j:=1;j<=i;j++{
			//fmt.Print(i,"*",j,"=",i*j,"\t")
			fmt.Printf("%d * %d = %d\t",i,j,i*j)
		}
		fmt.Println()
	}
}