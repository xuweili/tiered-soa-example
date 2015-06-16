# tiered-soa-example
An example to show deferpanic's dapper/zipkin style tracing in a SOA/micro-service environment.

Many users implement SOA - service oriented architecture or
'micro-services' in their companies. We believe this is a sea change and
want to show you that you can't just install old-school devops agents on
your boxes anymore - it's vital to be watching your APP traffic flow.

Servers aren't greek gods controlled by the BOFH anymore - they are free
range cattle.

## QuickStart

### Sign Up For a Free Account if You Haven't Yet
```bash
   curl https://api.deferpanic.com/v1.10/users/create \
        -X POST \
        -d "email=test@test.com" \
        -d "password=password"
```


### Find/Replace My Token with Yours
```bash
find . -name "*.go" -type f -print0 | xargs -0 sed -i '' 's/v00L0K6CdKjE4QwX5DL1iiODxovAHUfo/mytoken/g'
```

### Build && Run
```bash
make
make run
```

In this example we have a total of 5 'micro-services' running, named
naturally after the greek gods.

zeus
  - listens on port 3000
  - latency is 960ms

athena
  - listens on port 3001
  - talks to aphrodite
  - talks to dionysus
  - latency is 710ms

aphrodite
  - listens on port 3003
  - talks to athena
  - latency is 300ms

dionysus
  - listens on port 3004
  - talks to athena
  - latency is 400ms

hercules
  - listens on port 3002
  - talks to hercules
  - latency is 250ms
