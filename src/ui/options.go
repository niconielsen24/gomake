package ui

type Options map[int]string

var defaultOptions = Options{
	0: "Go     : echo, PostgreSQL",
	1: "Go     : echo, MySQL",
	2: "Python : fastapi, PostgreSQL",
	3: "Python : fastapi, MySQL",
	4: "Python",
	5: "Go",
	6: "Zig",
	7: "Custom project",
}

var goOptions = Options{
	0: "Echo",
	1: "MySQL",
	2: "PostgreSQL",
  3: "Bubbletea",
  4: "Bubbles",
  5: "Lipgloss",
}

var pythonOptions = Options{
	0: "Flask",
	1: "Django",
	2: "FastAPI",
	3: "Pandas",
	4: "NumPy",
	5: "Scikit-learn",
	6: "Matplotlib",
	7: "Celery",
  8: "MySQL",
  9: "PostgreSQL",
}

var zigOptions = Options{
	0: "std.io",
	1: "std.json",
	2: "std.fs",
	3: "std.math",
	4: "std.testing",
	5: "async/await",
}

var langOptions = map[string]Options{
	"go": goOptions,
  "python" : pythonOptions,
  "zig" : zigOptions,
}

func GetDefaultOptions() Options {
	return defaultOptions
}

func GetLangOptions(s string) Options {
	return langOptions[s]
}
