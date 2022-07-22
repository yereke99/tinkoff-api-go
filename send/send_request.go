package send

import(
  "net/http"
  "bytes"
  "encoding/json"
  "fmt"
  "io/ioutil"
)

func Send_notification(data interface{}){
  body, _ := json.Marshal(data)


  resp, err := http.Post("http://127.0.0.1:5000/api/notification", "application/json", bytes.NewBuffer(body))
  if err != nil{
    log.Fatalf("Failed to send request: %s", err)
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
