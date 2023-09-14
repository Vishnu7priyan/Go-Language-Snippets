package main
import(
 "log"  )
func init() {
 log.SetPrefix("LOG TRACE: ")
 log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)  
}

