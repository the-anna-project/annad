package main

import (
	"os"
	"sync"

	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/factory/server"
	"github.com/xh3b4sd/anna/file-system/os"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/net/char"
	charexecnet "github.com/xh3b4sd/anna/net/char/net/exec"
	"github.com/xh3b4sd/anna/net/core"
	coreexecnet "github.com/xh3b4sd/anna/net/core/net/exec"
	"github.com/xh3b4sd/anna/net/ctx"
	ctxexecnet "github.com/xh3b4sd/anna/net/ctx/net/exec"
	"github.com/xh3b4sd/anna/net/eval"
	"github.com/xh3b4sd/anna/net/idea"
	ideaexecnet "github.com/xh3b4sd/anna/net/idea/net/exec"
	"github.com/xh3b4sd/anna/net/in"
	inexecnet "github.com/xh3b4sd/anna/net/in/net/exec"
	"github.com/xh3b4sd/anna/net/out"
	outexecnet "github.com/xh3b4sd/anna/net/out/net/exec"
	"github.com/xh3b4sd/anna/net/pat"
	"github.com/xh3b4sd/anna/net/pred"
	"github.com/xh3b4sd/anna/net/resp"
	respexecnet "github.com/xh3b4sd/anna/net/resp/net/exec"
	"github.com/xh3b4sd/anna/net/strat"
	"github.com/xh3b4sd/anna/server"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/redis"
)

const (
	objectTypeAnna spec.ObjectType = "anna"
)

var (
	globalFlags struct {
		ControlLogLevels    string
		ControlLogObejcts   string
		ControlLogVerbosity int

		Addr string
	}

	annaCmd = &cobra.Command{
		Use:   "anna",
		Short: "Anna, Artificial Neural Network Aspiration, aims to be self-learning and self-improving software. For more information see https://github.com/xh3b4sd/anna.",
		Long:  "Anna, Artificial Neural Network Aspiration, aims to be self-learning and self-improving software. For more information see https://github.com/xh3b4sd/anna.",
		Run:   mainRun,
	}

	// Version is the project version. It is given via buildflags that inject the
	// commit hash.
	version string
)

func init() {
	annaCmd.PersistentFlags().StringVar(&globalFlags.ControlLogLevels, "control-log-levels", "", "set log levels for log control (e.g. E,F)")
	annaCmd.PersistentFlags().StringVar(&globalFlags.ControlLogObejcts, "control-log-objects", "", "set log objects for log control (e.g. core-net,impulse)")
	annaCmd.PersistentFlags().IntVar(&globalFlags.ControlLogVerbosity, "control-log-verbosity", 10, "set log verbosity for log control")

	annaCmd.PersistentFlags().StringVar(&globalFlags.Addr, "addr", "127.0.0.1:9119", "host:port to bind Anna's server to")
}

type annaConfig struct {
	CoreNet spec.Network

	FactoryServer spec.Factory

	Log spec.Log

	Server spec.Server
}

func defaultAnnaConfig() annaConfig {
	newConfig := annaConfig{
		CoreNet:       nil,
		FactoryServer: factoryserver.NewFactory(factoryserver.DefaultConfig()),
		Log:           log.NewLog(log.DefaultConfig()),
		Server:        server.NewServer(server.DefaultConfig()),
	}

	return newConfig
}

func newAnna(config annaConfig) spec.Anna {
	newAnna := &anna{
		annaConfig: config,
		Booted:     false,
		ID:         id.NewObjectID(id.Hex128),
		Mutex:      sync.Mutex{},
		Type:       spec.ObjectType(objectTypeAnna),
	}

	newAnna.Log.Register(newAnna.GetType())

	return newAnna
}

// mainObject is basically only to have an object that provides proper
// identifyable logging.
type anna struct {
	annaConfig

	Booted bool
	ID     spec.ObjectID
	Mutex  sync.Mutex
	Type   spec.ObjectType
}

func (a *anna) Boot() {
	a.Mutex.Lock()
	defer a.Mutex.Unlock()

	if a.Booted {
		return
	}
	a.Booted = true

	a.Log.WithTags(spec.Tags{L: "I", O: a, T: nil, V: 10}, "hello, I am Anna")

	go a.listenToSignal()

	a.Log.WithTags(spec.Tags{L: "I", O: a, T: nil, V: 10}, "booting factory")
	go a.FactoryServer.Boot()

	a.Log.WithTags(spec.Tags{L: "I", O: a, T: nil, V: 10}, "booting core-net")
	go a.CoreNet.Boot()

	a.Log.WithTags(spec.Tags{L: "I", O: a, T: nil, V: 10}, "booting server")
	a.Server.Boot()
}

func (a *anna) Shutdown() {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call Shutdown")

	go a.CoreNet.Shutdown()
	go a.FactoryServer.Shutdown()

	a.Log.WithTags(spec.Tags{L: "I", O: a, T: nil, V: 10}, "shutting down")
	os.Exit(0)
}

func mainRun(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		cmd.Help()
		os.Exit(1)
	}

	var err error

	// File system
	newOSFileSystem := osfilesystem.NewFileSystem()

	// Log
	newLog := log.NewLog(log.DefaultConfig())
	err = newLog.SetLevels(globalFlags.ControlLogLevels)
	panicOnError(err)
	err = newLog.SetObjects(globalFlags.ControlLogObejcts)
	panicOnError(err)
	err = newLog.SetVerbosity(globalFlags.ControlLogVerbosity)
	panicOnError(err)

	// Factory gateway
	newFactoryGatewayConfig := gateway.DefaultConfig()
	newFactoryGatewayConfig.Log = newLog
	newFactoryGateway := gateway.NewGateway(newFactoryGatewayConfig)

	// Text gateway
	newTextGatewayConfig := gateway.DefaultConfig()
	newTextGatewayConfig.Log = newLog
	newTextGateway := gateway.NewGateway(newTextGatewayConfig)

	// Factory
	newFactoryClientConfig := factoryclient.DefaultConfig()
	newFactoryClientConfig.FactoryGateway = newFactoryGateway
	newFactoryClientConfig.Log = newLog
	newFactoryGatewayClient := factoryclient.NewFactory(newFactoryClientConfig)
	newFactoryServerConfig := factoryserver.DefaultConfig()
	newFactoryServerConfig.FactoryClient = newFactoryGatewayClient
	newFactoryServerConfig.FactoryGateway = newFactoryGateway
	newFactoryServerConfig.FileSystem = newOSFileSystem
	newFactoryServerConfig.Log = newLog
	newFactoryServerConfig.TextGateway = newTextGateway
	newFactoryServer := factoryserver.NewFactory(newFactoryServerConfig)

	// Storage
	newRedisDialConfig := redisstorage.DefaultRedisDialConfig()
	newRedisDialConfig.Addr = "127.0.0.1:6379"
	newPoolConfig := redisstorage.DefaultRedisPoolConfig()
	newPoolConfig.Dial = redisstorage.NewRedisDial(newRedisDialConfig)
	newStorageConfig := redisstorage.DefaultConfig()
	newStorageConfig.Log = newLog
	newStorageConfig.Pool = redisstorage.NewRedisPool(newPoolConfig)
	newStorage := redisstorage.NewRedisStorage(newStorageConfig)

	// Pattern network
	newPatNetConfig := patnet.DefaultConfig()
	newPatNetConfig.Log = newLog
	newPatNetConfig.Storage = newStorage
	newPatNet, err := patnet.NewPatNet(newPatNetConfig)
	panicOnError(err)

	// Strategy network
	newStratNetConfig := stratnet.DefaultConfig()
	newStratNetConfig.Log = newLog
	newStratNetConfig.PatNet = newPatNet
	newStratNetConfig.Storage = newStorage
	newStratNet, err := stratnet.NewStratNet(newStratNetConfig)
	panicOnError(err)

	// Prediction network
	newPredNetConfig := prednet.DefaultConfig()
	newPredNetConfig.Log = newLog
	newStratNetConfig.PatNet = newPatNet
	newPredNetConfig.Storage = newStorage
	newPredNet, err := prednet.NewPredNet(newPredNetConfig)
	panicOnError(err)

	// Evaluation network
	newEvalNetConfig := evalnet.DefaultConfig()
	newEvalNetConfig.Log = newLog
	newStratNetConfig.PatNet = newPatNet
	newEvalNetConfig.Storage = newStorage
	newEvalNet, err := evalnet.NewEvalNet(newEvalNetConfig)
	panicOnError(err)

	// Character network
	newCharExecNetConfig := charexecnet.DefaultConfig()
	newCharExecNetConfig.Log = newLog
	newCharExecNet, err := charexecnet.NewExecNet(newCharExecNetConfig)
	panicOnError(err)
	newCharNetConfig := charnet.DefaultConfig()
	newCharNetConfig.Log = newLog
	newCharNetConfig.EvalNet = newEvalNet
	newCharNetConfig.ExecNet = newCharExecNet
	newCharNetConfig.PatNet = newPatNet
	newCharNetConfig.PredNet = newPredNet
	newCharNetConfig.StratNet = newStratNet
	newCharNetConfig.Storage = newStorage
	newCharNet, err := charnet.NewCharNet(newCharNetConfig)
	panicOnError(err)

	// Context network
	newCtxExecNetConfig := ctxexecnet.DefaultConfig()
	newCtxExecNetConfig.Log = newLog
	newCtxExecNet, err := ctxexecnet.NewExecNet(newCtxExecNetConfig)
	panicOnError(err)
	newCtxNetConfig := ctxnet.DefaultConfig()
	newCtxNetConfig.Log = newLog
	newCtxNetConfig.EvalNet = newEvalNet
	newCtxNetConfig.ExecNet = newCtxExecNet
	newCtxNetConfig.PatNet = newPatNet
	newCtxNetConfig.PredNet = newPredNet
	newCtxNetConfig.StratNet = newStratNet
	newCtxNetConfig.Storage = newStorage
	newCtxNet, err := ctxnet.NewCtxNet(newCtxNetConfig)
	panicOnError(err)

	// Idea network
	newIdeaExecNetConfig := ideaexecnet.DefaultConfig()
	newIdeaExecNetConfig.Log = newLog
	newIdeaExecNet, err := ideaexecnet.NewExecNet(newIdeaExecNetConfig)
	panicOnError(err)
	newIdeaNetConfig := ideanet.DefaultConfig()
	newIdeaNetConfig.Log = newLog
	newIdeaNetConfig.EvalNet = newEvalNet
	newIdeaNetConfig.ExecNet = newIdeaExecNet
	newIdeaNetConfig.PatNet = newPatNet
	newIdeaNetConfig.PredNet = newPredNet
	newIdeaNetConfig.StratNet = newStratNet
	newIdeaNetConfig.Storage = newStorage
	newIdeaNet, err := ideanet.NewIdeaNet(newIdeaNetConfig)
	panicOnError(err)

	// Response network
	newRespExecNetConfig := respexecnet.DefaultConfig()
	newRespExecNetConfig.Log = newLog
	newRespExecNet, err := respexecnet.NewExecNet(newRespExecNetConfig)
	panicOnError(err)
	newRespNetConfig := respnet.DefaultConfig()
	newRespNetConfig.Log = newLog
	newRespNetConfig.EvalNet = newEvalNet
	newRespNetConfig.ExecNet = newRespExecNet
	newRespNetConfig.PatNet = newPatNet
	newRespNetConfig.PredNet = newPredNet
	newRespNetConfig.StratNet = newStratNet
	newRespNetConfig.Storage = newStorage
	newRespNet, err := respnet.NewRespNet(newRespNetConfig)
	panicOnError(err)

	// Input network
	newInExecNetConfig := inexecnet.DefaultConfig()
	newInExecNetConfig.CharNet = newCharNet
	newInExecNetConfig.CtxNet = newCtxNet
	newInExecNetConfig.Log = newLog
	newInExecNet, err := inexecnet.NewExecNet(newInExecNetConfig)
	panicOnError(err)
	newInNetConfig := innet.DefaultConfig()
	newInNetConfig.Log = newLog
	newInNetConfig.EvalNet = newEvalNet
	newInNetConfig.ExecNet = newInExecNet
	newInNetConfig.PatNet = newPatNet
	newInNetConfig.PredNet = newPredNet
	newInNetConfig.StratNet = newStratNet
	newInNetConfig.Storage = newStorage
	newInNet, err := innet.NewInNet(newInNetConfig)
	panicOnError(err)

	// Output network
	newOutExecNetConfig := outexecnet.DefaultConfig()
	newOutExecNetConfig.IdeaNet = newIdeaNet
	newOutExecNetConfig.Log = newLog
	newOutExecNetConfig.RespNet = newRespNet
	newOutExecNet, err := outexecnet.NewExecNet(newOutExecNetConfig)
	panicOnError(err)
	newOutNetConfig := outnet.DefaultConfig()
	newOutNetConfig.Log = newLog
	newOutNetConfig.EvalNet = newEvalNet
	newOutNetConfig.ExecNet = newOutExecNet
	newOutNetConfig.PatNet = newPatNet
	newOutNetConfig.PredNet = newPredNet
	newOutNetConfig.StratNet = newStratNet
	newOutNetConfig.Storage = newStorage
	newOutNet, err := outnet.NewOutNet(newOutNetConfig)
	panicOnError(err)

	//
	// Core network
	//
	newCoreExecNetConfig := coreexecnet.DefaultConfig()
	newCoreExecNetConfig.InNet = newInNet
	newCoreExecNetConfig.Log = newLog
	newCoreExecNetConfig.OutNet = newOutNet
	newCoreExecNet, err := coreexecnet.NewExecNet(newCoreExecNetConfig)
	panicOnError(err)
	newCoreNetConfig := corenet.DefaultConfig()
	newCoreNetConfig.FactoryClient = newFactoryGatewayClient
	newCoreNetConfig.Log = newLog
	newCoreNetConfig.EvalNet = newEvalNet
	newCoreNetConfig.ExecNet = newCoreExecNet
	newCoreNetConfig.PatNet = newPatNet
	newCoreNetConfig.PredNet = newPredNet
	newCoreNetConfig.StratNet = newStratNet
	newCoreNetConfig.TextGateway = newTextGateway
	newCoreNet, err := corenet.NewCoreNet(newCoreNetConfig)
	panicOnError(err)

	//
	// create server
	//
	newServerConfig := server.DefaultConfig()
	newServerConfig.Addr = globalFlags.Addr
	newServerConfig.Log = newLog
	newServerConfig.TextGateway = newTextGateway
	newServer := server.NewServer(newServerConfig)

	//
	// create anna
	//
	newAnnaConfig := defaultAnnaConfig()
	newAnnaConfig.CoreNet = newCoreNet
	newAnnaConfig.FactoryServer = newFactoryServer
	newAnnaConfig.Log = newLog
	newAnnaConfig.Server = newServer
	a := newAnna(newAnnaConfig)

	a.Boot()
}

func main() {
	annaCmd.AddCommand(versionCmd)

	annaCmd.Execute()
}
