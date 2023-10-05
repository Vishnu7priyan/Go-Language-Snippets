//Declare
var array [5]int
//Declare with values 
array := [5]int{10, 20, 30, 40, 50}
//Capacity based on no. of values
array := [...]int{10, 20, 30, 40, 50}
//Declaring on specific index values
array := [5]int{1: 10, 2: 20}
//Change The value
array[2] = 35;
var array1 [5]string 
array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
array1 = array2

var array [4][2]int
array := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
array[0][0] = 10 
array[0][1] = 20 
array[1][0] = 30

slice := []int{10, 20, 30, 40, 50} 

newSlice := slice[1:3] 

newSlice = append(newSlice, 60)

slice2 := slice[2:3:4]
