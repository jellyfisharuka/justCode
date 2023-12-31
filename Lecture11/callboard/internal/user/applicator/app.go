package applicator

import (
	"callboard/internal/user/config"
	"callboard/internal/user/database"
	"callboard/internal/user/repository"
	"callboard/internal/user/server/http"
	"callboard/internal/user/user"
	"context"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)
type Applicator struct {
	logger *zap.SugaredLogger
	config *config.Config
}

func NewApplicator(logger *zap.SugaredLogger, config *config.Config) *Applicator {
	return &Applicator{
		logger: logger,
		config: config,
	}
}
func (a *Applicator) Run() {
	var (
		cfg = a.config
		l   = a.logger
	)

	ctx, cancel := context.WithCancel(context.TODO())
	_ = ctx

	mainDB, err := database.New(cfg.Database.Main)
	if err != nil {
		l.Fatalf("cannot сonnect to mainDB '%s:%d': %v", cfg.Database.Main.Host, cfg.Database.Main.Port, err)
	}

	defer func() {
		if err := mainDB.Close(); err != nil {
			l.Panicf("failed close MainDB err: %v", err)
		}
		l.Info("MainDB closed")
	}()

	replicaDB, err := database.New(cfg.Database.Replica)
	if err != nil {
		l.Fatalf("cannot сonnect to replicaDB '%s:%d': %v", cfg.Database.Replica.Host, cfg.Database.Replica.Port, err)
	}

	defer func() {
		if err := replicaDB.Close(); err != nil {
			l.Panicf("failed close replicaDB err: %v", err)
		}
		l.Info("replicaDB closed")
	}()

	repo := repository.NewRepository(mainDB, replicaDB)
	_ = repo

	userService := user.NewService(repo)

	endpointHandler := http.NewEndpointHandler(l, userService)

	router := http.NewRouter(l)
	httpCfg := cfg.HttpServer
	server, err := http.NewServer(httpCfg.Port, httpCfg.ShutdownTimeout, router, l, endpointHandler)
	if err != nil {
		l.Fatalf("failed to create server err: %v", err)
	}

	server.Run()
	defer func() {
		if err := server.Stop(); err != nil {
			l.Panicf("failed close server err: %v", err)
		}
		l.Info("server closed")
	}()

	a.gracefulShutdown(cancel)
}

func (a *Applicator) gracefulShutdown(cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch
	signal.Stop(ch)
	cancel()
}
