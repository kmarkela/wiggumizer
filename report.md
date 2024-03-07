# Wiggumize Report

__Scope:__
- https://0af1000403f5a14b81ce752a009d0095.web-security-academy.net:443
- https://content-autofill.googleapis.com:443
- https://www.google-analytics.com:443
- https://fonts.gstatic.com:443
- https://0abe003203cf09d58380418200940058.web-security-academy.net:443
- https://accounts.google.com:443
- https://www.gstatic.com:443
- https://ps.containers.piwik.pro:443
- https://www.googletagmanager.com:443
- https://portswigger.net:443
- https://0a0b00b404822135828b51d200ed008c.web-security-academy.net:443
- https://www.google.com:443
- https://ps.piwik.pro:443


__List of Checks:__
- __xml:__ This module is searching for XML in request parameters
- __redirect:__ This module is searching for Redirects
- __ssrf:__ This module is searching for URL in request parameters.
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
### Finding 1. - Redirect Found. Status: 304

__Host: https://www.google-analytics.com:443__ 

_Evidens:_

```
Path: /analytics.js

```
### Finding 2. - Redirect Found. Status: 304

__Host: https://ps.containers.piwik.pro:443__ 

_Evidens:_

```
Path: /287552c2-4917-42e0-8982-ba994a2a73d7.js

```
### Finding 3. - Redirect Found. Status: 304

__Host: https://www.google-analytics.com:443__ 

_Evidens:_

```
Path: /analytics.js

```
### Finding 4. - Redirect Found. Status: 304

__Host: https://ps.containers.piwik.pro:443__ 

_Evidens:_

```
Path: /287552c2-4917-42e0-8982-ba994a2a73d7.js

```
### Finding 5. - Redirect Found. Status: 302

__Host: https://0af1000403f5a14b81ce752a009d0095.web-security-academy.net:443__ 

_Evidens:_

```
Path: /login

```
_More Details:_

```
{"csrf":"oONNaOIWI5JJSLOvafTMskAuktH7Wyoc","username":"wiener","password":"peter"}
```
