//go:build !custom || outputs || outputs.serial

package all

import _ "github.com/influxdata/telegraf/plugins/outputs/serial" // register plugin
