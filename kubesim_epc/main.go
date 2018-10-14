package main
import (
  "log"
  "net/http"
  "strings"
  "gopkg.in/yaml.v2"
  "io/ioutil"
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


  //this example is working, from http://sweetohm.net/article/go-yaml-parsers.en.html
  //previous example(not working):  https://stackoverflow.com/questions/30947534/reading-a-yaml-file-in-golang

  type Config struct {
      Foo string
      Bar[] string
  }

  yamlFile, err := ioutil.ReadFile("config.yaml")
  log.Printf("yamlFile=%s", string(yamlFile))
  if err != nil {
      log.Printf("yamlFile.Get err   #%v ", err)
  }

  var config Config
  err = yaml.Unmarshal(yamlFile, &config)
  if err != nil {
      log.Fatalf("Unmarshal: %v", err)
  }
  log.Printf("Foo=%s, Bar=%s", config.Foo, config.Bar)

  //run server forever
  server.Server()
  http.HandleFunc("/", sayHello)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
