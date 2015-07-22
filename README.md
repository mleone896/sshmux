# sshmux
A lot of this code was google example code / somsone who put together a lib  


# Install sshmux lib
go get github.com/mleone896/sshmux

# Example Usage

```
import (
    "flag"

    "github.com/mleone896/sshmux"
)

var (
    file = flag.String("hostsFile", "", "use hostsfile")
    cmd  = flag.String("cmd", "hostname", "execute command")
)

func main() {
    flag.Parse()
    var hosts []string

    hosts = sshmux.GetHostsByFile(*file)

    var conf sshmux.Conf
    conf.User = "<your username>"
    conf.Key = "<your key here>"
    conf.Command = *cmd
    sshmux.Ssh(hosts, conf)
}

```

## License
apache  
