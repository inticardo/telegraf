# File Output Plugin

This plugin writes telegraf metrics to files

## Global configuration options <!-- @/docs/includes/plugin_config.md -->

In addition to the plugin-specific configuration settings, plugins support
additional global and plugin configuration settings. These settings are used to
modify metrics, tags, and field or create aliases and configure ordering, etc.
See the [CONFIGURATION.md][CONFIGURATION.md] for more details.

[CONFIGURATION.md]: ../../../docs/CONFIGURATION.md#plugins

## Configuration

```toml @sample.conf
# Send telegraf metrics to serial port
[[outputs.serial]]
  port = "/dev/ttyUSB0"

  ## Baud rate
  # brate = 9600

  ## Data bits
  # bits = 8

  ## Use HUPCL
  # hupcl = false

  ## Parity ("none"/"odd"/"even"/"mark"/"space")
  # parity = "none"

  ## Stop bit ("one"/"onepointfive"/"two")
  # stop = "one"

  ## Read timeout
  # timeout = "1s"

  ## Data format to output.
  ## Each data format has its own unique set of configuration options, read
  ## more about them here:
  ## https://github.com/influxdata/telegraf/blob/master/docs/DATA_FORMATS_OUTPUT.md
  # data_format = "influx"
```
