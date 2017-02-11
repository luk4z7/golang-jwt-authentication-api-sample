# Middleware JWT

### Getting started

Clone the repository in folder do you prefer
```bash
cd /var/www
git clone https://github.com/luk4z7/middleware-jwt
```

**Execute the file `init.sh` for up the docker containers**

```bash

https://github.com/luk4z7/middleware-jwt for the canonical source repository
Middleware-jwt


 __  __ _     _     _ _                                   _          _
|  \/  (_) __| | __| | | _____      ____ _ _ __ ___      (_)_      _| |_
| |\/| | |/ _` |/ _` | |/ _ \ \ /\ / / _` | '__/ _ \_____| \ \ /\ / / __|
| |  | | | (_| | (_| | |  __/\ V  V / (_| | | |  __/_____| |\ V  V /| |_
|_|  |_|_|\__,_|\__,_|_|\___| \_/\_/ \__,_|_|  \___|    _/ | \_/\_/  \__|
                                                       |__/
DOCKER
Generate new containers ? [ 1 ]
Delete all containers ?   [ 2 ]
Start new build ?         [ 3 ]

```

Test api `/test/hello` receive `401 Unauthorized`
```bash
➜  middleware-jwt (master) ✗ curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://127.0.0.1:6060/test/hello

HTTP/1.1 401 Unauthorized
Date: Sat, 11 Feb 2017 12:53:45 GMT
Content-Length: 0
Content-Type: text/plain; charset=utf-8

```

Get token
```bash
➜  middleware-jwt (master) ✗ curl -H "Content-Type: application/json" -X POST -d '{"Username":"root", "Password":"12345"}' http://127.0.0.1:6060/token-auth
{"token":"eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODcwNzcwNTksImlhdCI6MTQ4NjgxNzg1OSwic3ViIjoiIn0.omoQv8ViTOI--QI3TQB4VAdhOLtyKman_fUnG7Hf5EQx4AhNJDgGBmSEOsCMZokE4AHSxAIYFmhFQ6oW2XZQDSJM7jqpn9Mi2TmtKiWnOcU4bRNwYcI8wjFyoWcxQW8M8sPJ9h3MpAnRaZQ1Z3hippK8PjAuujtQfnNUNqdOPuqRpK_r48fiC2BRz5Z18_DBjS9Xl7ZZOTtBxHRA-BNcnSLzzoMCE9VJ83oEs2Q29aBGUW4Ghfz-k8eDHlI1q6l7dvs4Yz_pHAim0__m0do9j8hAHMrWahb2eVLO19kmug6V3Rd2JFY3qLNh0MNW7DHLbwnZ7n7mjXy5KyKGqK07Gg"}
```

Access the route `test/hello` with authorization
```bash
➜  middleware-jwt (master) ✗ curl -H "Authorization: Bearer eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODY5NTgwMjMsImlhdCI6MTQ4NjY5ODgyMywic3ViIjoiIn0.nUgmyKnVUGOLpzBNpCiV5B13oRVfunZi-5tmkDFizbfwKhIVi_2oysZeRwozZQJYSLkVrU8iVAoUxUfh43-1MhL2tUei_SyOiAnwBmCq1KIBzo0yznQQylg_zUptHEJFzofqNPNyr_NS60VJTCGW-FVElSX-k-ecTlOXmFLE-Geg4MAn1_wxnIJpRrLha3httMckk1zTkVkc6sWZVlQlfbS81e9xYWvrjF88xo3pb9XdBc4KWCfRT4Nz3pNMNdTcEWWvVC2Xk9DdWbTIXr0PH4AivDiIPeBh331wB04dduc9_ufXnal3TlS2ZnUm0jYJRaFPaDbacWpxERr402L8mw" http://127.0.0.1:6060/test/hello

Hello, World!
```