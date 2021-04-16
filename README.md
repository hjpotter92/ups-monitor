# ups-monitor #

A mini project for observing and logging the input voltage for the
APC-BR1500G ups installed.

The project's `script.py` however, works with any NUT compliant
device.

## Setup ##

The project is designed with a linux system in mind (and has been
deployed to a raspberry pi device). The system published collected
metrics to the thingspeak channel.

### systemd units ###

There are `systemd` example units with the project, to setup the
automated activity monitor. The units require defining certain
environment values before they can be used effectively.

There is a `.service` and a `.timer` unit defined.

#### environment values ####

  * `THINGSPEAK_CHANNEL_ID`: Channel id as given by thingspeak.com.
  * `THINGSPEAK_API_KEY`: Your channel generated api key with write
    permissions.
  * `NUT_UPS_NAME`: Name of the NUT compliant device. The device
    should have a `input.voltage` property.

### python script ###

The python file `script.py` talks to the NUT server to fetch provided
device's `input.voltage` variable and then forwards the same to the
thingspeak channel (`field1`).

### slack bot ###

You can use ThingHTTP service to trigger an HTTP call whenever there's
a powercut or low voltage, or line fluctuations etc. and send it over
to a lambda function. This lambda can then send a message over to
slack channel informing about the same.

#### environment values ####

  * `SLACK_TOKEN`: The token generated for your slack bot. Generally
    of the format `xoxb-*`
  * `SLACK_CHANNEL`: Slack channel id where the message should be
    published.

Keep in mind that the message has not been configured to act as an
embed. You can modify the message sent as per your requirements.
