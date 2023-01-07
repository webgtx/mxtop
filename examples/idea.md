# mxtop - top for your server park

# -v, --version flag
```
~> mxtop -v                                   
  __  __   ___                         mxtop
 |  |/  `.'   `.                      ╭───────────────────────────────────────────╮
 |   .-.  .-.   '                     │ Authors                                   │
 |  |  |  |  |  | ____     _____      │  ├── webgtx (https://github.com/webgtx)   │
 |  |  |  |  |  |`.   \  .'    /      │  └── ssleert (https://github.com/ssleert) │
 |  |  |  |  |  |  `.  `'    .'       │ License - BSD 3-Clause                    │
 |  |  |  |  |  |    '.    .'         │ Language - Go                             │
 |__|  |__|  |__|    .'     `.        │ Repo - https://github.com/webgtx/mxtop    │
                   .'  .'`.   `.      │ Commit - $COMMIT_HASH                     │
                 .'   /    `.   `.    │ Version - 0.0.1                           │
                '----'       '----'   ╰───────────────────────────────────────────╯
```

# -h, --help flag
```
~> mxtop -h

  mxtop - top for your server park
 ╭──────────────────────────────────────────────────╮
 │ flag             │                   description │
 │──────────────────────────────────────────────────│
 │ -i --interactive │ run mxtop in interactive mode │
 │ -f --fetch       │ fetch info from server        │
 │ -v --version     │ get info about version        │
 │ -h --help        │ get help message              │
 ╰──────────────────────────────────────────────────╯
```

# -f, --fetch flag
```
~> mxtop -f
 ╭────────────────────────────────────────────────────────────────────────────────╮
 │ name      │ ver │ cpu | ram           │ up         │ ping    │ ip              │
 │────────────────────────────────────────────────────────────────────────────────│
 │ fedora    │ 36  │ 82% | 2300/15884 mb │ 1d 22h 37m │   10 ms │ 157.116.227.217 │
 │ arch      │ >3  │  1% |  158/15884 mb │     1h 20m │ 1345 ms │  211.216.51.124 │
 │ debian    │ 11  │ 43% |  385/15884 mb │     12h 1m │  128 ms │   3.237.156.154 │
 │ slackware │ 15  │  0% | 4583/15884 mb │   1d 1h 1m │   29 ms │ 146.132.226.133 │
 ╰────────────────────────────────────────────────────────────────────────────────╯
~> next_command
```


# -i, --interactive flag
```
~> mxtop -i
 ╭────────────────────────────────────────────────────────────────────────────────╮
 │ name      │ ver │ cpu | ram           │ up         │ ping    │ ip              │
 │────────────────────────────────────────────────────────────────────────────────│
 │ fedora    │ 36  │ 82% | 2300/15884 mb │ 1d 22h 37m │   10 ms │ 157.116.227.217 │
 │ arch      │ >3  │  1% |  158/15884 mb │     1h 20m │ 1345 ms │  211.216.51.124 │
 │ debian    │ 11  │ 43% |  385/15884 mb │     12h 1m │  128 ms │   3.237.156.154 │
 │ slackware │ 15  │  0% | 4583/15884 mb │   1d 1h 1m │   29 ms │ 146.132.226.133 │
 ╰────────────────────────────────────────────────────────────────────────────────╯
                                q - quit | r - reload
```


# table idea
### go code
```go
package main

func main() {
	tbl := [][]string{
		[]string{"name", "ver", "cpu", "ram", "up", "ping", "ip"},
		[]string{"fedora", "36", "82%", "2300/15884 mb", "1d 22h 37m", "10 ms", "157.116.227.217"},
		[]string{"arch", ">3", "1%", "2300/15884 mb", "1h 20m", "1345 ms", "211.216.51.124"},
		[]string{"debian", "11", "43%", "158/15884 mb", "12h 1m", "128 ms", "157.116.227.217"},
		[]string{"slackware", "15", "0%", "4583/15884 mb", "1d 1h 1m", "29 ms", "146.132.226.133"},
	}

	genTable(tbl)
}
```
### table
```
 ╭────────────────────────────────────────────────────────────────────────────────╮
 │ name      │ ver │ cpu | ram           │ up         │ ping    │ ip              │
 │────────────────────────────────────────────────────────────────────────────────│
 │ fedora    │ 36  │ 82% | 2300/15884 mb │ 1d 22h 37m │   10 ms │ 157.116.227.217 │
 │ arch      │ >3  │  1% |  158/15884 mb │     1h 20m │ 1345 ms │  211.216.51.124 │
 │ debian    │ 11  │ 43% |  385/15884 mb │     12h 1m │  128 ms │   3.237.156.154 │
 │ slackware │ 15  │  0% | 4583/15884 mb │   1d 1h 1m │   29 ms │ 146.132.226.133 │
 ╰────────────────────────────────────────────────────────────────────────────────╯
```

### info
- name and ver from /etc/os-release and https://github.com/ssleert/ginip
- cpu usage info https://stackoverflow.com/questions/11356330/how-to-get-cpu-usage
- ram usage info https://github.com/ssleert/memory

<br>

### symbols
```
│─╭╮╰╯
```