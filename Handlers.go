package main

import (
	"fmt"
	"net/http"
	"gopkg.in/russross/blackfriday.v2"
	"bytes"
	"reflect"
	"encoding/json"
)


func SetupIndexPage()[]byte{
	var buffer bytes.Buffer
	for _, v := range routes {
		buffer.WriteString("### "+v.Name+"\n")
		buffer.WriteString("##### Route: "+v.Pattern+"\n")
		buffer.WriteString("##### Method: "+v.Method+"\n")
		buffer.WriteString("##### Input: \n")
		inputFields := reflect.Indirect(reflect.ValueOf(v.ExpectedInput))
		numOfInputFields := inputFields.Type().NumField()

		for i := 0; i<numOfInputFields ;i++  {
			buffer.WriteString("- ")
			buffer.WriteString(inputFields.Type().Field(i).Name)
			buffer.WriteString(" : ")
			buffer.WriteString(inputFields.Type().Field(i).Type.Kind().String()+"\n\n")
		}

		buffer.WriteString("##### Output: \n")
		outputFields := reflect.Indirect(reflect.ValueOf(v.ExpectedOutput))
		numOfOutputFields := outputFields.Type().NumField()

		for i := 0; i<numOfOutputFields ;i++  {
			buffer.WriteString("- ")
			buffer.WriteString(outputFields.Type().Field(i).Name)
			buffer.WriteString(" : ")
			buffer.WriteString(outputFields.Type().Field(i).Type.Kind().String()+"\n")
		}


		buffer.WriteString("\n##### Example:\n")
		buffer.WriteString("- Input:\n")
		b, err := json.Marshal(v.ExpectedInput)
		buffer.WriteString(string(b)+"\n")
		buffer.WriteString("\n- Output:\n")
		b, err = json.Marshal(v.ExpectedOutput)
		buffer.WriteString(string(b)+"\n")
		if err != nil {	fmt.Printf("Error: %s", err)	}
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

