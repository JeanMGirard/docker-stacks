BaseFunction for R
================

Use this FaaS function using R.

**Deploy the base R function**

(Make sure you have already deployed FaaS with ./deploy\_stack.sh in the root of this Github repository.

-   Option 1 - click *Create a new function* on the FaaS UI

-   Option 2 - use the [faas-cli](https://github.com/openfaas/faas-cli/) (experimental)

<!-- -->

    # curl -sSL https://get.openfaas.com | sudo sh

    # faas-cli -action=deploy -image=functions/base:R-3.4.1-alpine -name=baser
    200 OK
    URL: http://localhost:8080/function/baser

**Say Hi with input**

`curl` is good to test function.

    $ curl http://localhost:8080/function/baser -d "test"

**Customize the transformation**

If you want to customise the transformation then edit the Dockerfile or the fprocess variable and create a new function.

**Remove the function**

You can remove the function with `docker service rm baser`.
