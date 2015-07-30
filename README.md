# tiered-soa-example
An example to show deferpanic's dapper/zipkin style tracing in a SOA/micro-service environment.

There are a variety of tools out there that show you the HTTP latency
but very few allow you to drill down from the network call into the
app layer so you can see stats such as memory && GC at that point in time from
that service.

The ability to drilldown from a request just like you would do in chrome
is vital.

Many users implement SOA - service oriented architecture or
'micro-services' in their companies. We believe this is a sea change and
want to show you that you can't just install old-school devops agents on
your boxes anymore - it's vital to be watching your APP traffic flow.

Servers aren't greek gods controlled by the BOFH anymore - they are free
range cattle. If a single SRE can control 10k servers where does that
leave the developer?

We're here to help.

After going through this example you'll be able to see your distributed
requests like the following:

From Your Shell:
![shell](https://cloud.githubusercontent.com/assets/22410/8173886/493bad0e-1389-11e5-9f5c-5fa461c225b6.png)

From the High Level View:
![list](https://cloud.githubusercontent.com/assets/22410/8173888/507ddb32-1389-11e5-8958-6e56097a0196.png)

Clicking Through to Zeus's View from Olympus:
![zeus](https://cloud.githubusercontent.com/assets/22410/8173887/4c1aba60-1389-11e5-8ff7-53a5102966b1.png)

Athena's View from Zeus:
![athena](https://cloud.githubusercontent.com/assets/22410/8173889/55057cc8-1389-11e5-98b9-0d57fee72ac3.png)

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

In this example we have a total of 5 micro-services running, named
naturally after the greek gods.

zeus
  - listens on port 3000
  - talks to athena
  - talks to hercules
  - latency is ~ 1630ms
 
athena
  - listens on port 3001
  - talks to aphrodite
  - talks to dionysus
  - latency is ~ 1120ms

aphrodite
  - listens on port 3003
  - talks to athena
  - latency is ~ 610ms

dionysus
  - listens on port 3004
  - talks to athena
  - latency is ~ 510ms

hercules
  - listens on port 3002
  - talks to hercules
  - latency is ~ 510ms
