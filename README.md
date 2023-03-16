# Envoy rate limit example
To set up all the dependencies for this test: <br>
`docker compose up`

You can run the vegeta load test application, to see the rate limiting in action. <br>
`make load-test`

The rate limit service, should limit the requests based on the `Authorization` header, and allow up to 
10 requests per minute per unique value. This can be configured in [config.config](./examples/ratelimit/config/config.yaml)

To see how many requests are done by a specific authorization header: <br>
- connect to the redis container
- find the key that matches your request header. you can use `keys *` command to see all keys
- getting the value of the key will show how manyy requests this key has done: `get <key>`

### Monthly rate limits/ Quotas.
At the branch [customer-rate-limit](https://github.com/gegaryfa/envoy-rate-limit/tree/custom_rate-limit), you can find 
a solution where the reference rate limiter is added as a submodule and adjusted so it supports
 monthly rate limits as well.

## Readings/blogs/examples

- [Envoy rate limit](https://github.com/envoyproxy/ratelimit)
- [Envoy rate limit example](https://github.com/jbarratt/envoy_ratelimit_example/tree/master)
- [Envoy rate limits article 1](https://serialized.net/2019/05/envoy-ratelimits/)
- [Envoy rate limits article 2](https://medium.com/dm03514-tech-blog/sre-resiliency-bolt-on-sidecar-rate-limiting-with-envoy-sidecar-5381bd4a1137)
- [Understanding envoy rate limits](https://xuorig.medium.com/understanding-envoyproxys-rate-limiting-cb3e00c9be2d)
- [sophisticated-rate-limiting](https://www.funnel-labs.io/2022/10/10/envoyproxy-3-sophisticated-rate-limiting/#what-is-rate-limiting-for)