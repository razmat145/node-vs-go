# Node vs Go

With aim of learning, I'm comparing various operations between the two; starting with very easy (hello world) and towards wherever my understanding and experimentation takes me.

Currently, recording and comparing Prometheus metrics between the two:

- Requests Total `http_requests_total`
- Average response times `rate(http_response_time_seconds_sum[1m]) / rate(http_response_time_seconds_count[1m])`
- Request rate `rate(http_requests_total[1m])`
- 0.95 Perc response time `histogram_quantile(0.95, rate(http_response_time_seconds_bucket[1m]))`
- Current memory usage `process_resident_memory_bytes`
- Rate of change of memory usage `deriv(process_resident_memory_bytes[1m])`
- Peak memory usage `max_over_time(process_resident_memory_bytes[1m])`
- Rate CPU usage `rate(process_cpu_seconds_total[1m])`

Note: this is not meant to be a serious benchmark or anything; it's meant to be and should be treated as a learning exercise.

## Current core rules

Various tests will vary the rules based on available libraries;  
Two golang spammers are setup to spam 25k requests, with a 2 second sleep in between, constantly.

The receivers are 4 HTTP APIs:

- node-express - one of the most used 'slim' NodeJS libraries; in different stages with popular 'slim' libraries
- node-std - std library implementation of a http server; in different stages with bare bones libraries
- go-echo - one of the most used 'slim' Golang libraries; in different stages with popular 'slim' libraries
- go-std - std library implementation of a http server; in different stages with bare bones libraries

At time, the comparison might not be fair; or my usage of certain constructs and libraries might not be equal between the above. The comparisons and their results will evolve over time as my understanding does.

## Setup and Running

This playground has to be run via docker.

In order to build and run, one can simply use `npm run compose`

## Current bench tries

- [Hello World!](./readme/hello.md)
- [Factorial](./readme/factorial.md)
- [Garbage creation (on purpose this time)](./readme/garbage.md)
