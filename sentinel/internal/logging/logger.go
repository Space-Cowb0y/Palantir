package logging
import("log/slog"; "os"; "strings")
type Logger = *slog.Logger
func New(level string) Logger{ lvl:=slog.LevelInfo; switch strings.ToLower(level){case "debug":lvl=slog.LevelDebug; case "warn":lvl=slog.LevelWarn; case "error":lvl=slog.LevelError}; h:=slog.NewTextHandler(os.Stdout,&slog.HandlerOptions{Level:lvl}); return slog.New(h) }