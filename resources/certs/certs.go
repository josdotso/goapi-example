package certs

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _insecure_key_pem = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xd7\xa5\x12\xf4\x08\x16\x86\x61\x9f\xab\xf8\x7d\x6a\xab\xc3\x20\x46\x84\x99\x39\x2e\x8c\x1d\xec\xe0\xd5\x6f\xcd\xcc\xca\x3d\xf6\xb8\xa7\x3e\xf3\xfe\xe7\xef\x63\x05\x49\x31\xff\xb8\x1e\xf3\xc7\x76\x95\x90\xf1\x85\x3f\x9a\x90\xfc\xf3\x01\x0c\x45\x51\xb5\x46\x61\x19\x46\xe3\x98\x46\x60\x88\xe5\x0d\x67\x5a\xbc\x6c\xd3\x7e\x21\xd4\xbf\xb7\xac\xf1\x9c\x2f\xea\xa0\x57\xc5\xfd\x20\x02\x4a\x06\x41\xd9\x9e\x8c\x42\x4a\x4f\x1e\x05\x61\x04\x1c\x16\x0d\xef\x6a\x4f\xe9\x75\xe6\x42\x78\x9c\x66\x64\x79\x4a\x41\x56\x58\x38\x79\x37\xe8\x49\x27\xd2\xa2\xa7\x0f\x2e\xe5\x56\x18\x06\x90\xb7\x2e\x28\x0b\xdd\xa2\x2b\x26\x6e\xf7\xd9\x8b\x42\x01\x82\xa0\x7b\xf3\xcb\x27\xa7\x62\xaa\xf9\x6b\xe5\xc2\xa3\x85\xab\xa0\xf0\xa1\x40\x45\x09\xa7\x0e\xc4\xf8\x5d\x9a\x83\xbe\xca\x52\x61\xc5\x0d\x5b\x17\x06\x6d\xd0\xce\x82\x2c\xe6\xf9\xc0\x3c\xf6\x02\xe1\x92\x92\x3f\x0f\x4c\x42\xa8\xc6\x17\x8b\x3b\x08\xeb\x39\xd4\xda\xcd\x7c\x1b\x7e\x93\x69\x1c\xf9\xcd\x14\xd7\xb7\x97\xa7\x37\x2b\x1d\x4c\xb2\x8b\x28\xa9\xeb\xd6\x1a\xd3\xb1\xca\xa2\x30\xab\x00\x85\xb6\x3e\xb8\xb6\x79\x8d\x34\xb9\x7d\x73\xd1\x94\x12\xbb\xa7\x87\xca\x63\x56\xeb\xe9\x46\x2a\x90\xbd\xb2\xdd\xcf\x6e\xd4\x1b\x73\xbf\xf3\x8e\x2c\x4a\x40\xc6\x75\x99\xaf\x44\x19\x39\x85\x01\x02\xaf\x58\x5c\xc8\xad\x45\x24\xb6\x66\x78\xbb\xc4\xcd\x87\xb8\xb5\xaf\x0d\x62\x12\x0b\xb7\xfe\x7b\x26\xa9\x5b\x69\xc4\x2d\xd6\x49\xc9\x78\x04\xf4\xed\xe4\x96\x3f\x37\x1a\xdd\xd6\xe8\x64\x84\xe3\x01\xf4\x6e\xd8\xaa\x87\x5c\x39\x3e\x4e\x0d\xf4\x2d\xa4\x13\x3b\x5b\x54\x0e\x1d\x8c\x67\xd0\x00\x15\x65\x04\x6b\x29\x81\x68\x59\xef\x27\xb8\x2c\x65\xba\x2d\x2d\x2c\x9c\x58\x48\xdb\xc4\xa7\x06\x69\x03\xbb\xf6\x01\xe1\x3e\xc6\xab\xd1\x4d\x4a\xfa\x5c\x0e\x6d\xa7\xdf\x78\xed\xb0\xfc\x28\xeb\x3b\x16\x64\xb1\x79\xe8\xba\xfc\x66\xd4\x66\x21\xc5\x3c\x1f\x3f\x3a\xde\xd5\x7b\xe5\x13\x4b\xe1\x94\x30\x02\x94\x04\x1b\x3c\x39\x77\x17\xfc\xc9\x2c\x32\x30\xa8\x5d\x37\x97\x88\x41\xac\xfc\x20\xda\xa5\xa1\xa6\x26\xd9\x95\x34\xb8\xf4\x58\x15\x5d\xe1\xdc\x79\xc1\x4f\x7b\x8f\xce\x21\x8c\xe1\x0d\x1b\x71\x80\x3d\xc8\x68\x0f\x9b\x7e\x8f\x90\x99\xf1\x4f\xb4\x84\x46\xe4\x28\x02\xde\xf1\xe6\x24\x22\x67\x20\xd2\x79\x27\xab\x49\x8b\xb4\xcb\xc4\x88\x93\x4d\x44\x7e\xda\xf5\x2c\xde\xc5\x0d\xc5\x54\x6d\x86\x00\x05\x09\xbe\xbd\x65\x84\xcd\x89\xd0\x88\xb6\x4a\x06\x83\xc2\x63\x78\x33\xca\xac\xeb\x89\xb1\xc9\x9d\x05\x65\xea\xfe\x9b\x11\x24\x30\x1a\xba\x32\x3a\xf8\x49\xaa\x3c\xb5\x0b\x88\xe2\x98\x4b\x60\x00\xc6\xf9\x77\xc2\xa3\x05\xf6\x0a\xf7\x69\x60\xf9\x59\x24\xf3\x73\xc5\xaa\xa4\x2e\xd5\x9c\x60\xdb\x5b\x22\x73\x8b\x19\x5b\x55\xf9\xa4\xc6\x3f\x7a\x20\x6d\xba\x0e\x96\xfb\x88\xa3\x1f\xf2\x05\x56\xa5\x28\xa6\x90\x62\x7c\x7a\xb4\x4b\x28\xc9\xad\x0c\xca\xab\x69\xec\x6c\xcb\x4e\x1f\xef\xa8\x9a\x74\x13\x66\xf0\xf8\x54\xbd\x23\x6e\xd2\xdb\x91\xa3\xaf\x9f\xca\xd9\x86\x74\xe7\x1b\xbc\x01\x03\x5f\x49\xe2\x50\x27\xd2\x30\x6d\x50\xbb\xa1\xc6\x08\x6d\x37\x6b\x1e\x69\x98\xd0\xa9\x10\xcf\x39\x6b\xf1\x9a\x27\xe3\x07\xd9\xe3\xd0\x4a\x88\x23\x0b\x22\x3a\x7d\x43\xe1\xb5\x5b\xae\x7d\x29\x90\x04\x94\x06\xde\xb8\xcc\x99\xf1\xe8\x25\x20\x33\xec\xfc\x2a\x96\xdf\xa6\xd6\x1a\x9c\xe0\x70\xca\xb4\xe1\x91\x74\xfd\x13\xa6\xc6\xfe\xe1\xa2\xc4\x88\x3d\xdb\xdf\xb0\x57\xa7\x3b\xbf\x3f\x53\x3d\x2b\x80\xe1\xe5\x52\x10\x2e\x38\xed\xd1\x94\x11\xb6\xf7\x81\xe4\xc5\xee\x61\xd6\x00\xff\xbc\xad\xae\x0c\x69\xcf\x1b\x25\x6b\x06\x9e\xed\xe7\xfe\x5b\xed\x61\xdf\x38\x4b\x93\x3b\xc9\xc1\x7c\x3f\xc4\x81\x03\x79\xf6\xd1\xdf\x0c\xe4\xdd\xc7\x86\x7c\xbe\x8f\x66\x55\xc9\x3b\x03\x21\xa4\xd8\xc9\xe6\x1f\x88\x8d\xd6\x12\xc4\xb7\x0a\x05\x6d\xf6\x13\xb7\xf2\xcb\xcd\x6d\x27\x25\xfe\x73\x42\x55\x8d\xef\x6d\x08\x20\x6a\x82\xbb\x92\x5f\x76\x56\x32\x06\xc9\xb1\xd5\x77\xed\xce\x4e\x27\xeb\x19\x15\x3b\xee\xc6\x06\x2b\xd3\x5e\x6c\x1c\x38\x33\x21\x5f\x3e\x9d\x16\x38\x6d\x86\x4a\x0c\xb7\x2c\xc1\x7d\xc4\xe3\x01\xe4\x91\x62\x9d\x68\x5a\x72\x87\xd2\x8d\xa2\x4d\xd7\xea\x4a\x8f\x1d\x5e\xde\x10\xb3\xb9\x40\x40\x6e\x11\xf6\xf4\x82\x66\x09\x6a\xee\x43\x8f\xe5\xb4\xad\xe4\x74\x42\xed\x35\xd9\xaf\x8f\x5c\x4f\x81\x3b\x2e\x05\xf8\xd5\x23\xe9\xd4\x07\x66\xef\x1c\x6c\xdf\x24\x9d\x80\xb6\xb5\xec\x4e\xea\x58\x37\x8a\xe9\xf4\xef\x13\xa4\x49\x5f\x7a\xe8\x16\xfd\x9e\x8b\xf1\xc0\xc0\xdd\xd5\x31\x49\xe1\x72\x88\x01\x4d\xbf\xb9\x12\x3d\xd2\xac\x54\xed\xde\x14\x1c\x99\xdd\x4c\x8d\xfe\xa8\x88\xc2\x24\xbf\x33\x9a\x74\x4a\xe8\xf2\xd7\xd9\xc6\xc7\x28\x22\x9e\xd1\x1a\x27\x80\x95\x52\xf2\x4a\xa6\x2f\xc4\xa3\x0b\x80\x09\x2f\xc2\xf6\x9e\x45\x62\x58\x75\xbf\xf3\x9d\x21\x14\xa4\x0a\x41\x6f\x8f\xe3\x11\x39\x1a\x1a\xe5\xd3\x7f\x36\x27\x8c\xbe\x3a\x38\x20\x0d\x27\x38\x9e\xc0\x35\x8d\xc0\x32\x76\xed\xa8\x04\x35\x00\x99\x47\xf9\xd5\x3e\x18\x61\x64\x33\x6f\xee\x6c\xb1\x16\x37\x08\xec\xf4\x86\x4b\x0c\x5b\xfb\xa5\xa6\x6e\x27\x7a\xa4\xc5\x05\x0a\x13\x6e\x21\xa6\xc0\x4e\xaf\xa5\xcb\xb4\x62\x18\x7a\xa6\xd3\x3d\x01\x44\x19\xf8\x79\xe7\x39\x50\xc7\x48\x8f\xbc\x3e\x70\x41\x0b\x3a\x13\xd7\xe8\x16\xf3\xa4\x38\x95\x6a\xa5\xfd\x02\x0d\x17\xd4\x09\xaa\xc2\x11\x82\x9a\x99\x1f\x4e\x6a\xc3\x13\x53\x8a\x59\x8f\xef\x00\xd4\x58\xe8\xc6\xc6\x0f\xc7\xc3\x08\x44\xc6\xee\x40\x54\x83\x8b\x18\x3e\x84\x1f\xc9\xee\x57\x10\xdd\x01\x4b\x42\x1e\xda\x95\x6f\xd7\xb4\xfa\xdb\x99\xd8\x47\x89\x92\x75\x59\x20\xc8\xdf\xc2\xab\x04\x6e\x3a\xc1\x05\x0f\x35\x34\x1c\x45\x98\xb5\xc7\x92\x76\x32\xd2\x74\xd1\xe8\xcf\xe8\xae\x6b\x8c\x08\x09\xa4\x12\xef\x4c\xaa\x65\xa1\xbc\x02\xda\x0a\x19\x61\x6d\x10\x3d\x04\x65\x8c\xed\xb0\xc3\x03\x96\xfd\xba\x64\xc6\xba\x21\x6c\x28\x99\x11\x41\x0f\x14\x1b\x9c\x12\x09\xe5\x8e\x94\xcb\x94\x21\x72\x5f\x63\x42\x73\xd1\xcf\x40\x96\x3c\xb3\x49\x2b\xae\xa1\x21\x42\xa9\xaf\xbf\x60\xd5\xf4\xd2\x80\xd6\xb7\xaa\xa8\x9a\xeb\x62\xab\x92\x42\x18\xff\x13\x56\x4b\xae\x54\x40\x5b\xd2\x0f\xd1\x74\x5c\x93\xc5\xe4\x5e\x3c\x2e\xb4\xb0\xdd\x33\xcd\x68\xe2\xa9\x98\xc2\x57\xa5\x37\xed\x58\xc0\x88\x44\xc3\xad\x8a\x1a\xff\x3c\x05\xee\xc5\x8d\x77\xa3\xc1\x1d\x49\xe6\x04\x5e\xf8\x8e\xfb\xbf\x3e\x8a\x7e\xf8\x6c\xb6\xa2\xd8\x66\xaf\x85\xda\xf3\x48\x9b\xdd\x1a\x72\x08\x96\xbf\xa1\x36\xa6\x40\xc6\x1e\x74\x06\xe1\xb8\x7e\x51\x9d\xf5\xf0\x43\xd6\x4d\x85\x68\xdf\xf0\x6b\xb4\x99\xcd\x93\x66\x77\x08\x97\x62\xe2\x4c\x89\x17\xe5\x61\x1f\x64\x4c\xbf\xeb\x18\x66\xc7\x91\xe9\xc3\x0c\xd2\x0d\x30\xcf\xbc\xa3\x1d\xf1\x55\x26\x87\xdf\xb0\xfc\x0e\x7d\xc0\x9b\x8e\x2a\xe3\xab\xea\x9e\x7b\xe3\xe8\xa2\x92\x3c\xf3\xd2\x05\xf3\x01\xe9\xc2\x56\x2c\x9c\x12\xd7\x02\xee\xad\x40\x2b\x63\x86\xd0\x77\x40\x91\x2a\x89\xfd\xfa\x28\x95\x6e\xa0\x82\x19\x8c\xcd\xe1\x9b\xad\xa8\xd1\xa7\x8a\xaf\xe6\x56\x1f\x36\x29\x9e\xfe\xcb\x4d\x74\x7f\xb3\x39\xea\x84\xf1\xd1\x82\x56\x0d\xe3\x86\xa2\x8a\xd8\x33\xb0\x40\xc8\xc1\x09\x6c\xb7\x78\xf0\x54\x39\x66\x7a\xe3\xf7\x15\x54\x63\xa6\x66\x57\x6d\xee\x44\x13\x8f\x5e\xc2\xb1\x98\xd4\x46\x5f\x0a\xfe\x06\x66\x74\xd6\x51\xbd\xf3\x0d\x3f\xfa\xda\xba\xc8\x8f\x04\x0c\x83\x09\xc7\xf5\xf7\x42\xc3\x88\x7f\x73\xae\xb8\x2f\xd0\xdd\x38\x96\x9d\xea\xd7\xee\xa1\xa7\x05\x3b\xac\x40\xdf\xd9\x94\x3f\x3c\x6b\x0a\x6d\x9c\xb3\x66\x4a\x6b\x84\x60\xcb\x1a\x21\x41\xa2\x07\xbc\xef\x6f\x85\x78\xe7\xdb\x6d\x6e\xd1\xf2\x1b\x71\x78\x35\x6d\x07\x2b\xfe\x43\x02\xab\x2b\xb8\xb7\x02\x6d\xab\xdd\x8f\xd0\x3e\xa6\xdb\x44\x18\x50\x8a\xbe\x85\x7f\xcc\x9b\x9f\x06\xb5\x93\xa9\x0a\x00\x91\x23\x0b\xed\x6d\x73\xb4\x0e\xb9\xf9\x73\x0b\xb2\x2e\xb4\x41\xc6\x42\xbc\xcc\x77\x26\xc8\x0c\xfc\x18\xb8\xe7\x58\x36\x24\xab\x69\x73\x4a\x82\x93\xb2\x6f\xa1\x38\x71\xaf\x28\x0d\xef\xe2\x09\xf0\x52\x3a\x1a\x73\x1d\xc1\x0c\x04\x4e\x52\xbe\x4d\xe9\x75\x04\x12\x4d\x24\xbb\x43\xde\x5f\x06\xe3\x5d\x77\x30\x6c\xd2\x8a\xca\x6b\xb7\x77\xe6\x4e\x9e\x3f\x6d\x9e\xe4\xc1\x67\xce\x49\x88\x15\x09\xc8\x21\x98\x98\xc2\x1a\x53\xd9\xc0\xff\x9e\x4e\x0b\xa5\xf9\xc1\x35\x11\x73\x29\xb2\x95\x33\x3c\x3d\x73\xe6\xe9\x56\xe5\x3a\xbc\xc2\xe6\x3d\x76\xa6\x0f\xa2\x7a\xf1\x47\xbf\xd2\xa3\x8e\x69\x3c\x07\x14\xec\x57\xdf\x2f\x8d\x63\x1c\x81\x39\xa8\x9f\xf8\xad\xea\x25\x53\xd6\x9d\x3f\x16\x75\x31\xa9\x2c\x2f\x90\x1d\x14\xcb\xbe\x76\x8d\xe5\xa1\xf9\x31\x75\x9f\x76\x66\xb4\x57\x8e\x31\xdc\x05\xf1\x02\x10\x87\x4b\x26\x67\xd5\x12\x1a\x3c\xc0\x7f\x86\x1f\xe8\x73\xab\x04\x07\x14\x07\x98\xaf\x09\xd5\x74\xb8\x5a\x96\x15\x99\x67\x7b\x6c\xee\xfe\x5c\xcc\x26\x6a\xc2\x32\xd4\xd7\x2e\xcd\x08\x84\x53\x16\x10\x93\xb3\x9f\xae\x1a\xe4\x4f\xb3\xb2\x30\x2f\x5e\x0b\x7d\x99\x92\x6b\xef\x62\x15\x2c\xc3\x66\x18\x3a\xdc\xd9\xd5\xe8\x98\x1d\xc5\xfe\xbd\xe7\x02\x25\xc6\x13\x62\x6c\xe1\x0d\x08\x57\x83\x3b\x0f\x44\x99\x7b\x7e\x3a\xd8\x87\x50\x0b\x62\x94\x31\xef\x08\xf6\x77\xdb\x7d\x08\x19\xda\xf5\x95\x89\x98\x5d\xc7\x6f\x43\xf7\x0f\x64\xe6\x0a\x47\x91\x6a\x5b\x73\x6e\x45\x99\xa1\xa8\x28\x97\xc2\x2f\x1e\xf0\x1b\x8d\xa9\xbc\x45\x6b\x90\x03\xda\x57\xe0\x30\xd7\xa9\x96\x31\xf8\x8e\xca\x59\xf1\xe5\x03\xdc\x29\x03\x50\x65\x2a\x5b\x21\xdf\x07\xe4\xa6\x24\x49\xfa\x4f\xcf\xc3\x4d\x14\xe4\xc1\x6f\x2b\x7b\x60\x93\x5d\x59\x30\x71\x31\xb8\xba\x2c\x8b\x1f\xda\xa8\x5c\xe5\xd1\xe0\x52\x72\x9b\x6e\xd9\xa1\x24\x74\xfe\x05\x46\x90\x66\xc6\xb6\x8f\xb6\x1c\xbf\xb4\xba\x2a\xc8\x69\xb8\xea\x0d\x73\x0e\xe8\x3e\x60\x77\xf7\xda\x76\xae\x65\x7b\x0f\x70\x8a\x4c\x15\x45\xfe\x98\xab\x81\xb4\xe8\x81\x37\xd3\x33\xc8\x87\xbe\x75\x3c\xd3\x58\x36\xb7\xcd\xd3\x0e\x72\x09\xca\xca\x2b\x76\xd4\xb0\xcf\xdc\x1f\xa0\x3f\x22\x2e\xcc\x1e\x34\xf5\x1a\xae\xed\xa0\x82\x69\x15\x5f\x24\xcd\x6f\xa1\x85\x8f\x26\xfd\x0e\x8c\xad\xe3\x77\x29\xf7\x10\x95\x24\xad\xbc\xc3\xea\x00\xc9\x8e\xf8\x59\x38\xd5\x07\xc2\x29\x92\x80\xcd\x81\x21\x72\xbf\x83\x58\x1e\xc5\x84\x52\xaf\xae\x45\x65\x16\x53\xaa\x1a\x8a\x87\x66\x3b\x65\xc0\xb6\xec\xa9\x1d\x52\xf9\xdd\xf2\xb4\xfb\x0d\x4c\xb4\x36\xf6\xcd\xe4\xad\xea\x57\xa1\xbd\x36\xc0\x26\x3e\xed\xad\x92\x95\x5a\x37\xf2\xfe\xb8\x65\x03\xd2\xdb\x58\x86\x54\x5b\x7c\xce\xcf\xfb\x7a\x1d\x62\x91\x5c\xef\xd4\xd4\x09\xd1\xa6\x8c\x2f\x45\x18\xf0\xc8\x80\x6a\xbf\x6b\x6e\xa0\x51\xf2\x80\xa6\xe5\x1e\xb1\x37\x2f\xf7\x06\xc1\x5d\x6d\x64\x77\xd7\xb0\xc3\x4b\x73\x57\xeb\xa4\x35\xc2\xac\xd6\x69\xe3\x28\x76\xab\x1e\x46\xca\xee\xa3\xde\x0a\x6b\xb5\x3d\xfb\x28\xe7\x11\xbe\x95\xf3\xd7\x5f\xc0\x3f\xd9\x21\x98\xfc\xff\xcf\x91\xff\x06\x00\x00\xff\xff\xb9\xb5\x48\xfd\xaf\x0c\x00\x00")

func insecure_key_pem() ([]byte, error) {
	return bindata_read(
		_insecure_key_pem,
		"insecure-key.pem",
	)
}

var _insecure_csr = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x95\xb7\xce\x83\x6c\x93\x40\x7b\xae\xe2\xeb\xad\x15\xd9\x40\xf9\x90\xc1\x04\x93\x43\x07\x18\x93\x73\x78\x0c\x57\xbf\xfa\x5e\x69\xa5\xfd\x8b\x7f\xda\x33\xa3\xa9\x8e\xce\xff\xfc\x3b\xbc\xa4\x68\xd6\x3f\x82\xe4\xfa\x9a\xac\x09\xc0\x97\xfe\x71\x25\x27\x90\x3c\xff\x8f\x22\xa6\xa6\xc9\xb2\x2f\x08\xe0\xc4\x04\xe0\x00\x58\x65\xda\x4f\xb8\x81\xce\x57\x56\xc8\x83\xc4\x07\x7d\xe8\x9b\xae\x09\x25\x27\x11\x43\xc7\xd1\x24\x38\x8b\x49\xf4\x9b\xd3\x81\xbb\xf2\xa1\xaf\x11\xd3\x4d\xa0\x0c\xfe\xa0\x2a\x41\xdc\x4f\x22\xaa\x72\x47\xbd\xce\x07\x6b\x2e\x08\xeb\x34\x5d\x0c\x2a\xf0\x8f\xbf\xa4\x9f\xab\xe7\xa3\xd5\x27\x64\x78\xa5\x1e\x2f\xa6\xb1\x8e\x21\x59\x94\xce\x09\x21\x63\xa9\x2f\xed\xa6\xb0\x29\x00\x0f\x24\xe1\x67\x76\xee\xc0\x5d\x9a\xcc\x5f\x59\x9c\xd6\x1f\x25\xa8\x7c\x85\x6b\x93\x08\x56\x0e\xc1\xed\x79\x1c\x1e\x59\x64\xd5\x1f\xa5\x3f\x91\xbc\xe1\xdf\xf9\xf0\xa3\x4d\xb7\x82\x72\xf5\xf7\x49\x94\x20\xb7\xe5\x84\x55\xe7\x4a\x7d\x16\xa4\xf3\x34\x7d\x00\x6d\x11\xc0\xaa\xd2\x1a\x13\x60\x8a\xe0\x2d\x8a\xa7\xe5\xa4\xe8\x48\x3c\x02\x9c\x00\x00\x4a\x13\x44\xf8\xb7\xf0\x02\x93\x26\x00\x47\x5c\x46\x2b\x6a\xb1\x18\xb0\x18\xfa\xf6\x3f\x6f\x65\x5f\x78\x5d\xc8\x3f\xe2\x97\xa7\x5e\xa4\xbf\xba\x8d\xe3\x34\xa7\xbe\xde\x29\x82\x6b\x65\xe0\x78\x21\x50\x3f\xb1\xfc\xfc\x8d\xc4\x52\xe9\x71\x14\x95\x6d\x36\xea\x98\x53\x54\x6a\xd8\xda\xa7\xaf\x0f\xdf\x8d\x21\xbe\x68\x97\x6e\x74\x18\xfa\xee\x6b\x19\x0b\xd5\x97\x42\x29\xac\x11\xc5\x78\xfc\x7e\xb0\x76\x3d\xc3\x3a\xf9\xf7\xf1\xb9\x3e\xec\x5b\x98\xf4\xe8\xa9\x86\xb4\xfb\x7b\xbb\xce\xf8\x9d\x79\xce\x09\xbf\xf6\x20\x30\x24\x8f\x63\x95\x14\x9d\xcd\xd2\x89\xa5\x58\x6a\x8c\x43\x21\xa2\x81\x06\xef\xd6\x8a\xfa\xf1\xc0\x0d\xba\x8e\x5d\x74\xe8\xa8\x17\x33\x31\x2f\xaa\xe7\xe4\xd9\x44\x63\xab\x29\xa3\xc0\xde\xb0\x68\xb5\xcd\xf2\x6b\xcd\xb8\x58\x25\xdc\x2f\xad\xb9\x87\xf4\x0c\x07\x24\xa4\xe7\x34\x9c\xa9\x16\x63\xd0\x6e\xd9\x5f\x20\xfb\x28\xb6\x28\x92\xb0\xfe\x28\x9c\x51\x94\xe1\xcc\xee\xc3\xb9\xb5\xce\x63\x52\x5e\xd8\x03\x8c\x6a\xad\x64\xcd\x9d\x75\xb5\xbf\x7f\x38\xfc\x58\x91\x19\x4f\x79\x78\x33\x66\x78\x8b\xb9\x34\xe7\x67\x33\xaf\x83\xd2\x97\x22\xba\x4a\x4b\x82\x32\x55\xa2\x46\xf2\xdb\xe6\xea\x5e\xa6\xd6\x55\x0a\xb9\x1a\xe3\x8d\xd5\xcb\xb4\x52\xb2\x1f\x27\x59\x1e\xc8\x9c\xb1\xc0\x63\xb4\xe3\xe5\xed\x8c\x71\x4c\x9a\xf5\xe9\xee\xaf\x05\x33\xb4\x79\xa8\xc5\x27\x14\x2b\x53\xfc\xf8\x45\xf0\xcd\x1b\x15\xfa\x8b\x24\xbf\x30\x57\x52\x21\xae\xa8\x98\xd7\x69\x21\xcc\x90\xa3\xb0\xba\x7b\x7b\xc0\x15\x65\x22\xeb\xdb\x3f\x43\xb9\x26\x89\xc7\x40\x2d\x37\x69\x95\x8b\x51\x1f\x34\x8e\xca\x85\x53\x44\xa2\x81\x73\x44\xba\x5e\x1b\x93\x5e\x4d\xf3\x24\xf1\x12\x8e\xca\x64\x21\x15\x55\x69\x75\x98\xd4\x6d\xe3\x6a\x9f\x5d\x19\x47\x7d\x62\x76\xe7\xbe\xe0\x46\xf4\x89\x98\xd0\x07\xb3\x28\x83\x78\x95\x7c\x03\xeb\xde\x07\x7b\xd1\x85\x92\xf7\x80\xa6\x6b\x0d\x96\x81\x9f\x3e\x52\x01\x13\xa2\xa9\xf0\x73\x57\xdc\x68\x4a\x52\x73\x27\xbc\xc4\x02\xca\x03\x81\x71\xaf\xd1\x3d\x08\xe8\x9c\x0a\xfe\x3c\x34\xf7\xfe\x2c\xe6\x64\x05\x64\xbb\xb0\x72\x15\xd8\x25\xdc\xe7\x12\x46\xc8\xf9\xad\xba\xa5\x75\xda\xd4\xcb\x4c\xea\x17\x09\x68\x4e\x26\xcb\x94\x40\xf1\x1b\x85\xb1\x04\x9a\x6a\xdb\xab\x6b\x2b\x35\xc6\x9d\x86\x8b\x68\xf3\xd4\xb9\x44\x9c\xba\x8c\x58\xaf\xe9\xbd\xe3\x6e\xc4\x87\x9a\x08\x1c\xc0\x4f\x02\x06\x5f\x30\xd1\x5f\x53\xaa\xd5\x67\x61\x01\xa7\xb3\x4d\x97\x82\xaa\x08\xb2\x7f\xd5\x56\x5d\x49\x92\x6e\xe0\x56\xd5\x7f\x7a\x82\xfc\x9f\x28\xa2\xf3\xff\x8f\x25\x8b\x77\x80\x58\x55\x1a\x0f\x5e\xb0\x1c\x98\x00\x93\x9d\x2f\xba\xdb\x5c\x8d\xaa\xbd\x32\x24\x6f\x83\x20\x25\x32\x9b\x22\x03\x39\xef\x0a\x8b\x7e\x2f\x7b\x29\x38\x4f\x7f\x9c\x82\x62\x98\x7a\x76\xd1\xf8\xd0\x3e\x92\x34\xf4\xd9\xa0\x18\x71\x31\xda\xc1\xd6\x79\xb1\xec\xe8\x3f\xba\x5f\xc5\xb6\x92\x31\x12\x76\x53\xbf\x54\x09\xa2\x04\xb2\x70\x0f\x91\xf4\x78\x0f\x93\x6c\xbe\xaf\x35\x52\xb9\x8e\xe6\x01\xe8\x8d\x7e\xe4\xec\xde\x3c\x57\x45\x8b\xd9\x37\xa9\x09\xa7\x7e\x59\xd3\x62\xc8\xe7\xc7\xa6\x58\x03\xdf\x3a\xd3\x18\x74\x44\x51\x0d\x63\xbd\x6f\x0f\x90\xe3\xfe\x56\x9c\xc0\xb0\x40\x43\xf0\xc7\xe0\x1b\xd6\x8f\x70\x56\xbf\xb8\x6d\x46\x28\x9f\xe3\x45\x26\x1d\x89\x39\xc5\xf9\xb2\xa7\xc6\x07\xab\xd2\xa7\x02\x98\x9e\x2c\x02\x6b\xb1\x0c\x4d\x19\x23\x75\x39\x0c\x8f\x47\x0f\xdc\x92\x3f\xcc\x9e\xda\xdf\xbf\x60\x58\x48\x15\x7e\x1d\x65\xed\xe5\xa0\x88\xd9\x37\xbb\xa2\x03\x81\x87\xe9\xa1\xfc\x7c\x69\x56\x1e\x81\x7a\xf0\x48\xa3\x0c\x42\x83\x97\x19\x86\xca\xcc\x2b\xc5\x66\x5a\xe2\x7b\x0a\xb2\x3c\x17\xec\x73\xd3\xe3\x23\xca\x82\x80\xfa\x1d\xd1\x56\x3f\x46\x6d\x49\xf9\xc1\x64\x56\xb7\xc6\x15\xb4\xae\x9e\x5f\xf6\x83\x22\x63\x3a\xb4\x3a\x71\x15\x65\xa6\xe3\x41\x1f\x0e\xde\xce\xbb\x21\x6a\xe3\xb1\xe2\xda\x31\xdb\xd8\xf9\x11\x34\xc5\x3b\x64\xf6\x8b\x13\xfc\xa2\x79\x1b\x2a\xf1\x25\x16\x40\x1b\x04\xb0\x54\x8a\x7a\x22\x96\xd7\xde\xed\xd3\x8c\x9d\x0e\x1f\x9f\xb6\x29\x6c\x25\x27\xeb\xfd\xd0\x5e\x44\xe2\x90\x93\x16\x25\x79\xd1\xdc\x06\xeb\x50\xd9\x73\xa7\x58\xe7\x6a\x7f\x96\x84\x3d\xb0\x74\x25\x27\x5e\xac\x28\x19\x09\x55\xde\x2a\x49\x3a\x1f\x84\x4a\x5e\xe2\xa2\x12\x23\x5d\x0a\x4a\x2f\x8d\xf9\x81\x61\xce\xbe\x56\xe4\x61\xbb\x3e\xef\xf7\x3c\x9c\x8d\xfd\xd1\x1f\xe9\x34\xd0\xf8\x37\x56\x9e\x53\xef\xb6\x0f\xad\x44\x3a\x17\x6e\xb3\x72\x60\x9b\x94\xc6\xf9\xb2\xab\x65\xb6\x77\x1c\xd7\x9f\x5e\x67\xef\xa7\x40\xf3\x5e\xfc\x10\x37\xb7\x72\x0d\xd2\xca\x32\x92\xbd\x86\x20\x04\xec\x46\x61\x92\x21\xcb\xf8\x4c\x9f\x88\x4a\x3b\x01\x41\x46\x65\x9e\xbb\x9c\xc4\x68\x89\xae\xd1\xcf\x62\xff\x4e\xc5\xfe\xb8\x80\x91\x6e\x4c\x4e\x10\xa5\xc2\xb3\xc7\x29\xb6\xde\x02\xe7\xda\xd2\xa2\xb3\xa7\x59\xda\x99\x44\x0d\xcf\x00\xe2\x1f\xe9\x77\x1d\x34\xae\x09\xec\x11\xf9\xab\x9d\x64\x89\xff\xbd\x84\xff\x1b\x00\x00\xff\xff\xb7\x99\x76\xa7\x2e\x07\x00\x00")

func insecure_csr() ([]byte, error) {
	return bindata_read(
		_insecure_csr,
		"insecure.csr",
	)
}

var _insecure_pem = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x96\xc9\xae\x83\xba\xb2\x86\xe7\x3c\xc5\x9d\xa3\xab\xd0\x25\x84\xa1\x0d\xa6\xc7\x04\x42\x3f\xa3\x49\x20\x84\x2e\xf4\xe1\xe9\x8f\xd6\x5a\xfb\x0c\x8e\xb6\x87\xf5\x97\x55\xf6\xaf\xfa\x4a\xf5\xff\x3f\x07\x22\x45\xc3\xff\x27\x22\xd7\xd3\x64\x4d\x04\x1e\xfa\x8d\x12\x96\xa6\x29\x7a\x2d\x8a\x10\x5c\x4a\xb0\x69\x10\x94\x9a\x16\xb7\xaa\xe0\x89\x7e\x49\xdd\xb5\x4d\x72\x62\xdd\xe8\x13\xad\x5a\x73\x0c\x1c\x84\xa1\x03\xb6\x32\xd5\x76\xf1\x00\x3a\x2c\x71\x40\x40\x10\x7b\xa0\x09\x3c\xcb\xb5\x36\xe4\xc4\x52\xe0\x38\x1a\xda\x06\x29\x0e\xf7\x21\x69\x85\x6f\xd6\x36\x95\xe5\xc6\x9b\x0c\x7e\x35\x15\x6d\xb4\x17\x87\x5c\xe9\x76\x7a\x95\xb5\x78\x20\x72\x06\xaf\x96\x4b\x6d\xca\xf6\x9b\x60\xa0\xdd\xd5\xb3\x0e\x37\x31\x1b\x7c\x93\x3b\x94\x92\x48\xa7\xd2\x30\x19\x62\x46\xa6\x12\x0f\xcd\x96\x38\x29\x80\xf6\x91\xb8\x5b\x6f\xb7\x15\xbe\x84\x26\xc3\x6f\x1a\x25\x55\xa1\xf8\xa5\xa7\x08\x75\x1c\x6e\xa5\xc3\x08\x73\x16\x05\x4b\x1a\xe2\xaa\x50\x9a\x35\x7b\xc1\x5b\xd6\xee\x67\xcb\x2d\x37\xb9\xfc\x2d\x24\xa1\x4d\x98\x32\x06\x57\x44\xa6\x54\x6b\xce\x3a\x17\xcb\x03\x9b\x2d\x81\x4d\xad\x72\x6c\x79\xf9\x86\x3d\x40\x59\x52\xbc\xe1\xc3\xd9\xc2\x7f\x62\xb6\x04\x0e\xcb\xd3\x76\xab\x76\xb6\xb0\x16\x61\x5f\x23\x93\xb0\xc0\xfb\xf7\x49\xb0\xb2\xc4\x20\xb0\x76\x74\x00\xf7\xc7\x1c\x08\x4a\x4f\x7c\xe3\x2a\x53\x9a\x36\x63\xf5\x25\x0d\xd1\x2e\xd7\xc0\xff\xd3\x72\x4f\x0a\x70\x95\xbd\xa0\x42\xe4\xad\xbc\xc4\x4c\x73\xc4\xcc\x75\x57\x3d\x90\xfd\x25\xf4\x9e\x8c\x9a\x25\x67\x82\xba\x88\xf4\x46\x43\xb8\xc9\x3b\x77\x48\xda\xa6\x8e\x23\xb7\xb1\xee\xd4\x66\xfc\x79\x66\x12\xe8\xeb\x2a\x19\xab\x95\xbe\xaa\x0f\x45\x2b\x53\xc9\x1d\x5a\x3f\x7f\xcb\x44\x28\x65\x0c\x3d\x17\xe1\xf9\xd7\xc0\x34\x14\x16\x0d\x09\x4b\xa6\xbe\x77\x45\x02\xe1\x5f\x21\xcb\x23\x24\x66\x5f\x63\x46\x9e\x52\x45\x38\x0a\xa9\xdf\x2d\x09\x70\x96\x24\x8a\xe0\xf5\xaf\x1e\x80\xd0\x01\x52\x59\x6a\x37\xf0\xa3\x97\xbd\x58\x96\x1a\x24\x80\xfd\xc9\xe9\xf4\xe6\xe6\xd2\xe1\x5d\x05\x4c\x5d\x53\xa6\x47\x6f\x7d\xa6\xf0\x55\x2d\xc7\xc2\x9e\x15\x1d\x8a\xc6\xb5\x5d\x71\xe7\x57\x67\x17\xea\x3e\x28\xe8\x3c\x1a\x95\x67\xfa\x01\x4d\x91\x10\xc9\x05\x7f\xf2\xce\x85\x5f\x50\x84\xd6\x45\xb0\xda\x84\x3c\xd6\xe4\x44\xba\xed\xd1\x04\x01\x46\x9f\x4f\xb1\x15\x96\xeb\xbb\x21\x8a\x57\x5e\x55\x45\x59\xd6\x26\xe6\x8a\xc8\x33\xab\xb3\xdb\xd5\x20\xca\x26\xeb\x8b\xa8\x51\x91\x20\x8b\x05\xf9\x56\x69\x18\x5c\x2f\xb1\xb9\xe6\x6a\x74\x07\x6e\x4a\x1a\x1f\x07\x73\x98\xab\x67\xa9\xb4\x56\xc1\x21\x2d\x3a\x0d\x1f\xbc\xbf\xb6\x72\xa1\x92\xa9\x57\x12\xe3\x52\x4f\xe3\x2b\xa2\xc2\xf7\x21\xe4\x8c\x7e\x4e\x9c\xcb\xe1\xa6\x13\xb7\x9f\xaf\x4c\xe7\xdb\xb0\x66\xe5\x56\x65\x39\x6f\x08\x93\xa8\x6d\x9a\xb0\x7b\xdd\xbc\x95\xbc\x7f\xa8\x1e\x0e\x54\xcc\x59\xc4\xed\x29\xca\x54\x46\xcd\xfb\x39\xec\xbe\x4c\x4a\x7e\xb1\x74\x29\xe3\xd1\xe3\xc4\xfc\x81\x92\x8f\x89\x3f\x77\xd9\xa6\x59\x36\xbc\x2c\x5d\xf0\x56\xa5\xdb\xb4\x47\x16\x9e\xee\xcd\x42\x1a\xed\x12\xa7\x44\xc0\xdd\xc8\xe9\x3e\xd4\xa7\x05\x96\x45\xec\x5f\x79\x46\xf1\xa3\xd7\x32\xb9\x11\xa3\x3a\x68\x99\xf5\xbe\xe2\x1c\xfe\x72\x2a\x2e\x97\x66\xdc\xa0\x39\xbd\xb8\xc1\x60\xa7\xe5\x52\x96\x34\xe3\xe1\x2b\xc1\xc2\x91\x34\xb9\x62\xa7\x7d\x1b\x6c\x98\xc6\xbb\x2b\x2c\xda\x13\xda\xbd\xe3\x8f\x2e\x72\x9e\x92\x1f\x3f\x5d\xc3\xa9\x22\xf8\x39\x6f\x8c\x67\x1f\xbc\xb8\x9e\xe6\x98\x26\xa3\xc1\x0f\xd5\x67\x46\x5c\xb2\xd7\x78\xcb\xe9\x4b\xbf\x28\x7c\x17\x5d\x83\x1d\xee\xb1\xb5\x46\x6c\xd2\xae\xc6\x31\x77\x86\x61\x8e\xcf\x80\x17\xf3\xb4\x64\xa4\x17\x78\xc9\xc1\x4b\xb1\x75\x54\x31\x54\xfa\xc8\xdb\x7a\xa1\x09\xe9\x66\x8a\x47\x1a\x94\xb8\x6e\xf9\xa5\x4f\xe3\x9b\xce\x29\xa6\xa8\x04\x96\x48\x7f\x5d\xdf\x35\x79\xb0\x23\x26\xa6\xd6\x50\xb0\x01\x38\x24\xe1\x6d\xca\xca\xea\x2f\x3a\x9f\x57\x4a\x1d\xf0\x2e\xe1\x55\x1a\x74\x56\x7b\xce\xec\xd8\x24\x9b\xb7\x6e\xb7\xa9\x69\xb3\xc3\xb7\x79\x4d\xd2\xc3\x1e\xb7\x10\x3a\x67\x5e\x64\x1a\x1e\xa6\x02\x29\xde\x3f\x58\xc4\x8d\xde\x1f\xb5\x1c\x9b\xc2\x5a\xbc\x3e\x44\x55\x03\x2c\x24\x41\xee\x88\x86\xf8\x65\x24\x63\xe7\xea\x59\x79\x65\x66\x8a\xe7\x46\xd2\x6e\x51\xad\xcd\x79\xab\xb6\x8c\xc7\xe2\x1b\x28\x2d\x08\x80\x52\x47\x35\xcc\x2d\xc0\xfd\x50\x5c\x10\xd2\x86\xe0\x69\x73\x10\xd8\x34\xb9\x97\x40\xf1\x03\x89\x7a\xf7\xd1\x1f\xba\xe5\x08\x4b\x24\x43\x27\x97\x80\x13\x6b\xc6\x16\x43\xe8\xf8\x2a\xd8\xb4\x4d\xfa\x19\x71\x2e\xe5\x11\xc0\x51\x4f\x10\x68\x1b\xf8\xef\x65\x87\x43\x72\xe9\xf8\x13\xcf\x3a\x1c\xe9\x65\x7e\xb1\x60\x83\x3f\xae\xbb\xcb\x3e\x6e\x50\xdb\xfd\xeb\xff\x42\x46\xfc\x33\x69\x7f\x28\x83\x40\xf1\xea\xcb\x03\x2c\x28\xe9\x06\x8b\x8a\xef\xa7\x2e\x8c\xec\xba\x54\xe4\xfb\xdb\xef\x3e\x91\x3d\x3c\x30\x78\x1f\xdf\x62\x1d\x97\xec\xdb\x08\xd2\xa8\xf0\x02\xb1\x2b\xbd\x28\x53\x9a\xa9\x47\xe0\xc1\xca\x60\x7c\x01\x27\x87\x01\xec\x84\xf2\xc3\x34\xc6\x7c\x4a\x94\x1c\x14\xa6\x50\x89\x3d\xfd\x9c\x9e\xe9\xcb\x3e\xe3\xf3\x3e\x68\x8e\x50\xcc\x6b\x74\x13\x00\x71\x53\x44\x9d\xe3\x6c\xc7\x98\x35\xe3\x3e\xd1\x93\x66\xad\xb7\xa8\xaa\xb8\x2f\x5d\x42\xa3\x4e\xa9\x8c\x52\xd8\xd8\x71\x99\x25\x7d\xad\x8f\x31\x3e\x87\xb6\x55\x90\x7e\xb2\x3c\xd3\xe8\x83\xf3\xfb\x8b\xf0\xae\xe7\x22\x8d\x95\x43\x1e\x8d\x2a\x96\x6f\x39\xf5\x29\xbb\xcd\xbb\x08\x64\x9a\xdb\xc7\xd1\xc6\x23\x78\xa4\x01\x9b\x6c\x53\xd2\x1f\xa0\x3f\x25\x0e\xf9\x3a\x2f\x5a\x25\x79\x87\xd9\xdd\xfc\x74\x20\x9e\x3c\x50\xda\xb0\x4f\xa8\x39\x1c\xd8\x41\xcd\xef\xe7\x5d\x97\xf6\xfc\xc9\xa8\x4e\x12\x37\xcb\x74\x98\x36\x59\xd7\xf9\xd8\x7d\xd6\xd3\xd0\x05\xd4\xca\xb1\x64\x16\xcc\x60\xbf\x3a\xb7\xc4\x4f\x3e\x44\x9b\x06\x19\x1e\x4c\xe1\x39\x92\xab\xdd\x47\x2a\x8f\xae\xe1\x73\xca\x1f\x57\xf1\xc1\xee\xf5\x2e\xd4\xb4\xe1\x22\x99\x3d\x45\xe3\xd1\xeb\xe9\x64\x9b\x43\x70\xc6\x0c\x22\x57\xc7\x43\xe7\x9d\x4f\x08\x8b\xc7\xf3\x13\xf4\x77\xe6\x44\x8d\xbb\x49\xab\x77\xf7\x42\xbb\xd1\x78\x11\xcd\xdc\x3c\x5d\x80\x4f\x56\xa3\x10\x0d\x27\x6f\x8c\xb6\x88\x8d\x44\xff\x71\xdc\xa9\xe5\xd1\x47\xc8\x25\x03\xe9\x6c\x72\x04\x1f\x8e\x07\x98\xbd\xac\x1c\x42\x13\x8e\xe5\xf2\x04\x9e\xca\x23\xce\xbb\x7c\xe3\xfb\xce\xaa\x4c\x28\x41\xc7\xfd\x5e\xa2\xd7\x7b\xff\xb8\x6e\xfb\x4c\x56\xea\x62\x22\xb6\x82\x76\x93\xd0\x51\x54\x10\x0b\xfe\x08\x24\x4f\xba\x9b\x41\xfb\xcf\xc5\x68\xb9\xd4\x9b\x39\xc1\x6d\x58\xd9\x09\x1e\x56\x08\xce\xdf\xba\x52\xd4\xac\x9b\x05\xef\x40\x2f\x97\x5d\xa9\xfd\xd4\xe1\x87\x16\x30\x24\x68\x9c\xcf\x4a\x4c\x4f\x49\xc6\x0a\xec\x25\x8d\x4f\xad\xec\xb9\xd3\xfa\xec\xc3\x62\xfd\x92\xd8\x62\xbb\x56\x75\xcd\x61\x7c\xc1\x60\x10\xcf\xfc\xe3\xa0\xec\x52\x00\x76\x84\x1e\x13\x4d\x6f\xe6\x97\x9e\x19\x39\x25\xe6\x2b\x5d\xdc\x65\x41\x65\x93\x98\xdc\x14\x69\x35\x92\xba\x93\x06\xce\x60\xdc\x11\x3b\x23\x73\x32\x10\xf5\xbc\x55\x0a\x1d\x88\x9e\x32\xcb\xd0\xa9\x4f\x64\x85\x65\x91\x27\x7e\xd7\x0b\x84\xa5\x7f\xaf\x1c\xff\x09\x00\x00\xff\xff\x9a\xd6\x77\x81\x8f\x08\x00\x00")

func insecure_pem() ([]byte, error) {
	return bindata_read(
		_insecure_pem,
		"insecure.pem",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"insecure-key.pem": insecure_key_pem,
	"insecure.csr":     insecure_csr,
	"insecure.pem":     insecure_pem,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"insecure-key.pem": &_bintree_t{insecure_key_pem, map[string]*_bintree_t{}},
	"insecure.csr":     &_bintree_t{insecure_csr, map[string]*_bintree_t{}},
	"insecure.pem":     &_bintree_t{insecure_pem, map[string]*_bintree_t{}},
}}
