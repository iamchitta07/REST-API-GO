GET     /events                 => Get a list of available events
GET     /events/<id>            => Get a list of available events
POST    /events/                => Create a new bookable event (auth required)
PUT     /events/<id>            => Update an event (auth required)(only by creator)
DELETE  /events/<id>            => Delete an event (auth required)(only by creator)
POST    /signup                 => Create a new user
POST    /login                  => Authenticate user (Auth token(JWT))
POST    /events/<id>/register   => Register user for a event (auth required)
DELETE  /events/<id>/register   => cancel regestration (auth required)


// Packages
/models => event.go -> stores and fetches events data in/from DB
