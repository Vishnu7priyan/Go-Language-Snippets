//html-tool is developed by tomnomnom 

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "regexp"
    "os/exec"
    execute "github.com/alexellis/go-execute/pkg/v1"
)

func GetAllUrls(user_input string) {
    url := "https://"+user_input
    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
     data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    stringBody := string(data)

    re := regexp.MustCompile(`<a.*href\s*=\s*["'](https?://[^"']+)["'].*>`)
    links := re.FindAllStringSubmatch(stringBody, -1)

    fmt.Println(len(links))
    for _, link := range links {
        fmt.Println(link[1])
        
    }

 command := "echo"
    arg1 := "https://" + user_input
    arg2 := "|"
    arg3 := "html-tool"
    arg4 := "attribs"
    arg5 := "href"
    arg6 := "echo"
cmd := exec.Command(command,arg6, arg1, arg2, arg3, arg4, arg5)

    output, err := cmd.CombinedOutput()
    if err != nil {
        panic(err)
    }

    final_comm := string(output)
    executing := execute.ExecTask{
        Command: final_comm,
        Shell: true,}
    res , _ := executing.Execute()
    fmt.Println(res.Stdout)
   
    
}
