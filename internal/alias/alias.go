package alias

import (
	"github.com/JPZ13/dpm/internal/parser"
)

type aliasSetter bool

const (
	set   aliasSetter = true
	unset aliasSetter = false
)

const (
	binaryLocation     = "/usr/local/bin" // TODO: make variable for Windows
	bashFile           = "#!/bin/bash\nif [ -z ${DPM_ACTIVE+x} ]; then exec %s/%s-home \"$@\"; else exec \"$DPM_ACTIVE/%s\" \"$@\"; fi"
	bashFileIfNotExist = "if [ -z ${DPM_ACTIVE+x} ]; then exec %s/%s-home \"$@\"; else echo 'Error: command %s not found'; exit 1; fi"
)

// setOrUnsetAliases loops the entrypoints for
// each package and runs set or unset
func setOrUnsetAliases(packages map[string]parser.Command, setter aliasSetter) error {
	for _, pkg := range packages {
		for _, entrypoint := range pkg.Entrypoints {
			err := setOrUnsetAlias(entrypoint, setter)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// setOrUnsetAlias invokes setAlias or unsetAlias
// depending on value passed to setter
func setOrUnsetAlias(alias string, setter aliasSetter) error {
	if setter == set {
		return setAlias(alias)
	}

	return unsetAlias(alias)
}
