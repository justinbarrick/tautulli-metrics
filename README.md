Imports your Tautulli data into influxdb for processing, including any historical data
from before you launched the service (or even from during a service interruption!).

To run, set the following variables:

* `TAUTULLI_URL`: the URL of your tautulli instance.
* `TAUTULLI_API_KEY`: the api key to your tautulli instance (check the configuration file).
* `INFLUX_URL`: the url to your influxdb instance.
* `INFLUX_DB`: the name of the influxdb database to use.
* `INFLUX_USER`: (optional) the influxdb username.
* `INFLUX_PASS`: (optional) the influxdb password.
