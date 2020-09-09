# I'm learning about Docker!

I want to write a Go example that uses Go routines to collect soccer player statistics.

I want this Go code to run in a Docker container.

Future upgrades?
* [x] Another Docker container running a simple database to store results.
* [ ] Do more stuff with Kubernetes?
* [ ] Hook in to real API called API-FOOTBALL (paid)

# How to?

Grab the dependencies
```
dep ensure -v
```

Build the containers:
```
docker-compose build
```

Start the services:
```
docker-compose up
```
