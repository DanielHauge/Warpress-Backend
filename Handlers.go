package main

import (
	"bytes"
	"encoding/json"
	"gopkg.in/russross/blackfriday.v2"
	"net/http"
	"reflect"
)


func SetupIndexPage()[]byte{
	var buffer bytes.Buffer

	for _, v := range routes {
		buffer.WriteString("\n\n### "+v.Name+"\n")
		buffer.WriteString("##### Route: "+v.Pattern+"\n")
		buffer.WriteString("##### Method: "+v.Method+"\n")


		if v.ExpectedInput != nil {
			buffer.WriteString("##### Input: \n")

			inputFields := reflect.Indirect(reflect.ValueOf(v.ExpectedInput))
			numOfInputFields := inputFields.Type().NumField()

			for i := 0; i < numOfInputFields; i++ {
				buffer.WriteString("- ")
				buffer.WriteString(inputFields.Type().Field(i).Name)
				buffer.WriteString(" : ")
				buffer.WriteString(inputFields.Type().Field(i).Type.Kind().String() + "\n\n")
			}
		}

		if v.ExpectedOutput != nil {
			buffer.WriteString("##### Output: \n")
			outputFields := reflect.Indirect(reflect.ValueOf(v.ExpectedOutput))
			numOfOutputFields := outputFields.Type().NumField()

			for i := 0; i < numOfOutputFields; i++ {
				buffer.WriteString("- ")
				buffer.WriteString(outputFields.Type().Field(i).Name)
				buffer.WriteString(" : ")
				buffer.WriteString(outputFields.Type().Field(i).Type.Kind().String() + "\n")
			}
		}

		buffer.WriteString("\n##### Example:\n")
		buffer.WriteString("- Input:\n")
		var b []byte
		if v.ExpectedInput != nil{
			b, _ = json.Marshal(v.ExpectedInput)
		} else {
			b = []byte("Nothing")
		}
		buffer.WriteString(string(b)+"\n")

		buffer.WriteString("\n- Output:\n")
		if v.ExpectedOutput != nil {
			b, _ = json.Marshal(v.ExpectedOutput)
		} else {
			b = []byte("Nothing")
		}
		buffer.WriteString(string(b)+"\n")
	}
	return buffer.Bytes()
}

var IndexPage []byte

func Index(w http.ResponseWriter, r *http.Request) {
	var buffer bytes.Buffer
	buffer.WriteString("# Warpress API\n")
	buffer.WriteString("This is a api for the website of Warpress, this page is only available during development\n\n")
	buffer.WriteString("## Api endpoints:\n\n")
	buffer.Write(IndexPage)



	output := blackfriday.Run([]byte(buffer.Bytes()))
	w.Write(output)
}



