package main
import(
 "log"  )
func init() {
 log.SetPrefix("TRACE: ")
 log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)  
}

