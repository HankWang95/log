# log
A Golang log Third-party libraries .


![image](http://github.com/HankWang96/readme_add_pic/raw/master/images/example.jpg)

---

## How To Use It

### First 
> go get -u github.com/HankWang95/log

```go
	import github.com/HankWang95/log
```

### Second
```go
	log.NewLogger("your project's name", "/Where/Would/You/Like/To/Save/Your/LogFile")
```

### Third
You can use three function to Log. 

```go
	log.Error("test1", "are", "you", "ok?")
	log.Debug("test2", "are", "you", "ok?")
	log.Info("test3", "are", "you", "ok?")
```