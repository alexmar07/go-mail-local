package command

import "errors"

type Runnable interface {
	Run(args []string) error
}

type Kernel struct {
	cmds    map[string]Runnable
	options Options
}

func NewKernel(options Options) *Kernel {

	k := &Kernel{
		cmds:    make(map[string]Runnable),
		options: options,
	}

	k.boot()

	return k
}

func (k *Kernel) Run(cmd string, args []string) error {

	if runnable, ok := k.cmds[cmd]; ok {
		return runnable.Run(args)
	}

	return errors.New("command not found")
}

func (k *Kernel) Register(cmd string, runnable Runnable) {
	k.cmds[cmd] = runnable
}

func (k *Kernel) boot() {

	k.Register("list", &ListCommand{})
	k.Register("setup", &Setup{k.options})
	k.Register("sendmail", &Sender{k.options})
}
