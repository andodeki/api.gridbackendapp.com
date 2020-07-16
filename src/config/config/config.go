package config

import "flag"

//DataDirectory - is the path used for loading tempaltes/database migrations
var DataDirectory = flag.String("data-directory", "", "Path for loading griffinss and migrations scripts")
