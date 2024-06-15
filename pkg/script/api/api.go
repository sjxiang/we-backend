package main

import (
   "fmt"
   "strings"
   "net/http"
   "io"
)

func main() {

   url := "http://127.0.0.1:5678/api/v1/user/login"
   method := "POST"

   payload := strings.NewReader(`
{
    "email": "1535484943@qq.com",
    "password": "123456@qwe"
}`)

   client := &http.Client{}

   req, err := http.NewRequest(method, url, payload)

   if err != nil {
      fmt.Println(err)
      return
   }
   req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
   req.Header.Add("Content-Type", "application/json")

   resp, err := client.Do(req)
   if err != nil {
      fmt.Println(err)
      return
   }
   defer resp.Body.Close()

   
   body, err := io.ReadAll(resp.Body)
   if err != nil {
      fmt.Println(err)
      return
   }

   fmt.Println(resp.StatusCode)
   fmt.Printf("%+v\n", resp.Header)
   fmt.Println(string(body))
}