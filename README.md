## GoBalancer

GoBalancer is a simple load balancer implementation, showcasing the simplicity of concurrent programming in Go. The core of the load balancer is a min heap, which balances work based on the "least busy" worker, to ensure a fair distribution of load across the worker pool.

## References
 - The implementation above was inspired by Rob Pike's 2012 Go talk: https://talks.golang.org/2012/waza.slide
