package main
import(
   "fmt"
   "time"
)

func routineFunc(){

	for i:=0 ; i < 1000; i++{
		fmt.Println( "routine func : " , i)
	}

}


func main(){
	
	
	for i:=0 ; i < 1000; i++{
		if i == 100{
			go routineFunc()
			}

		
		if i == 500 {
			//time.Sleep(1000)
			}
			
		fmt.Println( "main func : " , i)
	}
	
	
   for i:=0; i< 10; {
	   
			time.Sleep(1000)

	   }	
	
	
	}
