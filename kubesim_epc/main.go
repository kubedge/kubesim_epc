package main
import (
  "log"
  "net/http"
  "strings"
  "github.com/kubedge/kubesim_base/grpc/go/kubedge_server"
  "github.com/kubedge/kubesim_base/config"
  "github.com/kubedge/kubesim_base/connected"
)
func sayHello(w http.ResponseWriter, r *http.Request) {
  message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  message = "kubesim epc simulator " + message
  w.Write([]byte(message))
}
func main() {
  log.Printf("%s", "kubesim_epc is running")

  var conf config.Configdata
  conf.Config()
  log.Printf("epc server:  product_name=%s, product_type=%s, product_family=%s, product_release=%s, feature_set1=%s, feature_set2=%s",
             conf.Product_name, conf.Product_type, conf.Product_family, conf.Product_release, conf.Feature_set1, conf.Feature_set2)

  var conn connected.Connecteddata
  conn.Readconnectvalues()
  log.Printf("epc server:  connected=%s", conn.Connected)

  //run server forever
  server.Server(conf)
  http.HandleFunc("/", sayHello)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
