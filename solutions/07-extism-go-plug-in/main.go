package main

import (
	"strconv"

	"github.com/extism/go-pdk"
)

// go

//export hello
func hello() int32 {

	// read the function argument from the memory
	input := pdk.Input()

	// display input
	pdk.Log(pdk.LogInfo, string(input))

	// read config
	// {"url":"https://jsonplaceholder.typicode.com/todos/1"}
	url, _ := pdk.GetConfig("url")
	// display input
	pdk.Log(pdk.LogInfo, "🌍" + url)
	
	// use a host function to make a request
	req := pdk.NewHTTPRequest("GET", url)
	res := req.Send()

	pdk.Log(pdk.LogInfo, "📝 status: " + strconv.FormatUint(uint64(res.Status()), 10) )
	
	pdk.Log(pdk.LogInfo, "📝 body: "+string(res.Body()))
	
	// create output
	//output := "🎉 Extism is 💜, 🌍: " + url
	output := "🎉 Extism is 💜"
	
	// return the value
	// copy output to host memory
	mem := pdk.AllocateString(output)
	pdk.OutputMemory(mem)

	return 0
	
}

func main() {}