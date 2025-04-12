package config

type Logger struct {
	Environment string // dev, production
	Level       string // debug, info, warn, error
	Format      string // json, text
	OutputFile  string // nếu rỗng thì stdout
	AddSource   bool   // có log file, line, func không
}
