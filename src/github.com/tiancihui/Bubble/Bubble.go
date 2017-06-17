package main
import "fmt"

func BubbleSort(BubbleSlice [] int) (ret [] int){
	
		length :=len(BubbleSlice)
		for i:=0; i<length; i++{
			for j:=length - 2; j >= 0; j-- {
				if BubbleSlice[ j ] > BubbleSlice[ j + 1 ] {
                v:= BubbleSlice[j]
	            BubbleSlice[j]  = BubbleSlice[j+1]
	            BubbleSlice[j+1] = v
				}
			}	
	}

    return BubbleSlice
	}


func BubblePrint(BubbleSlice [] int)( ){
	fmt.Println( "------------------ Print Slice Begin -------------")	
		for i:=0; i< len(BubbleSlice); i++{
			fmt.Println("BubbleSlice [" , i , "]", BubbleSlice[i])	
		}
	fmt.Println( "------------------ Print Slice End -------------")	
}




func main(){

	var BubbleArray [10] int = [10] int {10,9,8,7,6,5,4,3,2,1}
	var BubbleSlice [] int = BubbleArray[2:8]

		for _, v:=range BubbleArray{
			fmt.Println( v, "   ")	
		}

	fmt.Println( " hello word \r\n")	

   BubblePrint(BubbleSlice)


	fmt.Println( " hello word \r\n")	

		BubbleSlice = BubbleSort(BubbleSlice)


       BubblePrint(BubbleSlice)
	

	fmt.Println( " hello word \r\n")	
}
