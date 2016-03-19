# fossasia-2016-graphql-demo

Demo project used to showcase GraphQL capabilities

## Slides

Slides for the talk @ FOSSASIA 2016 available here:
__Introduction to GraphQL (or How I Learned to Stop Worrying About REST APIs__:
- [PDF](https://github.com/sogko/fossasia-2016-graphql-demo/raw/master/slides/fossasia-2016-presentation.pdf)
- [Slideshare](http://www.slideshare.net/AhmadHafizIsmail/introduction-to-graphql-or-how-i-learned-to-stop-worrying-about-rest-apis)

## Reminder

#### We are looking for more contributors for `graphql-go`
Feel free to dive right in and submit a PR @ [graphql-go/graphql](https://github.com/graphql-go/graphql)

Participate in our discussion =)

### Features
Forked from [https://github.com/graphql-go/playground](https://github.com/graphql-go/playground)

- [graphql-go](https://github.com/graphql-go/graphql): Golang GraphQL library
- [graphql-relay-go](https://github.com/graphql-go/relay): Golang GraphQL library helper to construct Relay-compliant server
- [graphiql](https://github.com/graphql/graphiql): In-browser IDE to explore GraphQL queries
- [Starwars GraphQL Schema](https://github.com/graphql-go/relay/tree/master/examples/starwars): GraphQL example schema defined with Relay capabilities with the help of `graphql-relay-go`.

### To run locally
```bash
# `cd` to project directory
$ cd <project_dir>

# get all dependencies
$ go get ./...

# launch server
$ go run main.go

# Go to http://localhost:8080 on your browser
```
