package builder

// maybe not the best name
type BuildArguments struct {
	lang  string
	techs []string
}

func (ba *BuildArguments) SetLang(l string) {
	ba.lang = l
}

func (ba *BuildArguments) GetLang() string {
	return ba.lang
}

func (ba *BuildArguments) SetTechs(techs []string) {
	if techs != nil {
		ba.techs = append(ba.techs, techs...)
	}
}

func (ba *BuildArguments) GetTechs() []string {
	return ba.techs
}
