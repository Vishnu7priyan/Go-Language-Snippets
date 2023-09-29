import ( 
"fmt"  
"io"
"net/http" 
"os" ) 

func init() {
 if len(os.Args) != 2 { 
 fmt.Println("Usage: ./example2 <url>") 
os.Exit(-1)  } 
}  
