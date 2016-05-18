package main

import (
	"os"
	"sync"

	"github.com/cenk/backoff"
	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/index/clg"
	"github.com/xh3b4sd/anna/index/clg/collection"
	"github.com/xh3b4sd/anna/index/clg/profile"
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
	"github.com/xh3b4sd/anna/scheduler"
	"github.com/xh3b4sd/anna/server"
	"github.com/xh3b4sd/anna/server/control/log"
	"github.com/xh3b4sd/anna/server/interface/text"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
	"github.com/xh3b4sd/anna/storage/redis"
)

const (
	objectTypeAnna spec.ObjectType = "anna"
)

var (
	globalFlags struct {
		Addr string

		ControlLogLevels    string
		ControlLogObejcts   string
		ControlLogVerbosity int

		Storage     string
		StorageAddr string
	}

	annaCmd = &cobra.Command{
		Use:   "anna",
		Short: "Anna, Artificial Neural Network Aspiration, aims to be self-learning and self-improving software. For more information see https://github.com/xh3b4sd/anna.",
		Long:  "Anna, Artificial Neural Network Aspiration, aims to be self-learning and self-improving software. For more information see https://github.com/xh3b4sd/anna.",
		Run:   mainRun,
	}

	// version is the project version. It is given via buildflags that inject the
	// commit hash.
	version string
)

func init() {
	annaCmd.PersistentFlags().StringVar(&globalFlags.Addr, "addr", "127.0.0.1:9119", "host:port to bind Anna's server to")

	annaCmd.PersistentFlags().StringVar(&globalFlags.ControlLogLevels, "control-log-levels", "", "set log levels for log control (e.g. E,F)")
	annaCmd.PersistentFlags().StringVar(&globalFlags.ControlLogObejcts, "control-log-objects", "", "set log objects for log control (e.g. core-net,impulse)")
	annaCmd.PersistentFlags().IntVar(&globalFlags.ControlLogVerbosity, "control-log-verbosity", 10, "set log verbosity for log control")

	annaCmd.PersistentFlags().StringVar(&globalFlags.Storage, "storage", "redis", "storage type to use for persistency (e.g. memory)")
	annaCmd.PersistentFlags().StringVar(&globalFlags.StorageAddr, "storage-addr", "127.0.0.1:6379", "host:port to connect to storage")
}

type annaConfig struct {
	CLGIndex spec.CLGIndex
	CoreNet  spec.Network
	Log      spec.Log
	Server   spec.Server
	Storage  spec.Storage
}

func defaultAnnaConfig() annaConfig {
	newCLGIndex, err := clg.NewIndex(clg.DefaultIndexConfig())
	if err != nil {
		panic(err)
	}

	newConfig := annaConfig{
		CLGIndex: newCLGIndex,
		CoreNet:  nil,
		Log:      log.NewLog(log.DefaultConfig()),
		Server:   nil,
		Storage:  memorystorage.NewMemoryStorage(memorystorage.DefaultConfig()),
	}

	return newConfig
}

func newAnna(config annaConfig) (spec.Anna, error) {
	newAnna := &anna{
		annaConfig: config,

		BootOnce:     sync.Once{},
		ID:           id.NewObjectID(id.Hex128),
		Mutex:        sync.Mutex{},
		ShutdownOnce: sync.Once{},
		Type:         spec.ObjectType(objectTypeAnna),
	}

	if newAnna.CLGIndex == nil {
		return nil, maskAnyf(invalidConfigError, "CLG index must not be empty")
	}
	if newAnna.CoreNet == nil {
		return nil, maskAnyf(invalidConfigError, "core network must not be empty")
	}
	if newAnna.Log == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
	}
	if newAnna.Server == nil {
		return nil, maskAnyf(invalidConfigError, "server must not be empty")
	}
	if newAnna.Storage == nil {
		return nil, maskAnyf(invalidConfigError, "storage must not be empty")
	}

	newAnna.Log.Register(newAnna.GetType())

	return newAnna, nil
}

// mainObject is basically only to have an object that provides proper
// identifyable logging.
type anna struct {
	annaConfig

	BootOnce     sync.Once
	ID           spec.ObjectID
	Mutex        sync.Mutex
	ShutdownOnce sync.Once
	Type         spec.ObjectType
}

func (a *anna) Boot() {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call Boot")

	a.BootOnce.Do(func() {
		a.Log.WithTags(spec.Tags{L: "I", O: a, T: nil, V: 10}, "hello, I am Anna")

		go a.listenToSignal()
		go a.writeStateInfo()

		a.Log.WithTags(spec.Tags{L: "I", O: a, T: nil, V: 10}, "booting CLG index")
		go a.CLGIndex.Boot()

		a.Log.WithTags(spec.Tags{L: "I", O: a, T: nil, V: 10}, "booting core-net")
		go a.CoreNet.Boot()

		a.Log.WithTags(spec.Tags{L: "I", O: a, T: nil, V: 10}, "booting server")
		a.Server.Boot()
	})
}

func (a *anna) Shutdown() {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call Shutdown")

	a.ShutdownOnce.Do(func() {
		go a.CLGIndex.Shutdown()
		go a.CoreNet.Shutdown()

		a.Log.WithTags(spec.Tags{L: "I", O: a, T: nil, V: 10}, "shutting down")
		os.Exit(0)
	})
}

func mainRun(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		cmd.Help()
		os.Exit(1)
	}

	var err error

	// log
	newLog := log.NewLog(log.DefaultConfig())
	err = newLog.SetLevels(globalFlags.ControlLogLevels)
	panicOnError(err)
	err = newLog.SetObjects(globalFlags.ControlLogObejcts)
	panicOnError(err)
	err = newLog.SetVerbosity(globalFlags.ControlLogVerbosity)
	panicOnError(err)

	// text gateway
	newTextGatewayConfig := gateway.DefaultConfig()
	newTextGatewayConfig.Log = newLog
	newTextGateway := gateway.NewGateway(newTextGatewayConfig)

	// storage
	var newStorage spec.Storage
	switch globalFlags.Storage {
	case "redis":
		newRedisDialConfig := redisstorage.DefaultRedisDialConfig()
		newRedisDialConfig.Addr = globalFlags.StorageAddr
		newPoolConfig := redisstorage.DefaultRedisPoolConfig()
		newPoolConfig.Dial = redisstorage.NewRedisDial(newRedisDialConfig)
		newStorageConfig := redisstorage.DefaultConfig()
		newStorageConfig.BackOffFactory = func() spec.BackOff {
			return backoff.NewExponentialBackOff()
		}
		newStorageConfig.Log = newLog
		newStorageConfig.Pool = redisstorage.NewRedisPool(newPoolConfig)
		newStorage, err = redisstorage.NewRedisStorage(newStorageConfig)
		panicOnError(err)
	case "memory":
		newStorage = memorystorage.NewMemoryStorage(memorystorage.DefaultConfig())
	default:
		panic("invalid storage flag")
	}

	// scheduler
	newSchedulerConfig := scheduler.DefaultConfig()
	newSchedulerConfig.Log = newLog
	newSchedulerConfig.Storage = newStorage
	newScheduler, err := scheduler.NewScheduler(newSchedulerConfig)
	panicOnError(err)

	// pattern network
	newPatNetConfig := patnet.DefaultConfig()
	newPatNetConfig.Log = newLog
	newPatNetConfig.Storage = newStorage
	newPatNet, err := patnet.NewPatNet(newPatNetConfig)
	panicOnError(err)

	// strategy network
	newStratNetConfig := stratnet.DefaultConfig()
	newStratNetConfig.Log = newLog
	newStratNetConfig.PatNet = newPatNet
	newStratNetConfig.Storage = newStorage
	newStratNet, err := stratnet.NewStratNet(newStratNetConfig)
	panicOnError(err)

	// prediction network
	newPredNetConfig := prednet.DefaultConfig()
	newPredNetConfig.Log = newLog
	newStratNetConfig.PatNet = newPatNet
	newPredNetConfig.Storage = newStorage
	newPredNet, err := prednet.NewPredNet(newPredNetConfig)
	panicOnError(err)

	// evaluation network
	newEvalNetConfig := evalnet.DefaultConfig()
	newEvalNetConfig.Log = newLog
	newStratNetConfig.PatNet = newPatNet
	newEvalNetConfig.Storage = newStorage
	newEvalNet, err := evalnet.NewEvalNet(newEvalNetConfig)
	panicOnError(err)

	// character network
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

	// context network
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

	// idea network
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

	// response network
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

	// input network
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

	// output network
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

	// core network
	newCoreExecNetConfig := coreexecnet.DefaultConfig()
	newCoreExecNetConfig.InNet = newInNet
	newCoreExecNetConfig.Log = newLog
	newCoreExecNetConfig.OutNet = newOutNet
	newCoreExecNet, err := coreexecnet.NewExecNet(newCoreExecNetConfig)
	panicOnError(err)
	newCoreNetConfig := corenet.DefaultConfig()
	newCoreNetConfig.Log = newLog
	newCoreNetConfig.EvalNet = newEvalNet
	newCoreNetConfig.ExecNet = newCoreExecNet
	newCoreNetConfig.PatNet = newPatNet
	newCoreNetConfig.PredNet = newPredNet
	newCoreNetConfig.Scheduler = newScheduler
	newCoreNetConfig.StratNet = newStratNet
	newCoreNetConfig.TextGateway = newTextGateway
	newCoreNet, err := corenet.NewCoreNet(newCoreNetConfig)
	panicOnError(err)

	// CLG collection
	newCollectionConfig := collection.DefaultConfig()
	newCollectionConfig.Log = newLog
	newCollection, err := collection.New(newCollectionConfig)
	panicOnError(err)

	// CLG profile generator
	newGeneratorConfig := profile.DefaultGeneratorConfig()
	newGeneratorConfig.Collection = newCollection
	newGeneratorConfig.Log = newLog
	newGeneratorConfig.Storage = newStorage
	newGenerator, err := profile.NewGenerator(newGeneratorConfig)
	panicOnError(err)

	// CLG index
	newCLGConfig := clg.DefaultIndexConfig()
	newCLGConfig.Generator = newGenerator
	newCLGConfig.Log = newLog
	newCLGIndex, err := clg.NewIndex(newCLGConfig)
	panicOnError(err)

	// log control
	newLogControlConfig := logcontrol.DefaultConfig()
	newLogControlConfig.Log = newLog
	newLogControl := logcontrol.NewLogControl(newLogControlConfig)

	// text interface
	newTextInterfaceConfig := textinterface.DefaultConfig()
	newTextInterfaceConfig.Log = newLog
	newTextInterfaceConfig.Scheduler = newScheduler
	newTextInterfaceConfig.TextGateway = newTextGateway
	newTextInterface, err := textinterface.NewTextInterface(newTextInterfaceConfig)
	panicOnError(err)

	// server
	newServerConfig := server.DefaultConfig()
	newServerConfig.Addr = globalFlags.Addr
	newServerConfig.Log = newLog
	newServerConfig.LogControl = newLogControl
	newServerConfig.TextGateway = newTextGateway
	newServerConfig.TextInterface = newTextInterface
	newServer, err := server.NewServer(newServerConfig)
	panicOnError(err)

	// anna
	newAnnaConfig := defaultAnnaConfig()
	newAnnaConfig.CLGIndex = newCLGIndex
	newAnnaConfig.CoreNet = newCoreNet
	newAnnaConfig.Log = newLog
	newAnnaConfig.Server = newServer
	newAnnaConfig.Storage = newStorage
	a, err := newAnna(newAnnaConfig)
	panicOnError(err)

	a.Boot()
}

func main() {
	annaCmd.AddCommand(versionCmd)

	annaCmd.Execute()
}
