package params

import "fmt"

var Port = 8182
var Origin = fmt.Sprintf("http://localhost:%d", Port)
var Path = "/"
var URL = fmt.Sprintf("%s%s", Origin, Path)
var File = "go.mod"
var Comment = "module descriptor"
