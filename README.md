# Aryan's assignment


## Running the project
After entering the project directory, run the following commands to run the project.

- `cp .env.example .env` . For this assignment, an `.env` should be present with the submission.
- `export $(grep -v '^#' .env | xargs)` or export .env values if you haven't
- `make docker-up`

The service should be available on `APP_PORT`. The default value is 8080. And for mongodb, the default is `27017`.

**Running tests**

- While the container is running, you can run `make test` to run the tests. 

The current project contains tests for lambda events. Also, for convenience, I'm using a db connection for testing. But ideally, it'll be done against mocks (eg: database/repository).


**API Endpoints**
> All post endpoints needs to include the header "Content-Type" : "application/x-www-form-urlencoded" and "Authorization" : "$token". `/auth/login` endpoint doesn't need to have auth token. 

- Postman collection is present inside `docs/`.

**[GET] `/health-check`**

Health check returns a response with timestamp if all the services are okay. During this check, it performs a `ping()` on the database, if it fails, the service will return `ServiceUnavailable (503)`.
```json{
"status": "ok",
"time": "2022-11-07 13:42:09.953685721 +0000 UTC m=+11.603784839"
}
```

**[POST] `/user/create`**

This endpoint is used to create random user and return responses. The password is set to `secret` for all users and being returned to the json response. Though in a real life scenario the sensitive data will be hidden.
```json
{
"id": "6368ff49985c599efc1a5c47",
"name": "Tessie Kohler",
"email": "stephenlarkin@koepp.info",
"password": "secret",
"age": 11,
"gender": "male"
}
```

**[POST] `/auth/login`**
This endpoint validates a user's credentials and returns a token. Which can be used to authenticate protected routes. 

```json
{
"token": "D799E13E8F7D9619297B39A11764D056D66EDD769E89DAC8BE382ED4BE2DE80B"
}
```

**[GET] `/profile`**
This endpoint returns a list of profiles. It accepts `age=$numeric` and `gender=male` or `gender=female` query parameters. Based on the parameters, it will send a filtered or unfiltered list (if any query parameter is not present).
The endpoint doesn't include location based search model. And the logic of aggregation is documented.
```json
{
"id": "6368d78d18506de8e0401bd9",
"name": "Annabelle Boyer",
"email": "donavonmclaughlin@legros.name",
"age": 22,
"gender": "male"
},
{...},
{
"id": "6368d4c805b6486d005a3ee9",
"name": "Karlee Robel",
"email": "macstiedemann@kub.biz",
"age": 20,
"gender": "male"
}]
```

**[POST] `/swipe`**
The endpoint records swipes for the authenticated user, and returns a response if there was already a swipe against the authenticated user's profile.

```json
{
"message": "swipe recorded",
"preference_matched": false,
"recorded_swipe_id": "63678b1fc63904ad04823dd3",
"matched_swipe_id": "" // If there is a match, it will include that in the response.
}
```

## Improvements
- The repositories can have some common methods.
- Add more test cases, mostly some happy and some sad path were tested.
- The testcase now tests with lambda. But the core logic is true for the app in general sense. We can extract the core of the test and simply add the wrapper for lambda on lambda's scenarios.
- Some handlers have direct bindings with services, it can be replaced by interfaces.
- Handling graceful shutdowns.
- With the .env value, `DB_URI` and `DB_PORT` is separated, so an accurate value depends on 2 different values, but it could have been done by one.
- Logging is not present at the moment, but in a production environment, there would be more logging, including tracing and spanning. 
- There is still room for improvements.
- Traits are as of now, static. 

## Code Structure
The communication flow among the components are uni-directional. Because of that, tracking error and making changed become much faster.
```
┌──────────┐    ┌──────────┐    ┌───────────┐    ┌────────────────────────────────────────┐
│          │    │          │    │           │    │           Request Serializer           │
│  Server  ├────►  Router  ├────►  Handler  ├────►────────────┬───────────┬───────────────┼───┐
│          │    │          │    │           │    │ Serializer │ Validator │ Entity Mapper │   │
└──────────┘    └──────────┘    └───────────┘    └────────────┴───────────┴───────────────┘   │
                                                                                              │
     ┌────────────────────────────────────────────────────────────────────────────────────────┘
     │
┌────▼────┐               ┌────────────┐              ┌─────────────────┐    ┌────────────┐
│         │               │            │              │    Presenter    │    │            │
│ Service ├───────────────► Repository ├──────────────►─────────────────┼────►  Renderer  │
│         │               │            │              │ Response Mapper │    │            │
└─────────┘               └─────┬──────┘ ┌──────────┐ └─────────────────┘    └────────────┘
                                │        │          │ 
                                └────────► Database │
                                         │          │
                                         └──────────┘
```

## Attractiveness logic
The logic is as follows:

Every user has some `traits`. Those are some random values from the Traits list. In users, its stored with it's id and value. The idea was to have something like
fifa's player statistics. They have some sort of values for different types of attributes (eg: Attack, defence etc). 

So, I'll take the authenticated user's traits, collect the ids. Check of other users where they have similar traits and sum their value. And sort them by value in the end.

## Absent of location based model

At first the idea was to simply add lat, long to the user's entities and do an aggregated query where I'll the take difference (geometric) between current user's lat, long with 
other users lat,long. But, obviously that would not give us the actual distance. I think we need some sort of proximity service for this. While I wasn't able to work
on this feature. But this is an idea how I can achieve this,

The logic turns into this query, 
```javascript
db.users.aggregate([
  {
    '$match' : {
      '$and': [
        { 'gender': 'male' },
        { 'age' : { '$lte': 100 } },
        { '_id' : { '$ne' : 'currentUserId' } },
        { 
          '$or' : [
              {'traits.id' : '1'},
              {'traits.id' : '4'}
            ] 
          }
      ]
    },
  },
  {
    '$project' :  { 
      attractiveness_score : { '$sum' : '$traits.value' },
      name: '$name',
      email: '$email',
      age: '$age',
      gender: '$gender',
      traits: '$traits'
    }
  },
  {
     $count: "email"
   }
]).sort({ 'attractiveness_score' : -1 })
```



**Idea** 

We want to store the user's location in a datastore that supports geohash. Modern relational (mysql, postgres) or other databases (redis) has this feature out of the box.
We want to keep users location on redis (or alternative). And while making a query, we want to plot a radius and select the user ids from geohash supported datastore
and simply add `{ 'id' : { 'in' : [ids...] }}` to the query. 

**Some reference**
- [https://www.memurai.com/blog/geospatial-queries-in-redis](https://www.memurai.com/blog/geospatial-queries-in-redis)
- [FAANG System Design Interview: Design A Location Based Service (Yelp, Google Places)](https://www.youtube.com/watch?v=M4lR_Va97cQ&t=1s&ab_channel=ByteByteGo)


