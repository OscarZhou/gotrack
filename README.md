# gotrack


### gotrack is a tool, developed by Golang, of tracking a function's runtime condition. It is not only able to log out at the beginning and ending of every function, but it is also allow to print logs in the progress of the function if it will take a long time.


### Inspiration comes from some of the situation encountered at work. I have to add `fmt.Println` in different functions. Too laborious. 


### How to use? 

`go get github.com/OscarZhou/gotrack`  



### Example

```
func main() {
	t := track.Default()
	Loop5(t)

}

// Loop5 starts a 5s loop
func Loop5(t *track.Track) {
	t.Start()
	defer t.End()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(1 * time.Second))
	}
}

```

### Result:  

>Start function: main.Loop5  
>InProgress function:    main.Loop5  
>InProgress function:    main.Loop5  
>InProgress function:    main.Loop5  
>InProgress function:    main.Loop5  
>InProgress function:    main.Loop5  
>End function:   main.Loop5 took 5.0123114s



### This project will keep updating and welcome to pull request.  
