[![GoDoc](https://godoc.org/github.com/simar7/crdt?status.svg)](https://godoc.org/github.com/simar7/crdt)

# crdt 
A CRDT for LWW (Last-Writer-Wins) element set implementation in Go.

####  Prerequisites    
<code> redis-server </code>    
<code> radix.v2 </code>    


#### How to use?
1) Testing native Go implementation <code>crdt.go</code>     
<code> go test -v ./. </code> 

2) Testing redis based Go implementation <code>crdt_redis.go</code>    
<code> cd redis </code>    
<code> go test -v ./... -zset your-zset-name --port redis-server-port-number </code>

