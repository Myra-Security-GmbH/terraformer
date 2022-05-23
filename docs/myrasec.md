### Use with Myra Security

Example using a Myra Security API Key and corresponding Token:

```
export MYRASEC_API_SECRET=[MYRASEC_API_SECRET]
export MYRASEC_API_TOKEN=[MYRASEC_API_TOKEN]
./terraformer import myrasec --resources=domain
```

List of supported Myra Security services:
* `domain`
  * `myrasec_domain`
* `dns`
  * `myrasec_dns_record`
* `error pages`
  * `myrasec_error_pages`