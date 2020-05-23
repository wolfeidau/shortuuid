# shortuuid

This library uses [base58 encoding](https://en.wikipedia.org/wiki/Base58) to shorten the string representation of [Universally Unique IDentifier (UUID)](https://tools.ietf.org/html/rfc4122) values making them smaller when stored in a document database such as DynamoDB.

This library provides a reduction of ~39% in storage of the string representation of these values.

[![GitHub Actions status](https://github.com/wolfeidau/shortuuid/workflows/Go/badge.svg?branch=master)](https://github.com/wolfeidau/shortuuid/actions?query=workflow%3AGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/wolfeidau/shortuuid)](https://goreportcard.com/report/github.com/wolfeidau/shortuuid) [![Coverage Status](https://coveralls.io/repos/github/wolfeidau/shortuuid/badge.svg?branch=master)](https://coveralls.io/github/wolfeidau/shortuuid?branch=master)


# example

before `1c807ffe-b179-4a15-aa92-e06f5b5cf268`
after  `4X8k1aRv5UMp9pdKzxaHwu`

# License

This application is released under Apache 2.0 license and is copyright [Mark Wolfe](https://www.wolfe.id.au).