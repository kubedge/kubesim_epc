package main
import (
  "log"
  "net/http"
  "strings"
  "github.com/kubedge/kubesim_base/grpc/go/kubedge_server"
)
func sayHello(w http.ResponseWriter, r *http.Request) {
  message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  message = "kubesim epc simulator " + message
  w.Write([]byte(message))
}
func main() {
  log.Printf("%s", "kubesim_epc is running")
  //run server forever
  server.Server()
  http.HandleFunc("/", sayHello)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
