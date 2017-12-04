# Docker ELK stack

## Requirements

### Host setup

1. Install [Docker](https://www.docker.com/community-edition#/download) version **1.10.0+**
2. Install [Docker Compose](https://docs.docker.com/compose/install/) version **1.6.0+**
3. Clone this repository


## Usage

### Starting up the stack

```console
$ docker-compose up -d
```

You can check that everything is running this way:

```console
$ docker ps
```

Give Kibana a few seconds to initialize, then access the Kibana web UI by hitting
[http://localhost:5601](http://localhost:5601) with a web browser.

By default, the stack exposes the following ports:
* 5000: Logstash TCP input.
* 9200: Elasticsearch HTTP
* 9300: Elasticsearch TCP transport
* 5601: Kibana

**WARNING**: If you're using `boot2docker`, you must access it via the `boot2docker` IP address instead of `localhost`.

**WARNING**: If you're using *Docker Toolbox*, you must access it via the `docker-machine` IP address instead of
`localhost`.

Now that the stack is running, you will want to inject some log entries. The shipped Logstash configuration allows you
to send content via TCP:

```console
$ nc localhost 5000 < /path/to/logfile.log
```

## Initial setup

### Default Kibana index pattern creation

When Kibana launches for the first time, it is not configured with any index pattern.

#### Via the Kibana web UI

**NOTE**: You need to inject data into Logstash before being able to configure a Logstash index pattern via the Kibana web
UI. Then all you have to do is hit the *Create* button.

Refer to [Connect Kibana with
Elasticsearch](https://www.elastic.co/guide/en/kibana/current/connect-to-elasticsearch.html) for detailed instructions
about the index pattern configuration.

#### On the command line

Run this command to create a Kibana index pattern:

```console
curl -XPUT -D- 'http://localhost:9200/.kibana/doc/index-pattern:docker-elk' \
    -H 'Content-Type: application/json' \
    -d '{"type": "index-pattern", "index-pattern": {"title": "logstash-*", "timeFieldName": "@timestamp"}}'
```

This will automatically be marked as the default index pattern as soon as the Kibana UI is opened for the first time.


#### Once kibana is ready

You can test your stack producing dummy log with the following code compiled (golang), just fire up a terminal and run:
###### For linux:

```console
cd docker-elk
./producer/linux/producer
```
###### For MAC OSX:
```console
cd docker-elk
./producer/osx/producer
```

###### For windows:
```console
cd docker-elk
./producer/osx/producer.exe
```

##### You can find the source file in producer/src directory

### Now you can open your Kibana dashboard and you should see the following:

<img width="1676" alt="image" src="https://user-images.githubusercontent.com/10232243/33532600-bbee678a-d879-11e7-9855-209b6b1f8795.png">
