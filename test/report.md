# Wiggumize Report

__Scope:__
- https://0abe003203cf09d58380418200940058.web-security-academy.net:443
- https://0a0b00b404822135828b51d200ed008c.web-security-academy.net:443
- https://0af1000403f5a14b81ce752a009d0095.web-security-academy.net:443


__List of Checks:__
- __xml:__ This module is searching for XML in request parameters
- __redirect:__ This module is searching for Redirects
- __ssrf:__ This module is searching for URL in request parameters.
- __subd:__ This module is searching for 404 messages form hosting platformas
- __secret:__ This module is searching for secrets (eg. API keys)
- __lfi:__ This module is searching for filenames in request parameters. Could be an indication of possible LFI
--------------------

## xml_checker
### Finding 0. - Posible XML tag in params
__Host: https://0a0b00b404822135828b51d200ed008c.web-security-academy.net:443__ 

_Evidens:_

```
Path: /product/stock

```
_More Details:_

```
Body: <?xml version="1.0" encoding="UTF-8"?><stockCheck><productId>2</productId><storeId>2</storeId></stockCheck>

```
## redirect_checker
### Finding 0. - Redirect Found. Status: 302

__Host: https://0af1000403f5a14b81ce752a009d0095.web-security-academy.net:443__ 

_Evidens:_

```
Path: /my-account

```
### Finding 1. - Redirect Found. Status: 302

__Host: https://0af1000403f5a14b81ce752a009d0095.web-security-academy.net:443__ 

_Evidens:_

```
Path: /login

```
_More Details:_

```
{"csrf":"oONNaOIWI5JJSLOvafTMskAuktH7Wyoc","username":"wiener","password":"peter"}
```
