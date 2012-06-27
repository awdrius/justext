package gojustext
import ( "io"; "os"; "bytes"; "compress/gzip" )

func UrduStoplist() ([]byte, os.Error) {
var gz *gzip.Decompressor
var err os.Error
if gz, err = gzip.NewReader(bytes.NewBuffer([]byte{
0x1f,0x8b,0x08,0x00,0x00,0x00,0x00,0x00,0x00,0xff,0x44,0x95,
0x4d,0x72,0xe2,0x58,0x10,0x84,0xf7,0x3a,0xcc,0xdc,0xd1,0x0e,
0x44,0xb0,0x60,0x33,0x1d,0x73,0x82,0x5e,0xb4,0xb0,0x03,0x50,
0xd8,0x86,0xb6,0xe9,0x85,0xce,0xf1,0x24,0x2e,0x33,0xdf,0x97,
0x25,0xdc,0x0b,0x4b,0xe2,0xbd,0xfa,0xc9,0xca,0xca,0x2a,0x2f,
0x2f,0xf7,0x1f,0xdd,0xdc,0xdf,0xf7,0xcb,0x9f,0x8e,0xef,0x7d,
0xd7,0x86,0x79,0xd7,0xde,0xba,0xf6,0xc1,0xf9,0xf2,0xd2,0x06,
0x0e,0xda,0x07,0x5f,0xf3,0xae,0xbb,0x3f,0x71,0xe6,0xe3,0xfe,
0x5f,0x37,0x6f,0xf5,0x9b,0xb4,0x1c,0x70,0x7e,0xd1,0xf9,0xa9,
0xbb,0xef,0x79,0xb4,0xc3,0x32,0x55,0xa0,0xad,0x11,0xde,0x74,
0x21,0x3c,0x4e,0xed,0x48,0x14,0xb3,0x10,0x76,0xde,0x69,0x9f,
0x1b,0xbe,0xb9,0x22,0x4b,0x3b,0xb4,0x5b,0x1b,0x39,0xf5,0xe0,
0x95,0x07,0x59,0xca,0x22,0x2e,0xd8,0x0c,0xed,0x55,0xdf,0x0d,
0x89,0x7e,0x60,0x42,0x5e,0x0d,0xcc,0x35,0x05,0x91,0xf9,0xf2,
0xc1,0xdd,0xa4,0xcb,0xcf,0x79,0x4a,0x14,0x0a,0xd2,0x79,0x99,
0xd6,0x1b,0xd0,0x2c,0x63,0x01,0xd9,0xe2,0xdd,0xeb,0xf9,0x61,
0xa4,0xcf,0x94,0x2f,0x84,0xa4,0x5a,0x23,0x73,0x71,0xb9,0x3f,
0x71,0x61,0x6a,0xc8,0xd8,0x76,0xad,0xbc,0xb9,0xdf,0x70,0x66,
0xd5,0x1e,0x12,0x3b,0xf4,0xd4,0x87,0x49,0x0f,0x5d,0x7b,0xd7,
0x12,0xfb,0xa7,0x90,0xa4,0xd3,0x81,0xd7,0xab,0xf4,0x0d,0xa6,
0x35,0xd1,0xca,0x55,0x7b,0xd5,0xeb,0x88,0x97,0x08,0x8a,0xe4,
0xdd,0xfc,0x8c,0x71,0x10,0xfd,0xf2,0x96,0x7e,0xac,0x7d,0x68,
0xc0,0x3a,0xf1,0xfb,0x3a,0x6f,0xea,0x9e,0xd0,0x9f,0x04,0x39,
0xa5,0x5f,0x58,0xef,0xf1,0x16,0xe1,0xc3,0x0d,0xfb,0x51,0x5a,
0x77,0xed,0x28,0xe4,0x53,0xfb,0x6d,0xca,0xdc,0x01,0xe2,0x36,
0xf7,0x1a,0xc7,0xc6,0xc6,0x1e,0xad,0xe8,0x22,0x3b,0x1e,0x02,
0xd1,0x1c,0x34,0xa3,0x9d,0x0c,0x8f,0x31,0xa8,0x68,0x38,0x45,
0xd0,0xa0,0x10,0x66,0x02,0x2b,0xe7,0x05,0xdd,0x78,0xcd,0x7c,
0xb7,0x11,0x5c,0x45,0xb6,0xc2,0xe9,0xdb,0x89,0xbf,0x31,0xe1,
0x97,0x11,0xa4,0xcb,0x36,0x5d,0x49,0xf2,0x1e,0x82,0xab,0xca,
0x84,0x89,0x66,0xec,0x9b,0x95,0x4c,0xe4,0xb0,0x92,0x12,0x90,
0xa1,0x26,0x3a,0x09,0xce,0xd1,0x06,0x95,0x52,0xdf,0x52,0x5f,
0x75,0xca,0xca,0x9f,0xdb,0x2d,0x94,0xcc,0x54,0xfb,0x8e,0xbc,
0xc2,0xf3,0x58,0xcc,0xf4,0xd2,0xc0,0x1f,0x50,0xae,0xb1,0x50,
0xee,0x22,0xe8,0xb1,0x50,0xd7,0xf5,0xba,0x70,0xb7,0x23,0xcc,
0xda,0x00,0x29,0xe4,0x41,0xc2,0x5b,0xe9,0x62,0xe0,0xa2,0x57,
0xba,0xc0,0xb4,0xb3,0x8e,0x93,0x62,0x09,0x87,0x00,0x18,0xda,
0x2f,0xef,0x07,0x95,0x70,0xab,0x04,0xd5,0x13,0x7c,0x6f,0xf5,
0x3a,0x07,0x97,0xbf,0x9c,0xbe,0xe8,0x64,0x85,0x08,0x76,0x1c,
0xe7,0xe7,0x07,0xbf,0x09,0x7e,0x86,0xac,0x8d,0x15,0xbd,0x65,
0x22,0x95,0x4b,0xa6,0xcf,0x36,0x26,0x04,0x6a,0xb6,0xc6,0x1a,
0x87,0x92,0xf1,0x3b,0x32,0x1f,0x6d,0xc2,0x9a,0xb4,0xda,0x7f,
0xf5,0x75,0x0e,0x07,0x46,0x6c,0x5f,0x04,0x5b,0x1b,0x5a,0xfd,
0xce,0xf0,0xd9,0x02,0x26,0x29,0x25,0x26,0x82,0xfa,0xf8,0x20,
0xbb,0x1b,0x00,0x22,0x2e,0xca,0x4d,0xde,0xe6,0x0d,0x03,0x49,
0xa0,0xdd,0x82,0x45,0x89,0x42,0x66,0x96,0x7f,0x75,0x22,0x48,
0x4d,0xe6,0x41,0x33,0x13,0x63,0x34,0x6f,0x63,0x58,0x63,0x6b,
0x16,0x4a,0xf2,0xf5,0x9b,0x3a,0x6e,0x8e,0xe6,0x17,0x24,0xda,
0x70,0xe7,0xd5,0xe1,0x09,0x05,0xaa,0xa0,0xf4,0xae,0x40,0x56,
0x60,0x6b,0x62,0x09,0x82,0x85,0xa8,0xa9,0x82,0x87,0x94,0x4f,
0x62,0x29,0x74,0xe0,0x3a,0x3d,0x51,0x55,0x74,0x16,0xad,0x17,
0xff,0x02,0xe5,0xe4,0x8a,0x60,0xe9,0x3b,0x34,0x55,0x51,0xb4,
0xd0,0x32,0xd0,0x60,0x35,0x29,0xd3,0x80,0xad,0x38,0xcf,0xe1,
0x3b,0x0d,0x83,0x66,0x5b,0x5c,0x72,0xcd,0x7c,0xc2,0x9d,0x13,
0x9a,0xc1,0x22,0x2f,0x31,0xd9,0x11,0x7a,0x41,0x62,0x69,0x6b,
0x8c,0xd2,0x29,0x3a,0xa1,0x8d,0xc0,0x67,0x3b,0x87,0x7c,0x2d,
0xf8,0xc1,0xf2,0x00,0x12,0x20,0xd9,0x19,0x7f,0x99,0x5f,0xc3,
0xd2,0x48,0xa0,0x3d,0x56,0x83,0x64,0x94,0x86,0x9e,0x6b,0xe0,
0x11,0x4a,0x56,0x0f,0x99,0xb0,0x76,0x0e,0x15,0xfb,0xc3,0x84,
0x70,0x0c,0x37,0xaf,0xd0,0x9c,0x1a,0x1f,0x43,0x9a,0x55,0xe0,
0x16,0x64,0x68,0xe2,0x7e,0xf0,0x42,0x8c,0x98,0x99,0x3e,0x0d,
0x34,0xfd,0xfa,0x0f,0xa2,0x64,0xf7,0x0d,0x0b,0xc1,0xae,0xe3,
0xca,0xa4,0xa0,0xd1,0x7f,0x94,0xfe,0x57,0xc0,0x4a,0x97,0x84,
0x8b,0x84,0x0a,0xc3,0x4c,0x6d,0x3f,0xae,0xd1,0x43,0x45,0xb6,
0xa2,0x23,0x9d,0x2b,0x5e,0x57,0xf1,0x48,0xef,0x9f,0x22,0xf2,
0x27,0xb7,0x35,0x46,0xb9,0x45,0xd8,0x6e,0xe0,0xf0,0xe7,0xd6,
0x4d,0xd5,0xae,0x3b,0xc5,0x80,0x33,0xb9,0xf4,0x77,0x27,0x59,
0x43,0x44,0x90,0xff,0x1a,0xb5,0x73,0xea,0xe4,0x7b,0xdb,0x57,
0x4d,0xee,0xf3,0x1a,0x23,0x09,0x79,0x17,0xe8,0x63,0x49,0x4b,
0xa4,0x0b,0x82,0xec,0x30,0xc6,0xe9,0x43,0x73,0xca,0x25,0x3b,
0x94,0x39,0xc0,0xa6,0x16,0xdd,0xba,0xd7,0xf7,0xdd,0xff,0x01,
0x00,0x00,0xff,0xff,0xfd,0xdf,0x41,0x97,0x6c,0x07,0x00,0x00,
})); err != nil {
	return nil, err
}

var b bytes.Buffer
io.Copy(&b, gz)
gz.Close()
return b.Bytes(), nil}