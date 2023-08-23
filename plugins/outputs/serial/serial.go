//go:generate ../../../tools/readme_config_includer/generator
package serial

import (
	_ "embed"

	"github.com/albenik/go-serial"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/internal"
	"github.com/influxdata/telegraf/plugins/outputs"
	"github.com/influxdata/telegraf/plugins/serializers"
)

//go:embed sample.conf
var sampleConfig string

type Serial struct {
	Port    string `toml:"port"`
	Brate   int    `toml:"brate"`
	Hupcl   bool   `toml:"hupcl"`
	Parity  string `toml:"parity"`
	Stop    string `toml:"stop"`
	Timeout int    `toml:"timeout"`

	serialOutput *SerialOutput
	encoder      internal.ContentEncoder
	serializer   serializers.Serializer
}

func (*Serial) SampleConfig() string {
	return sampleConfig
}

func (f *Serial) SetSerializer(serializer serializers.Serializer) {
	f.serializer = serializer
}

func (f *Serial) Init() error {
	// var err error

	// return err
	// Creación de objeto serial
	return nil
}

func (f *Serial) Connect() error {

	// Conexión efectiva del serial
	return nil
}

func (f *Serial) Close() error {
	// Cierre
	return nil
}

func (f *Serial) Write(metrics []telegraf.Metric) error {
	return nil
}

type SerialOutput struct {
}

func (so *SerialOutput) mapParity(p string) (serial.Parity, bool) {
	rtn, ok := map[string]serial.Parity{
		"none":  serial.NoParity,
		"odd":   serial.OddParity,
		"even":  serial.EvenParity,
		"mark":  serial.MarkParity,
		"space": serial.SpaceParity,
	}[p]
	return rtn, ok
}

func (so *SerialOutput) mapStop(s string) (serial.StopBits, bool) {
	rtn, ok := map[string]serial.StopBits{
		"one":          serial.OneStopBit,
		"onepointfive": serial.OnePointFiveStopBits,
		"two":          serial.TwoStopBits,
	}[s]
	return rtn, ok
}

func init() {
	outputs.Add("serial", func() telegraf.Output {
		return &Serial{}
	})
}
