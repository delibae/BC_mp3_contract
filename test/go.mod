module example.com/hello

go 1.16

require (
	rsc.io/quote v1.5.2 // indirect
	example.com/helloworld v0.0.0 
)

replace (
	example.com/helloworld v0.0.0 => ./helloworld
)
