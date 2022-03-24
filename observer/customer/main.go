package main

func main() {

	shirtItem := NewItem("Nike Shirt")

	observerFirst := &Customer{Id: "abc@gmail.com"}
	observerSecond := &Customer{Id: "xyz@gmail.com"}

	shirtItem.Register(observerFirst)
	shirtItem.Register(observerSecond)

	shirtItem.UpdateAvailability()
}
