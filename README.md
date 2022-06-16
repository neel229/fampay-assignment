
# fampay-assignment

### Setup
To get the application up and running, make sure docker is installed. Then, run the following commands:

To build the docker image
```
docker build -t fampay-assignment:latest .
```
Once the docker image is built, create a container using the below command
```
docker run -d -p 8080:8080 fampay-assignment:latest
```

Now that the server is running on port `8080`, play around with the API endpoints.

<br />

### API Endpoints
#### Retrieve Videos
- Method: `GET`
- Path: `localhost:8080/`
- Query Params: `limit` and `offset`
- Example Query: `localhost:8080/?limit=5&offset=0` returns the first 5 videos

#### Search With Title
- Method: `POST`
- Path: `localhost:8080/title`
- Body: `{"title": "<title>"}` must be in json 
- Example Query: `localhost:8080/title` with body `Mohammed Salah football skill football short video`

#### Search With Description
- Method: `POST`
- Path: `localhost:8080/description`
- Body: `{"description": "<description>"}` must be in json 
- Example Query: `localhost:8080/description` with body `Cristiano ronaldo vs fcbarcelon #ronaldo #fcbarcelona #football.`

#### Full-text search
- Method: `POST`
- Path: `localhost:8080/fts`
- Body: `{"keyword": "<keyword>"}` must be in json 
- Example Query: `localhost:8080/fts` with body `skill`

<br/>

> All the example queries can be ran without any changes as the data for the same is present. 
> Upon starting the server, the server will fetch latest videos for keyword "football"
> and store new videos. You can first call the "retrieve videos" endpoint and note down
> title, description and a keyword of your choice to carry out test queries for other endpoints.
