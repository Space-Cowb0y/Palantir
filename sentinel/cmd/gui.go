package cmd

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Space-Cowb0y/Palantir/sentinel/internal/config"
	"github.com/Space-Cowb0y/Palantir/sentinel/internal/logging"
	"github.com/Space-Cowb0y/Palantir/sentinel/internal/plugin"
	ui "github.com/Space-Cowb0y/Palantir/sentinel/pkg/ui"
	web "github.com/Space-Cowb0y/Palantir/sentinel/pkg/web"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run sentinel services (gRPC + HTTP) and watch for Eyes",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfgPath, _ := cmd.Flags().GetString("config")
		cfg, err := config.Load(cfgPath)
		if err != nil { return err }

		log := logging.New(cfg.LogLevel)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Start HTTP server (serves /healthz and static webui)
		httpSrv := web.NewHTTPServer(cfg, log)
		go func(){ _ = httpSrv.Start() }()

		// Start gRPC server (for Eyes)
		grpcSrv := web.NewGRPCServer(cfg, log)
		go func(){ _ = grpcSrv.Start() }()

		// Start plugin watcher/loader
		mgr := plugin.NewManager(cfg, log)
		go mgr.Watch(ctx)

		log.Info("Sentinel running", "http", fmt.Sprintf("http://%s", httpSrv.Addr()), "grpc", grpcSrv.Addr())
		<-ctx.Done()
		return nil
	},
}

var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "Open the TUI (Bubble Tea) dashboard",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load("sentinel.yaml"); if err != nil { return err }
		// log := logging.New(cfg.LogLevel)

		model := ui.NewManagerModel(func() ([]ui.EyeRow, error) {
			// Poll known Eyes from plugin manager’s registry
			return ui.QueryEyesOverHTTP(cfg.HTTP.Listen), nil
		})
		return ui.Run(model)
	},
}

var pluginsCmd = &cobra.Command{ Use: "plugins", Short: "Manage/list Eyes" }
var pluginsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List discovered Eyes",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, _ := config.Load("sentinel.yaml")
		resp, err := http.Get("http://" + cfg.HTTP.Listen + "/api/eyes")
		if err != nil { return err }
		defer resp.Body.Close()
		fmt.Println("HTTP:", resp.Status)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd, uiCmd, pluginsCmd)
	pluginsCmd.AddCommand(pluginsListCmd)

	// small delay to help Windows console attach
	time.Sleep(50 * time.Millisecond)
}