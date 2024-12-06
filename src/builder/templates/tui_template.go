package templates

import (
  "gomake/builder"
  "github.com/google/uuid"
)


var TuiConfig = builder.BuilderConfig{
  Lang: "",
  DirTree: app, 
  GitRepo: builder.INIT_GIT_REPO,
  Techs: []string{"bubbletea","lipgloss","bubbles"},
  ProjectName: uuid.NewString(),
}


var app = builder.DirTree{
	DirName:    "app",
	InnerDirs:  []builder.DirTree{src},
	InnerFiles: []string{},
}

var src = builder.DirTree{
	DirName:    "src",
	InnerDirs:  []builder.DirTree{ui},
	InnerFiles: []string{},
}

var ui = builder.DirTree{
	DirName:    "ui",
	InnerDirs:  []builder.DirTree{models},
	InnerFiles: []string{"initial.go"},
}

var models = builder.DirTree{
	DirName:    "models",
	InnerDirs:  []builder.DirTree{},
	InnerFiles: []string{},
}
