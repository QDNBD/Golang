package main

type Phone interface {
	call()
}

type NokiaPhone struct {

}

func (nokiaPhone NokiaPhone) call() {
	println("I am NokiaPhone")
}

type IPhone struct {

}

func (iphone IPhone) call()  {
	println("I am IPhone")
}




func main() {
	var phone Phone

	phone = new (NokiaPhone)
	phone.call()

	phone = new (IPhone)
	phone.call()
}