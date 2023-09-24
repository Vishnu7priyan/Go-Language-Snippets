package main 
import (
  "fmt")
service, err := selenium.NewSeleniumService(    seleniumPath,     8080,     selenium.GeckoDriver(geckoDriverPath)) 
if err != nil {    
  panic(err)  } 
defer service.Stop() 
caps := selenium.Capabilities{"browserName": "firefox"}  
wd, err := selenium.NewRemote(caps, "http://localhost:8080/") 
if err != nil {    panic(err)  } 
defer wd.Quit()
}
