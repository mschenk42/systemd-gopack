package systemd

import (
	"github.com/mschenk42/gopack"
)

// Run initializes the properties and runs the pack
func Run(props *gopack.Properties, actions []string) {
	pack := gopack.Pack{
		Name: "systemd",
		Props: &gopack.Properties{
			"systemd.prop1": "value",
		},
		Actions: actions,
		ActionMap: map[string]func(p *gopack.Pack){
			"default": run,
		},
	}
	pack.Run(props)
}

func run(pack *gopack.Pack) {
}
