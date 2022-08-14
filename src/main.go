package main

import (
 "github.com/gorilla/mux"
 "log"
 "net/http"
 "net/url"
 "io/ioutil"
 "errors"
 "github.com/mateuszpoland/chatbot/lib/structs"
)

// app token : EAAIP9JjQb78BAOcmPFLjzEdqlHWIw6aC598yfZAxfuV4APFhJ6BYz7BV2JXoMrpodt3yMS7du93m1vqSbOkZAfqXk5fjIHnn6lTrmFbh8seqPSGZAT0tOxo6mkd61VVNniXMkcCtkrmlZCTawqZATL4pZB5GkD5wtZA9o6ejosnufsAr1XiN7DL90ADHaLRZByZB0x6Ua2EKhZAgZDZD

func HandleMessenger(resp http.ResponseWriter, request *http.Request) {
 secretKey := "secret_token"
 if request.Method == "GET" {
  u, _ := url.Parse(request.RequestURI)
  values, _ := url.ParseQuery(u.RawQuery)
  token := values.Get("hub.verify_token")
  if token == secretKey {
   resp.WriteHeader(200)
   resp.Write([]byte(values.Get("hub.challenge")))
   return
  }
  resp.WriteHeader(400)
  resp.Write([]byte(`Bad token`))

  return
 }

 // else is post
 body, err := ioutil.ReadAll(request.Body)
 if err != nil {
    log.Pritnf("Failed processing request body: %s", err)
    resp.WriteHeader(400)
    resp.Write([]byte("An error occurred"))
    return
 }

 var message InputMessage
 err = json.Unmarshal(body, &message)
 if err != nil {
   log.Printf("failed to decode request into json: %s", err)
   resp.WriteHeader(400)
   resp.Write([]byte("An error occurred"))

   return
 }

 log.Printf("Message: %#v", message)
 for _, entry := range.message.Entry {
   if(len(entry.Messaging)) == 0 {
      log.Printf("No messages in packet.")
      resp.WriteHeader(400)
      resp.Write([]byte("Error. Empty messsage sent"))

      return
   }

   event := entry.Messaging[0]
   err = handleMessage(event.Sender.ID, event.Message.Text)
 }

}


func handleMessage(senderId, message string) error {
   if len(message) == 0 {
      return errors.New("No message found.")
   }

   response := ResponseMessage {
      Recipient: Recipient {
         ID: senderId
      },
      Message: OutputMessage {
         Text: "Welcome to FUH EKO-AR."
      }
   }

   data, err := json.Marshal(response)
   if err != nil {
      log.Printf("Error encoding output message: %s", err)
      return err
   }

   sendResponseToFacebook(data)
}

func sendResponseToFacebook(data) {
   uri := os.Getenv("FACEBOOK_APP_WEBHOOK_URL")
   uri = fmt.Sprintf("%s?access_token=%s", uri, os.Getenv("FACEBOOK_ACCESS_TOKEN"))
   log.Printf(uri)
   
   req, err := http.NewRequest(
      "POST",
      uri,
      bytes.NewBuffer(data)
   )

   if err != nil {
      log.Printf("Failed creating message: %s.", err)
      return err
   }

   req.Header.add("Content-Type", "application/json")

   client := http.Client{}
   res, err := client.Do(req)
   if err != nil {
      log.Printf("Failed doing request: %s", err)
      return err
   }

   log.Printf("MESSAGE SENT?\n%#v", res)
   return nil
}



// Initialize request
func main() {
 router := mux.NewRouter()
 router.HandleFunc("/", HandleMessenger).Methods("POST", "GET")
 port := ":8000"
 log.Printf("Server started on %s", port)
 log.Fatal(http.ListenAndServe(port, router))
}