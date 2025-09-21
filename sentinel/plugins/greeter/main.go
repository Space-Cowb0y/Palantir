package main
import(
  "flag"; "fmt"; "net"; "net/http"; "time"; "google.golang.org/grpc"
)
var listen=flag.String("listen","127.0.0.1:0","listen address")
func main(){ flag.Parse(); go func(){ _=http.ListenAndServe("127.0.0.1:0", http.NewServeMux()) }(); lis,err:=net.Listen("tcp",*listen); if err!=nil{panic(err)}; srv:=grpc.NewServer(); fmt.Println("greeter eye at", lis.Addr()); time.Sleep(10*time.Minute); _=srv.Serve(lis) }