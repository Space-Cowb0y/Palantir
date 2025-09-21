package cmd
import(
    "context"; "fmt"; "time"; "net/http"
    "github.com/Space-Cowb0y/Palantir/sentinel/internal/config"
    "github.com/Space-Cowb0y/Palantir/sentinel/internal/logging"
    "github.com/Space-Cowb0y/Palantir/sentinel/internal/plugin"
    ui "github.com/Space-Cowb0y/Palantir/sentinel/pkg/ui"
    web "github.com/Space-Cowb0y/Palantir/sentinel/pkg/web"
    "github.com/spf13/cobra"
)
var runCmd=&cobra.Command{Use:"run",Short:"Run services",RunE:func(cmd *cobra.Command,args[]string)error{
    cfg,err:=config.Load("sentinel.yaml"); if err!=nil{return err}
    log:=logging.New(cfg.LogLevel)
    ctx,cancel:=context.WithCancel(context.Background()); defer cancel()
    httpSrv:=web.NewHTTPServer(cfg,log); go func(){ _=httpSrv.Start() }()
    grpcSrv:=web.NewGRPCServer(cfg,log); go func(){ _=grpcSrv.Start() }()
    mgr:=plugin.NewManager(cfg,log); go mgr.Watch(ctx)
    log.Info("Sentinel running","http",fmt.Sprintf("http://%s",httpSrv.Addr()),"grpc",grpcSrv.Addr())
    <-ctx.Done(); return nil}}
var uiCmd=&cobra.Command{Use:"ui",Short:"Open TUI",RunE:func(cmd *cobra.Command,args[]string)error{
    cfg,err:=config.Load("sentinel.yaml"); if err!=nil{return err}
    m:=ui.NewManagerModel(func()([]ui.EyeRow,error){return ui.QueryEyesOverHTTP(cfg.HTTP.Listen),nil})
    return ui.Run(m)}}
var pluginsCmd=&cobra.Command{Use:"plugins",Short:"Manage/list Eyes"}
var pluginsListCmd=&cobra.Command{Use:"list",Short:"List Eyes",RunE:func(cmd *cobra.Command,args[]string)error{
    cfg,_:=config.Load("sentinel.yaml")
    resp,err:=http.Get("http://"+cfg.HTTP.Listen+"/api/eyes"); if err!=nil{return err}
    defer resp.Body.Close(); fmt.Println("HTTP:",resp.Status); return nil}}
func init(){ rootCmd.AddCommand(runCmd,uiCmd,pluginsCmd); pluginsCmd.AddCommand(pluginsListCmd); time.Sleep(50*time.Millisecond)}