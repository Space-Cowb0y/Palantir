package plugin
import(
  "context"; "encoding/json"; "errors"; "fmt"; "net"; "os"; "os/exec"; "path/filepath"; "sync"; "time"
  "github.com/Space-Cowb0y/Palantir/sentinel/internal/config"
  "github.com/Space-Cowb0y/Palantir/sentinel/internal/logging"
)
type Manifest struct{ Name string `json:"name"`; Exec string `json:"exec"`; Args []string `json:"args"`; Proto string `json:"proto"`; Health string `json:"health"` }
type Eye struct{ Manifest Manifest; Cmd *exec.Cmd; Addr string; Status string; PID int; Started time.Time }
type Manager struct{ cfg *config.Config; log logging.Logger; mu sync.RWMutex; eyes map[string]*Eye }
func NewManager(cfg *config.Config, log logging.Logger)*Manager{ return &Manager{cfg:cfg,log:log,eyes:map[string]*Eye{}} }
func (m *Manager) Watch(ctx context.Context){ m.log.Info("Plugin watcher started","dir",m.cfg.Plugins.Dir); t:=time.NewTicker(time.Duration(m.cfg.Plugins.PollSecond)*time.Second); defer t.Stop(); for{ select{ case <-ctx.Done(): m.log.Info("Plugin watcher stopped"); return; case <-t.C: _=m.scan() } } }
func (m *Manager) scan() error{ entries,err:=os.ReadDir(m.cfg.Plugins.Dir); if err!=nil{return err}; for _,e:= range entries { if e.IsDir(){continue}; if filepath.Ext(e.Name())!=".json"{continue}; b,err:=os.ReadFile(filepath.Join(m.cfg.Plugins.Dir,e.Name())); if err!=nil{continue}; var mf Manifest; if err:=json.Unmarshal(b,&mf); err!=nil{continue}; if _,ok:=m.eyes[mf.Name]; ok{continue}; if m.cfg.Plugins.Autostart{ _=m.startEye(mf) } }; return nil }
func freePort()(string,error){ l,err:=net.Listen("tcp","127.0.0.1:0"); if err!=nil{return "",err}; defer l.Close(); return l.Addr().String(),nil }
func (m *Manager) startEye(mf Manifest) error{ addr,err:=freePort(); if err!=nil{return err}; args:=append(mf.Args,fmt.Sprintf("--listen=%s",addr)); cmd:=exec.Command(mf.Exec,args...); cmd.Stdout,cmd.Stderr=os.Stdout,os.Stderr; if err:=cmd.Start(); err!=nil{return err}; m.mu.Lock(); m.eyes[mf.Name]=&Eye{Manifest:mf,Cmd:cmd,Addr:addr,Status:"starting",PID:cmd.Process.Pid,Started:time.Now()}; m.mu.Unlock(); m.log.Info("Eye started","name",mf.Name,"pid",cmd.Process.Pid,"addr",addr); return nil }
func (m *Manager) List() []*Eye{ m.mu.RLock(); defer m.mu.RUnlock(); out:=make([]*Eye,0,len(m.eyes)); for _,e:=range m.eyes{ out=append(out,e)}; return out }
func (m *Manager) Stop(name string) error{ m.mu.Lock(); defer m.mu.Unlock(); e,ok:=m.eyes[name]; if !ok{ return errors.New("not found") }; if e.Cmd!=nil && e.Cmd.Process!=nil { return e.Cmd.Process.Kill() }; return nil }