module github.com/semihalev/sdnsexampleplugin

require (
	github.com/miekg/dns v1.1.30
	github.com/semihalev/log v0.1.1
	github.com/semihalev/sdns v1.1.2
)

replace github.com/semihalev/sdns => ../sdns

go 1.13
