# Health Check for Docker Clusters

Study of a docker-compose cluster whose overall health is checked by a dedicated container;
in essence, it runs smoke tests from _inside_ the cluster network.

Inspired by an actual issue at work, this mini project mainly serves as investigation grounds for 
which language provides the nicest way to implement "scripts" in a container context.
Criteria are:

 - Easy to write correct scripts. Ideally testable.
 - Good abstractions for readable scripts.
 - Minimal tooling required and available as Docker images.
 - Low runtime requirements for the final container.
 
Yes, I want to get rid of the ubiquituous "small" shell script as a default approach.

### Go

<!-- TODO write some observations -->

## Run

```bash
docker build -t healthcheck --build-arg checks=healthchecks.json healthcheck/
docker-compose -d up
```
