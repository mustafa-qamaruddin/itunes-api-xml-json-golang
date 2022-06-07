# Project Description

Parse JSON/XML APIs using GoLang & Expose in GoLang

iTunes APIs Ref. https://developer.apple.com/library/archive/documentation/AudioVideo/Conceptual/iTuneSearchAPI/Searching.html#//apple_ref/doc/uid/TP40017632-CH5-SW1

Example GraphQL Queries

```graphql
{
  search(term: "mercedes sosa") {
    url
  }
}
```

```graphql
{
  search(term: "mercedes sosa", media: "music", entity: "musicTrack", attribute: "artistTerm") {
    url
  }
}
```