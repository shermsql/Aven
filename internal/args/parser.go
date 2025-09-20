package args

import "flag"

type Config struct {
	Mode      string
	Algorithm string
	Text      string
	Help      bool
}

func Parse() Config {
	mode := flag.String("m", "", "Mode: Encrypt Or Decrypt.")
	algorithm := flag.String("a", "", "Algorithm: Caesar Or XOR.")
	text := flag.String("t", "", "Text To Process.")
	help := flag.Bool("help", false, "Show Help Message.")

	flag.Parse()

	return Config {
		Mode:      *mode,
		Algorithm: *algorithm,
		Text:      *text,
		Help:      *help,
	}
}