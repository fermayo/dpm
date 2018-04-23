package alias

import (
	"github.com/fermayo/dpm/parser"
)

type aliasSetter bool

const (
	set   aliasSetter = true
	unset aliasSetter = false
)

const (
	binaryLocation = "/usr/local/bin" // TODO: make variable for Windows
	bashFile       = "#!/bin/bash\nif [ -z ${DPM_ACTIVE+x} ]; then exec %s/%s-home \"$@\"; else exec \"$DPM_ACTIVE/%s\" \"$@\"; fi"
)

// setOrUnsetAliases loops the commands
func setOrUnsetAliases(aliases map[string]parser.Command, setter aliasSetter) error {
	for alias := range aliases {
		err := setOrUnsetAlias(alias, setter)
		if err != nil {
			return err
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
