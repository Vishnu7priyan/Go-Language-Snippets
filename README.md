This repository consist of simple Go Language snippets

1. Anonymous Add and Print 
``` Go
package main

import "fmt"

func main(){
  add := func(x, y int) int {
    return x + y
  }
  result := add(3, 5)
  prin := func(){
    fmt.Println("Hii")
    }
  prin()
}
```
2.Ansi color
```go
package main

import (
        "fmt"
)

func main() {
    red := "\033[31m"
    green := "\033[32m"
    yellow := "\033[33m"
        
    fmt.Println(yellow + "This is yellow")
    fmt.Println(green + "This is green")
    fmt.Println(red + "This is red")
    /*
    More colors
    Black: \033[30m
    Blue: \033[34m
    Magenta: \033[35m
    Cyan: \033[36m
    White: \033[37m
    */
}
```

3.Array
```Go
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
```

4.Binary File
```Go
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("binaryfile.bin")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		panic(err)
	}
}
```

5.CSV
```go
package main

import (
    "encoding/csv"
    "os"
)

func main() {
    // Create a CSV file for writing
    file, err := os.Create("data.csv")
    if err != nil {
        panic(err)
    }
    defer file.Close()
   writer := csv.NewWriter(file)
    defer writer.Flush()

    data := [][]string{
        {"Name", "Age", "City"},
        {"vishnu", "19", "Bash"},
    }
    for _, row := range data {
        if err := writer.Write(row); 
        err != nil {
            panic(err)
        }
    }
    
}
```
6.Database
```go

package main

import (
    "database/sql"
    "fmt"
    "github.com/mattn/go-sqlite3"
    //Install with go get <URL>
)

func main() {
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println("Error :", err)
    }
    //Table creation
    createTable := `
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT,
            email TEXT
        )
    `
    _, err = db.Exec(createTable)
    if err != nil {
        fmt.Println("Error :", err)
    }

    //Updation
    updateSQL := "UPDATE users SET email = ? WHERE username = ?"
    _, err = db.Exec(updateSQL, "test@example.com", "vishnu")
    if err != nil {
        fmt.Println("Error updating data:", err)
        return
    }
    //Deletion
    deleteSQL := "DELETE FROM users WHERE username = ?"
    _, err = db.Exec(deleteSQL, "vishnu")
    if err != nil {
        fmt.Println("Error deleting data:", err)
        return
    }

    
}
```

7.HTTP
```go
package main

import (
    "fmt"
    "net/http"
    ")

func main(){
  url := "URL_HERE"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println("GET Response:")
	fmt.Println(string(body))

  jsonData := `{"title": "Sample Post", "body": "This is a sample", "userId": 1}`
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("POST Response Status:", resp.Status)
  }

```

8. Mongo DB
```go
package main

import (
    "context"
    "fmt"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)


func main() {
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        fmt.Println(err)
        return
    }

    db := client.Database("mydb")
    collection := db.Collection("mycollection")

    document := bson.D{{"name", "John Doe"}, {"age", 30}}

    _, err = collection.InsertOne(context.TODO(), document)
    if err != nil {
        fmt.Println(err)
        return
    }

    defer client.Disconnect(context.TODO())

    fmt.Println("Document inserted successfully!")
}
```

9.TextFile
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    filePath := "textfile.txt"
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Println("Error opening the file:", err)
        return
    }
    defer file.Close() 

    buffer := make([]byte, 1024)

    for {
        n, err := file.Read(buffer)
        if err != nil {
            break
        }
        fmt.Print(string(buffer[:n]))
    }
}
```

10.Time
```go
package main
import (
    "fmt"
    "time")

func main() {
    currentTime := time.Now()
    fmt.Println("Current Time is :", currentTime)
    Tomm := currentTime.Add(24 * time.Hour)
    pastTime := currentTime.Add(-1 * time.Hour)
    fmt.Println("Future Time:", Tomm)
    fmt.Println("Past Time:", pastTime)
    Sleephrs := 10* time.Hour
    time.Sleep(Sleephrs)
}
```
11.Count Link
```go
package main

import (  "fmt"  
"io/ioutil"  
"net/http"  
"strings")

func main() {  
   resp, err := http.Get("URL")  
   if err != nil {    
      panic(err)  }  
   data, err := ioutil.ReadAll(resp.Body) 
   if err != nil {   
      panic(err)  }

   stringBody := string(data)
   numLinks := strings.Count(stringBody, "<a") 
   fmt.Printf("Packt Publishing homepage has %d links!\n", numLinks)
}
```

12. Css Selector
```go

package main
import (  "fmt"  "strconv"  "github.com/PuerkitoBio/goquery")
func main() {  
  doc, err := goquery.NewDocument("URL")
  if err != nil {   
    panic(err)  }  
  doc.Find(`css selector here `).   
  Each(func(i int, e *goquery.Selection) {     
    a,_ = e.Attr("change here")     
    b, _ := e.Attr("change here")     
    c, err = strconv.ParseFloat(b, 64)     
    if err != nil {        
      println("Failed to parse ")      }    
    fmt.Printf("%s ($%0.2f)\n", a, c)    })}
```

13. Curl with IO
```go
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
func main() { 
 r, err := http.Get(os.Args[1])
  if err != nil { 
 fmt.Println(err)
 return  } 
 io.Copy(os.Stdout, r.Body) 
 if err := r.Body.Close(); err != nil { 
fmt.Println(err) 
 } 
}
```

14.Execute Linux
```go
//Default commands
package main

import (
    "fmt"
    "os/exec"
)

func main() {
    cmd := exec.Command("cd", "/tmp")
    err := cmd.Run()
    if err != nil {
        fmt.Println(err)
    }
}
//////////////////////////////////////////////
package main
import (
    "fmt"
    "os/exec"
    "runtime"
)
func main() {
    out, err := exec.Command("ls").Output()
    if err != nil {
        fmt.Printf("%s", err)
    }
    output := string(out[:])
    fmt.Println(output)

  }

/////////////////////////////////////////////
//Custom shell scripts
package main

import (
    "fmt"
    "os/exec"
)

func main() {
    cmd := exec.Command("PATH_HERE")
    err := cmd.Run()
    if err != nil {
        fmt.Println(err)
    }
}
```

15. Flag
```go
package main

import (
    "flag"
    "fmt"
)

func main() {
    var yourname string
    var age int

    flag.StringVar(&yourname, "name", "", "Your name")
    flag.IntVar(&age, "age", 0, "Your age")

    flag.Parse()

    fmt.Printf("Name: %s\n", yourname)
    fmt.Printf("Age: %d\n", age)
}
```

16.Get website File
```go
package main

import (
"net/http"
"os"
"fmt"
)

func main(){
r , err := http.Get("http://www.example.com/index.html")
if err != nil{
        fmt.Println("Error")
        }
	defer response.Body.Close()

bodyLen := 1300 //Change the length based on your requirement
webOut := make([]byte,bodyLen)
r.Body.Read(webOut)

	//writing the content in local file
file,err := os.Create("index.html")
if err != nil {
fmt.Println("Error")
}
defer file.Close()
file.Write(webOut)
fmt.Println("Done")
}
```

17. Internet Connectivity check
```go
package main
import (
        "fmt"
        "net"
        "os"
)
func main(){
        host := "google.com"
        port := "80"
        conn,err := net.Dial("tcp", host+":"+port)
        if err != nil {
                fmt.Println("No internet connection")
                os.Exit(1)
        }
        defer conn.Close()

        fmt.Println("Internet connection is available")
}
```

18.IPInfo
```go
package mypackage
import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)
func Ipinfo(ip string) {
    response, err := http.Get("https://ipinfo.io/" + ip + "/json")
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    defer response.Body.Close()
    data, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println("Error:", err)
    }
    fmt.Println(string(data))
}
```

19. JSON Response Handle
```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "io/ioutil"
)

func main() {
    response, err := http.Get("https://web-check.as93.net/.netlify/functions/dns?url=https://ikea.com")
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    if response.StatusCode == http.StatusOK {
        fmt.Println("GET request successful")

        responseBody, err := ioutil.ReadAll(response.Body)
        if err != nil {
            log.Fatal(err)
        }

        var jsonData interface{}

        if err := json.Unmarshal(responseBody, &jsonData); err != nil {
            log.Fatal(err)
        }
        formattedJSON, err := json.MarshalIndent(jsonData, "", "    ")
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println("JSON Response:")
        fmt.Println(string(formattedJSON))
    } else {
        log.Printf("GET request failed with status: %s\n", response.Status)
    }

      
}
```
20. Log
```go
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
```

21. Map
```go
package main

import(
  "fmt"
  )
func main(){
  colors := map[string] string{
    "Blue":"#f0f8ff",
    "Coral":"#ff7f50",
  }
  for key,value := range colors{
    fmt.Println("%s and %s " , key,value}
}
delete(colors,"Blue")
```

22. POST Request
```go
package main
import(  "net/http"  
        "net/url")
func main() {  
   data := url.Values{}  
   data.Set("s", "Golang")  
   response, err :=http.PostForm("URL", data) 
}
```

23. Selenium
```go
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
```

24. Struct
```go
package main
import (
"fmt")

type student struct{
rollno int
name string
}
func main(){
 var s1 Student = Student{101,"vishnu"}
 var s2 Student  = Student{name:"priyan"}
 fmt.Println(s1)
 fmt.Println(s1.name)
 fmt.Println(s2)
 fmt.Printf("Name of the student with 101 roll no is %v\n",s1.name)
}
```

25. XPATH
```go
package main
import (
 "strings"
 "github.com/antchfx/htmlquery"
)
func main() {
 doc, err := htmlquery.LoadURL("{URL}")
println(doc)
 if err != nil {
 panic(err)
 }
 TextNodes := htmlquery.Find(doc,`//{XPATH_HERE}//text()`)
 if err != nil {
 panic(err)
 }
 println("Result !")
 println("----------------------------------")
 for _, node := range TextNodes {
 text := strings.TrimSpace(node.Data)
 println(text)
 }
}
```
