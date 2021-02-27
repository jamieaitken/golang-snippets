# Encrypted Elasticache Communication

With Redis 6.0, native support for SSL is introduced. Until then additional tools like [Stunnel](https://www.stunnel.org/)
had to be used in order to achieve an SSL connection with your redis instance.

If you're using a managed service like AWS' Elasticache, you have been able to enable encryption-in-transit and encryption-at-rest
[since 2017](https://aws.amazon.com/about-aws/whats-new/2017/10/amazon-elasticache-for-redis-now-supports-in-transit-and-at-rest-encryption-to-help-protect-sensitive-information/).

Now if you're using the latter, you may be wondering how to connect to your cluster? It's as simple as
specifying an empty `tls.Config` in your go-redis client. No public or private keys, just an empty `tls.Config`.
