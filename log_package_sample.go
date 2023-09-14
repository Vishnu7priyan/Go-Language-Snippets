package main
import(
 "log"  )
func init() {
 log.SetPrefix("LOG TRACE: ")
 log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)  
}

func main() {
 log.Println("message") 
 log.Fatalln("fatal message")  //Println() followed by a call to os.Exit(1).
 log.Panicln("panic message") //Println() followed by a call to panic().
  }
