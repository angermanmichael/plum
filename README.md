
#### Building

To Build Spnee Run the following commands:

```
./run/buildonce
./run/buildspnee
```

The **buildonce** script only needs to be run once :)

And any time you add new go packages or delete old ones.

#### Running

To Run Spnee Run the following command:

```
./spnee/spnee
```

#### Testing

Depending on the APIs you are testing different scripts and simulators
need to be run...

The rules API can be tested via the [rules simulator](https://github.com/stormasm/spinnakr-rule-simulator).

All other APIs can be tested by running

```
./run/buildtest
```

#### Notes

* All of the above commands assume you are in the top level directory.
* [RabbitMQ](http://www.rabbitmq.com/) must be running and configured with the correct channel names.
* [Redis](http://redis.io/) must be running and populated with initial data
