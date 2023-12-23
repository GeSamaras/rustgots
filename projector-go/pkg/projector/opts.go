package projector

import "github.com/hellflame/argparse"

type Opts struct {
	Args   []string
	Config string
	Pwd    string
}

func GetOpts() (*Opts, error) {
	parser := argparse.NewParser("projector", "A tool for managing project templates",
	&argparse.ParserConfig{
	DisableDefaultShowHelp: true,
	})
	args := parser.Strings("a", "args", &argparse.Option{
		Help: "Arguments to pass to the template",
		Positional: true,
		Required: false,
		Default: "",
	})

	config := parser.String("c", "config", &argparse.Option{
		Help: "Path to the config file",
		Required: false,
		Default: "",
	})
	
	pwd := parser.String("p", "pwd", &argparse.Option{
		Help: "Path to the config file",
		Required: false,
		Default: "",
	})
	err := parser.Parse(nil)
	if err != nil {
		return nil, err
	}
	return &Opts{
		Args:   *args,
		Config: *config,
		Pwd:    *pwd,
	}, nil
}