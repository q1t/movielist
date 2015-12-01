# Watch List
This is just a simple watch list. You can add a list of movies.

~~Account isn't necessary, you can continue using watch list as a guest.
Everything will be stored at the browser's [local storage](https://developer.mozilla.org/en/docs/Web/API/Window/localStorage).~~ TODO

~~However, having registered email allows you to connect your data with your email,
access it on the website or backup it to that email .~~ TODO

~~To build and host it by yourself (optionally remove the registration part), see
[github](https://github.com/sysashi/watchlist).~~ Not ready yet.

### API

Every api endpoin accepts json and returns json.

``` 
POST /api/login
POST /api/register

Expect username and password, return error or validation errors, or if
login/registration was successful, return bearer (token).
```
Since you've obtained a token, you're able to interact with private api by
sending requests to ``` /api/private ```

There is only one endpoint availible for the private api.
``` 
GET /api/private/user 

Return all user data (id, username, email (if exists), lists).

```


### Todo
- submit forms by pressing enter
- adding lists
- adding movies



