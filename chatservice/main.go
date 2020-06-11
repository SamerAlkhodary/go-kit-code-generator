package chatService
import(
"github.com/go-kit/kit/log"
"github.com/go-kit/kit/log/level"
"fmt"
"flag"
"net/http"
"os"
"os/signal"
"syscall"
"context"
)
func main(){
var httpAddr= flag.String("http",":8080","http listen address")
var logger log.Logger
{


var httpAddr= flag.String("http",":8080","http listen address")
logger= log.NewLogfmtLogger(os.Stderr)
logger=log.NewSyncLogger(logger)
logger= log.With(logger,
"service","chatService",
"time", log.DefaultTimestampUTC,
"caller", log.DefaultCaller,
)
}
level.Info(logger).Log("msg","service started")
defer level.Info(logger).Log("msg","service ended")
flag.Parse()
ctx:=context.Background()
var service ChatService
{
service= NewService(logger)
}

errs:=make(chan error)
go func(){

c := make(chan os.Signal,1)
 signal.Notify(c,syscall.SIGINT, syscall.SIGTERM)
errs<- fmt.Errorf("%s",<-c)
}()
endpoints:=MakeEndpoints(service)
go func(){
fmt.Println("Listening on port",*httpAddr)
handler:=NewHTTPServer(ctx,endpoints)

}()
level.Error(logger).Log("exit", <-errs)
}