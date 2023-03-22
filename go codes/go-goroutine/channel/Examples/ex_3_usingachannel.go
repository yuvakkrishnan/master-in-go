//solve the Error........... 

package main

func main()  {
	
val, ok := <- channel
if (ok) {
    fmt.Printf("Value is : %d", val)
}


for {
    select {
        case val, ok := <- ch:
            if (ok) {
                list = append(list, err)
            } else {
                break;
            }
        case default :
			
    }
}
for err := range ch {
    list = append(list, err)
}
}