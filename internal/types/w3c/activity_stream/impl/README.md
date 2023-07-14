# Activity Stream Implementations

## TODO
- [ ] Finish building out necessary types
  - [ ] Activity
    - [ ] Update
  - [ ] Object
    - [ ] Article
    - [ ] Document
  - [ ] Properties
    - [ ] TBD given what's needed for an initial simple API

## Comment on Dependency Injection
Since ActivityStreams are recursive in nature keeping the files within the same package makes our lives much easier.
Initially attempted to do a manager style singleton DI pattern, used in `go-fed/activity`, however seemed to complex
given we can just keep all ActivityStream implementations in the root `impl` package. This does have the trade-off of
having 100's of files in this folder, however greatly simplifies/remove necessity to think about DI which I feel is
worth it.