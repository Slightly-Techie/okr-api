# okr-api
Objectives and Key results platform

### Run App
```bash
docker-compose up
```

### Authentication notes

I followed this tutorial: https://jeanklaas.com/blog/set-up-oauth-with-react-and-go-with-gin/
Not entirely so the code might be a bit different but frontend may need this if they're confused.

Here's how to set up google auth from google's end. https://developers.google.com/identity/gsi/web/guides/overview

I need to work on that.

```
{
  sub: '101506029931667196787',
  email: 'addodiabene69@gmail.com',
  email_verified: true,
  name: 'Addo Diabene',
  picture: 'https://lh3.googleusercontent.com/a/ACg8ocKjUusyFCStfj5lQ7s9k3Jk_sGHAiqHZRavGzZRZ5TZ=s96-c',
}
```


## objectives crud docs

> create endpoint 
>post : http://localhost:5000/objective/create
### sample body
```
{
  "title":"test",
  "description":"testing our endpoints",
  "assignee":"alex",
  "userid":"111-222-333"
}
```
>read endpoint
>get: http://localhost:5000/objective/get-all/:user-id
```
```

>update endpoint
>put endpoint : http://localhost:5000/objective/update/:objective-id
### sample body
```
{
  "title":"tested 2",
  "description":"testing our endpoints",
  "assignee":"alex",
}
```

>delete endpoint
>delete : http://localhost:5000/objective/update/objective-id
```
```