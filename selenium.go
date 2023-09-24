package main 
import (
  "fmt")
service, err := selenium.NewSeleniumService(    seleniumPath,     8080,     selenium.GeckoDriver(geckoDriverPath)) 
if err != nil {    
  panic(err)  } 
defer service.Stop() 
}
