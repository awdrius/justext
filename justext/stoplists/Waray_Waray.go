package gojustext
import ( "io"; "os"; "bytes"; "compress/gzip" )

func Waray_WarayStoplist() ([]byte, os.Error) {
var gz *gzip.Decompressor
var err os.Error
if gz, err = gzip.NewReader(bytes.NewBuffer([]byte{
0x1f,0x8b,0x08,0x00,0x00,0x00,0x00,0x00,0x00,0xff,0xca,0x48,
0xcc,0xe3,0xca,0x48,0xe4,0x72,0xcc,0xe3,0x4a,0xcc,0xcd,0xe7,
0x2a,0x2d,0x4e,0xe4,0xca,0x4e,0xe4,0xca,0xcc,0xe3,0xca,0x4b,
0x2c,0xce,0x4f,0xe1,0xca,0x4b,0x07,0x4a,0x17,0xa5,0x66,0x64,
0x56,0xe6,0x03,0x85,0xd2,0x13,0xb9,0x00,0x01,0x00,0x00,0xff,
0xff,0xc1,0xc7,0x06,0x06,0x2f,0x00,0x00,0x00,
})); err != nil {
	return nil, err
}

var b bytes.Buffer
io.Copy(&b, gz)
gz.Close()
return b.Bytes(), nil}