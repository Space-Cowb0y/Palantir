package cmd
import(
    "fmt"; "os"; "github.com/spf13/cobra"
)
var rootCmd = &cobra.Command{Use:"sentinel",Short:"Sentinel — plugin manager & control-plane for Eyes"}
func Execute(){ if err:=rootCmd.Execute(); err!=nil{ fmt.Println(err); os.Exit(1)} }
func init(){ rootCmd.PersistentFlags().StringP("config","c","sentinel.yaml","config file path") }