package main

import(
  "net/http"
  "bytes"
  "encoding/json"
  "fmt"
  "io/ioutil"
)

type User struct{
  Name string `json:"name"`
  Job string  `json:"job"`
}

func main(){
  user := User{
    Name: "User",
    Job: "Developer",
  }

  body, _ := json.Marshal(user)
  resp, err := http.Post("http://127.0.0.1:5000/api/notification", "application/json", bytes.NewBuffer(body))

  if err != nil{
    panic(err)
  }

  defer resp.Body.Close()

  if resp.StatusCode == http.StatusCreated{
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil{
      panic(err)
    }

    jsonStr := string(body)
    fmt.Println("Response: ", jsonStr)
  }
}
