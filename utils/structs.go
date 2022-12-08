package utils

type TestStruct struct {
	name string
}

func alterName(data *TestStruct, name string) {
	data.name = name;
}