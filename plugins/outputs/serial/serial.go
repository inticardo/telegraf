//go:generate ../../../tools/readme_config_includer/generator
package serial

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/albenik/go-serial/v2"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/config"
	"github.com/influxdata/telegraf/plugins/outputs"
	"github.com/influxdata/telegraf/plugins/serializers"
)

const ReadTimeout = 5000

//go:embed sample.conf
var sampleConfig string

type Serial struct {
	Port    string          `toml:"port"`
	Brate   int             `toml:"brate"`
	Bits    int             `toml:"bits"`
	Hupcl   bool            `toml:"hupcl"`
	Parity  string          `toml:"parity"`
	Stop    string          `toml:"stop"`
	Timeout config.Duration `toml:"timeout"`
	Log     telegraf.Logger `toml:"-"`

	serialOutput *SerialOutput
	serializer   serializers.Serializer
}

func (*Serial) SampleConfig() string {
	return sampleConfig
}

func (s *Serial) SetSerializer(serializer serializers.Serializer) {
	s.serializer = serializer
}

func (s *Serial) Init() error {
	s.serialOutput = &SerialOutput{}
	return nil
}

func (s *Serial) Connect() error {
	return s.serialOutput.connect(
		s.Port,
		s.Brate,
		s.Bits,
		s.Hupcl,
		s.Parity,
		s.Stop,
		int(time.Duration(s.Timeout).Milliseconds()),
	)
}

func (s *Serial) Close() error {
	return s.serialOutput.close()
}

func (s *Serial) Write(metrics []telegraf.Metric) error {
	octets, err := s.serializer.SerializeBatch(metrics)
	if err != nil {
		return fmt.Errorf("could not serialize metric: %v", err)
	}
	err = s.serialOutput.write(octets)
	if err != nil {
		return fmt.Errorf("could not send metric: %v", err)
	}
	return nil
}

type SerialOutput struct {
	ser *serial.Port
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

func (so *SerialOutput) connect(
	port string,
	brate int,
	bits int,
	hupcl bool,
	parity string,
	stop string,
	timeout int,
) error {
	realParity, ok := so.mapParity(parity)
	if !ok {
		return errors.New("invalid parity")
	}
	realStop, ok := so.mapStop(stop)
	if !ok {
		return errors.New("invalid stop bit")
	}

	if timeout < 0 {
		timeout = 0
	}

	var err error
	so.ser, err = serial.Open(port,
		serial.WithDataBits(bits),
		serial.WithBaudrate(brate),
		serial.WithHUPCL(hupcl),
		serial.WithParity(realParity),
		serial.WithReadTimeout(ReadTimeout),
		serial.WithStopBits(realStop),
		serial.WithWriteTimeout(timeout),
	)
	if err != nil {
		return fmt.Errorf("error opening port: %v", err)
	}

	go io.Copy(io.Discard, so.ser)
	return nil
}

func (so *SerialOutput) write(octets []byte) error {
	_, err := so.ser.Write(octets)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}
	return nil
}

func (so *SerialOutput) close() error {
	return so.ser.Close()
}

func init() {
	outputs.Add("serial", func() telegraf.Output {
		return &Serial{
			Brate:   9600,
			Bits:    8,
			Hupcl:   false,
			Parity:  "none",
			Stop:    "one",
			Timeout: config.Duration(time.Second),
		}
	})
}
