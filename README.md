# spotifyApiGo

Wrapper for spotify API using Go

To use this repository with your spotify account, First set the `.env` file with:

- `CLIENT_ID`
- `CLIENT_SECRET`

For more details to obtain those values, see [Spotify Web Api docs][spotify-api].

## Run the app
Go to [http://localhost:8080/test-cred](http://localhost:8080/test-cred) to check if the app is working

### Login
Go to [http://localhost:8080/login](http://localhost:8080/login) to log into your profile

### Get user
Go to [http://localhost:8080/get-user](http://localhost:8080/get-user) to view your spotify account profile

### Get Top Artist/Tracks 
Go to [http://localhost:8080/get-top/{type}/{term}/{limit}](http://localhost:8080/get-top/{type}/{term}/{limit}) to get your top tracks/artist in the past time period.
- `type` value is either "tracks" or "artists", to get either top tracks or artists
- `term` value is either "short" or "medium", to set the time period either short term (4 weeks) or medium term (6 months)
- `limit` value is integer specifying the maximum number of items to return

### Get Top Artist/Tracks Default
Go to [http://localhost:8080/get-top/{type}](http://localhost:8080/get-user/{type}) to get your top tracks/artist in the default value where it is the same as above, with default value of medium term and 10 items limit

### Get recommendations
Go to [http://localhost:8080/recommendations](http://localhost:8080/recommendations) to view a list of tracks recommendation based on your listening activities it consist of 4 types:
- `by_tracks` where the recommendation parameters is coming from your top tracks
- `by_artists` where the recommendation parameters is coming from your top artists
- `by_genres` where the recommendation parameters is coming from your top genres
- `by_random` where the recommendation parameters is coming from the mix of your top tracks, artists and genres.

### Refresh Token
Go to [http://localhost:8080/refresh-token](http://localhost:8080/refresh-token) if the token from /login is expired (it is around 1 hour)

This API is meant to simplified the Spotify Web API request response so it can be used for a monthly spotify recap like a receiptify, feel free to use it with your front end.

[spotify-api]: https://developer.spotify.com/documentation/web-api/