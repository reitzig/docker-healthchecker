# Health Check for Docker Clusters

Study of a docker-compose cluster whose overall health is checked by a 
dedicated container; in essence, it runs smoke tests from _inside_ the 
cluster network.
The actual checks being run are defined in
    [healthchecks.json](healthcheck/healthchecks.json).

### Motivation

Inspired by an actual issue at work, this mini project mainly serves as 
investigation grounds for which language provides the nicest way to 
implement "scripts" in a container context. 
Criteria are:

 - Easy to write correct scripts. Ideally testable.
 - Good abstractions for readable scripts.
 - Minimal tooling required and available as Docker images.
 - Low runtime requirements for the final container.
 
Yes, I want to get rid of the ubiquitous "small" shell script as a 
default approach.

### Tested Features

More specifically, the following language or library features are needed.

 - Command-line parameters.
 - Environment variables.
 - JSON parsing & generation.
 - File I/O.
 - External command execution. 

## Comparison

First, some numbers. Smaller is better.

|      | LOC | LOD |  SOF   |  SLI  |
|------|-----|-----|--------|-------|
| Ruby | ??? |  29 | 34.0MB | 258MB |
|  Go  | 135 |  15 |  5.4MB | 356MB |
<!-- TODO: update after Ruby is done -->
<!-- TODO: update after Go is done -->

Where the metrics are defined as follows:

 - _LOC_ -- Lines of "script" (and tooling) code
 - _LOD_ -- Lines in Dockerfile.
 - _SOF_ -- size overhead in the final image, 
     that is the size of the final image minus the size of `healthcheck`.
 - _SLI_ -- size of the largest intermediate image.

And now some subjective observations.

<!-- ### TODO: Bash -- the reference -->

### Ruby -- why scripting languages don't help

### Go

 - The standard library (or, for that matter, the official Docker container)
   has everything we need here.
 - No tooling beyond the compiler required.
 - Code is reasonably neat and could be tested.
 - Struct-based JSON marshalling makes for safe code.
 - Some gaps in the standard library that are off for a modern language,
   e.g. higher-order collection functions.

<!-- ### TODO: Kotlin Native -->

<!-- ### TODO: Crystal -->

<!-- ### TODO: Swift? -->

<!-- ### TODO: Rust? -->

<!-- ### TODO: Nim? -->

## Run

```bash
docker build -t healthcheck --build-arg checks=healthchecks.json healthcheck/
docker-compose -d up
```

Check
    [docker-compose.yml](docker-compose.yml)
for the port mappings of the individual healthcheckers.
