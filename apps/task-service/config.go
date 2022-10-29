package main

import "util"

var DATABASE_CONNECTION_STRING = util.MustGetEnv("DATABASE_CONNECTION_STRING")
var PORT = util.MustGetEnv("PORT")
