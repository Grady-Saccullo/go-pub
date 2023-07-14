# Activity Stream Implementations

## TODO
- [ ] Rename root files correctly under each package
- [ ] Pull in deserializers for all property packages (currently only property_object in brought in) to standardize
- [ ] Finish building out necessary types
  - [ ] Activity
    - [ ] Update
  - [ ] Object
    - [ ] Article
    - [ ] Document
  - [ ] Properties
    - [ ] TBD given what's needed for an initial simple API

## Dependency Injection
Since activity streams are recursive in nature needed to be able to inject deserializers in a singleton pattern
so there are no circular dependencies (cycle import). To achieve this a root level manager is created and setup
with all the deserializers. Local managers are then created for each package and set their manager to the root
level and define each necessary deserializer