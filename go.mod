module github.com/semihalev/sdnsexampleplugin

require (
	github.com/miekg/dns v1.1.30-0.20200514105037-b7da9d95e0ed
	github.com/semihalev/log v0.0.0-20200513200305-7f20b48b19b7
	github.com/semihalev/sdns v1.0.1-0.20200624190509-960e4a656fa1
)

replace github.com/semihalev/sdns => ../sdns

go 1.13
