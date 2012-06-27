package gojustext
import ( "io"; "os"; "bytes"; "compress/gzip" )

func AragoneseStoplist() ([]byte, os.Error) {
var gz *gzip.Decompressor
var err os.Error
if gz, err = gzip.NewReader(bytes.NewBuffer([]byte{
0x1f,0x8b,0x08,0x00,0x00,0x00,0x00,0x00,0x00,0xff,0x34,0x92,
0x4d,0x6e,0x1b,0x39,0x10,0x85,0xf7,0xef,0x22,0x1e,0x03,0x1e,
0xcd,0x19,0x0c,0xc3,0x98,0x18,0x70,0xe0,0x00,0xce,0x26,0xcb,
0x12,0xbb,0x24,0x11,0x21,0x59,0x1d,0xfe,0x04,0x51,0x6e,0xe3,
0xa5,0x17,0xde,0x24,0x37,0x48,0x5f,0x2c,0xaf,0xa8,0x64,0x25,
0x75,0x77,0x15,0xf9,0xbd,0x9f,0x45,0xa1,0x05,0x82,0x33,0x0c,
0x67,0xc5,0x97,0xa1,0x18,0x45,0xb0,0x5c,0x09,0x82,0x15,0xfe,
0x1a,0x9f,0x71,0x0b,0x6b,0x68,0xe3,0x2c,0x58,0x6d,0x9f,0x24,
0xc4,0xed,0x67,0x41,0x12,0x34,0xe5,0x8b,0x8a,0x64,0x50,0x48,
0xe3,0x46,0xe6,0xb8,0x2f,0x35,0xdc,0x17,0x3c,0x71,0x65,0xd5,
0x7a,0x88,0x21,0x72,0xae,0xda,0xd7,0x58,0x42,0xe4,0xd9,0x5a,
0x5a,0x5c,0xa4,0xe3,0x24,0xfb,0xff,0x3e,0xe7,0x5f,0x3f,0x76,
0xfe,0x2f,0x76,0x29,0xbd,0xa1,0x0b,0xf2,0x28,0xdc,0x58,0xa3,
0xf9,0x8d,0x86,0x62,0x4e,0xd3,0xe0,0x93,0x68,0xb1,0x0f,0xe9,
0xfc,0x42,0x36,0xde,0x26,0x35,0x08,0xb4,0x75,0xf2,0x10,0xa2,
0xe1,0x50,0xa5,0x84,0xed,0xb5,0x21,0x6f,0x2f,0x0d,0xab,0xd4,
0xae,0x3c,0x31,0xc7,0xed,0xb5,0x50,0x5d,0x15,0x3c,0x11,0x52,
0x56,0xde,0x95,0x78,0xa7,0x93,0x4c,0xfd,0xbd,0x2a,0x96,0xc8,
0xa3,0x9a,0xed,0xf9,0xf7,0x51,0x50,0x35,0x9c,0xa6,0xc8,0xbb,
0x93,0xda,0xb1,0xca,0x61,0x7b,0x93,0xdd,0x44,0xb8,0x21,0x83,
0x1e,0xa9,0x98,0xd7,0xbb,0x53,0x89,0x68,0x59,0x7b,0x75,0x83,
0x26,0x1c,0x97,0xa5,0xe0,0x93,0x72,0x1a,0x5d,0x6b,0x8e,0x74,
0x91,0x1f,0xa9,0x21,0x75,0xa9,0xd4,0x4f,0x4b,0xb8,0x1d,0x87,
0x5b,0x10,0xc9,0xe0,0x98,0x92,0x49,0x61,0xd8,0x47,0xd2,0x04,
0x21,0x68,0xb1,0xec,0x28,0xff,0x30,0x9c,0xec,0x82,0x92,0x05,
0x49,0xd3,0xb5,0x85,0x7a,0x6b,0xe4,0x30,0x2d,0x38,0x50,0x43,
0xcc,0xeb,0x98,0x13,0xe3,0x28,0x95,0xea,0x4a,0x77,0xee,0x45,
0x13,0xc3,0x24,0x08,0x0f,0x3b,0x5a,0xf1,0x81,0xba,0xbd,0x19,
0x08,0xba,0x9f,0x08,0x17,0xfe,0x79,0xe2,0x33,0x77,0x50,0x24,
0x7e,0xf3,0x53,0x6f,0x89,0x1a,0x1b,0xc3,0xf0,0xa8,0x1c,0xcd,
0x5f,0xca,0xe8,0x46,0x22,0x81,0x79,0x98,0x92,0x9c,0xf6,0x7a,
0xc6,0xe3,0xc2,0x58,0x90,0xcc,0xcb,0x61,0x21,0x78,0x8a,0x4c,
0x50,0xea,0xcd,0x7c,0x3f,0xc3,0x21,0x15,0x1d,0x24,0x2c,0xbb,
0x95,0x86,0x16,0x4e,0x7a,0x95,0x48,0x7c,0x25,0x85,0x01,0xe7,
0x38,0x13,0x68,0xd4,0xec,0x79,0x05,0x5a,0x98,0xb6,0x17,0x06,
0x26,0x78,0xd6,0xbf,0x6d,0x60,0x60,0x07,0x9e,0xb1,0x0a,0x3a,
0xc7,0xe1,0x8b,0xde,0xc8,0x05,0xf7,0x6d,0xe5,0x03,0xc3,0x59,
0x87,0x76,0x3c,0x1a,0x0e,0x56,0x09,0x7a,0x27,0x6c,0x85,0x26,
0x76,0xf4,0x21,0x53,0x2d,0xcb,0x14,0x4e,0x83,0x19,0x13,0x7e,
0x32,0x51,0x54,0x1e,0x5a,0xf9,0x43,0xf5,0x17,0x72,0xde,0x79,
0xc3,0xe0,0x23,0x09,0x5a,0xe7,0x2a,0x9f,0xaf,0x2f,0x0e,0xe0,
0x7f,0x8e,0x32,0x14,0xaf,0x50,0xc1,0x83,0x37,0x86,0x22,0xb5,
0xb1,0xed,0x7b,0xda,0xe6,0x5d,0xa0,0x02,0xc1,0x3b,0x1a,0x67,
0x8c,0x78,0x87,0x6e,0x9d,0x8d,0x8f,0xd9,0x3b,0xf7,0xf1,0x4f,
0x01,0xdd,0x66,0xf7,0xa6,0xe1,0xfb,0x25,0xfc,0x99,0x58,0xc3,
0x7b,0x3d,0xdb,0xc2,0x7e,0xfd,0xfb,0x81,0x9c,0x45,0x47,0xdb,
0xe1,0x77,0x00,0x00,0x00,0xff,0xff,0x3f,0x5a,0xa0,0x17,0x96,
0x03,0x00,0x00,
})); err != nil {
	return nil, err
}

var b bytes.Buffer
io.Copy(&b, gz)
gz.Close()
return b.Bytes(), nil}