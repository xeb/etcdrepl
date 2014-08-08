etcdrepl
========

[![Build Status](https://travis-ci.org/xeb/etcdrepl.png)](https://travis-ci.org/xeb/etcdrepl)

`etcdrepl` is a REPL is a wrapper around [etcdctl](https://github.com/coreos/etcdctl).

Examples

```
etcdrepl> help

OPTIONS:
   --debug					output cURL commands which can be used to reproduce the request
   --no-sync					don't synchronize cluster information before sending request
   --output, -o 'simple'			output response in the given format (`simple` or `json`)
   --peers, -C '--peers option --peers option'	a comma-delimited list of machine addresses in the cluster (default: {"127.0.0.1:4001"})
   --version, -v				print the version
   --help, -h					show help
   

COMMANDS:
   mk		make a new key with a given value
   mkdir	make a new directory
   rm		remove a key
   rmdir	removes the key if it is an empty directory or a key-value pair
   get		retrieve the value of a key
   ls		retrieve a directory
   set		set the value of a key
   setdir	create a new or existing directory
   update	update an existing key with a given value
   updatedir	update an existing directory
   output	output response in the given format (`simple` or `json`)
   quit		exits the current program
   help, h	Shows a list of commands or help for one command
   
etcdrepl> 
```

And also:

```
etcdrepl> get test
test
etcdrepl> set test anothervalue
anothervalue
etcdrepl> get test
anothervalue
etcdrepl> output json
etcdrepl> get test
{"action":"get","node":{"key":"/test","value":"anothervalue","modifiedIndex":39,"createdIndex":39},"etcdIndex":39,"raftIndex":3115860,"raftTerm":1}
etcdrepl> set test somethingelse
{"action":"set","node":{"key":"/test","value":"somethingelse","modifiedIndex":40,"createdIndex":40},"prevNode":{"key":"/test","value":"anothervalue","modifiedIndex":39,"createdIndex":39},"etcdIndex":40,"raftIndex":3115872,"raftTerm":1}
etcdrepl> get test
{"action":"get","node":{"key":"/test","value":"somethingelse","modifiedIndex":40,"createdIndex":40},"etcdIndex":40,"raftIndex":3115874,"raftTerm":1}
etcdrepl> 
```

[![Screenshot](https://raw.githubusercontent.com/xeb/etcdrepl/master/screenshot.png)](https://raw.githubusercontent.com/xeb/etcdrepl/master/screenshot.png)
