{#/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */#}

{{- self.add_import("math") }}

type {{ ffi_converter_name }} struct{}

var {{ ffi_converter_name }}INSTANCE = {{ ffi_converter_name }}{}

func ({{ ffi_converter_name }}) lift(cRustBuf C.RustBuffer) string {
	reader := fromCRustBuffer(cRustBuf).asReader()
	b, err := io.ReadAll(reader)
	if err != nil {
		panic(fmt.Errorf("reading reader: %w", err))
	}
	return string(b)
}

func ({{ ffi_converter_name }}) read(reader io.Reader) string {
	length := readInt32(reader)
	buffer := make([]byte, length)
	read_length, err := reader.Read(buffer)
	if err != nil {
		panic(err)
	}
	if read_length != int(length) {
		panic(fmt.Errorf("bad read length when reading string, expected %d, read %d", length, read_length))
	}
	return string(buffer)
}

func ({{ ffi_converter_name }}) lower(value string) C.RustBuffer {
	return stringToCRustBuffer(value)
}

func ({{ ffi_converter_name }}) write(writer io.Writer, value string) {
	if len(value) > math.MaxInt32 {
		panic("String is too large to fit into Int32")
	}

	writeInt32(writer, int32(len(value)))
	fmt.Printf("writing string of length %d\n", len(value))
	write_length, err := io.WriteString(writer, value)
	if err != nil {
		panic(err)
	}
	if write_length != len(value) {
		panic(fmt.Errorf("bad write length when writing string, expected %d, written %d", len(value), write_length))
	}
}

type {{ type_|ffi_destroyer_name }} struct {}

func ({{ type_|ffi_destroyer_name }}) destroy(_ {{ type_name }}) {}
