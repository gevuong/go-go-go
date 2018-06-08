In another file in the channels package, we define a `writeToChannel` function:

```
func writeToChannel(channel chan string) {
channel <- "hello"
}
```

Update the `readFromChannel` function to do the following:

Create a new channel that takes string values
Call writeToChannel as a goroutine, and pass your new channel as an argument
Read a string from the channel and return that string

```
package channels

func readFromChannel() string {
// CREATE A CHANNEL FOR string VALUES HERE
    channel := make(chan string)

// HERE, CALL writeToChannel AS A GOROUTINE, AND PASS IT YOUR CHANNEL
    go writeToChannel(channel)

// HERE, READ A STRING FROM YOUR CHANNEL AND RETURN IT
    return <-channel
}
```