# Resource Helper

[![N|Solid](https://cldup.com/dTxpPi9lDf.thumb.png)](https://github.com/nguyencatpham/request-handler)

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://github.com/nguyencatpham/request-handler)

Resource helper support for parser data from resource service.

# Getting Started!

  #### Model input
  ```
	var toMap map[string]interface{}
	var result map[string]interface{}

	// handle bussiness here

	conv := marshaller.ConventionalMarshaller{Value: toMap}
	b, err := conv.MarshalJSON()
	if err != nil {
		// handle error here
	}
	json.Unmarshal(b, &result)
  ```

### Installation

Request handler requires [Go](https://golang.org/) v1.11+ to run.

Install the package.

```sh
$ go get github.com/onskycloud/marshaller
```

#### Kubernetes + Google Cloud

See [KUBERNETES.md](https://github.com/joemccann/dillinger/blob/master/KUBERNETES.md)


### Todos

 - Write MORE Tests
 - Add Night Mode

License
----

MIT