Message of the Neverchanging
===

*Message of the Neverchanging* is a demo service that exposes a REST endpoint which returns a never-changing message. Think of it as a test mule for working with infrastructure.

Message of the Neverchanging is inspired by the old [message of the day (motd)](https://en.wikipedia.org/wiki/Motd_(Unix)) tradition from UNIX. This one, however, does not change day by day nor at all. 

Endpoints
---

##### GET `/v1/motn`

Returns a JSON payload containing a never-changing message and the current UNIX time.

For example, 

```
{
    "message": "Code for the masses",
    "timestamp": 1550013554
}
```

##### GET `/v1/healthz`

Returns a JSON payload saying to indicate that everything is OK.

For example, 

```
{
    "status": "green"
}
```
