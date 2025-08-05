# go-wanikani-api

An unofficial API client designed for the WaniKani API.

## Install

```
go get github.com/KaniLeap/go-wanikani-api
```

## Quick start

```go
ctx := context.Background()
cl := wk.NewClient(os.Getenv("WANIKANI_TOKEN"))

// Get current user
u, err := cl.GetUser(ctx)
if err != nil { log.Fatal(err) }
fmt.Println(u.Data.Name, u.Data.Level)

// List subjects (first page)
subs, err := cl.GetSubjects(ctx, wk.WithLevels([]int{1,2,3}))
if err != nil { log.Fatal(err) }
for _, it := range subs.Data.Data {
    fmt.Println(it.Id, it.Data.Slug)
}

// Paginate
for subs.HasNext() {
    if err := subs.Next(); err != nil { break }
}

// Get assignments available for review
asgn, err := cl.GetAssignments(ctx, wk.WithAvailableReviews(true))
if err != nil { log.Fatal(err) }
fmt.Println("assignments count:", asgn.Data.Count)
```
