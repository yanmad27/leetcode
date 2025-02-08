package main

func main() {
	obj := Constructor()
	obj.Change(1, 10)
	obj.Change(1, 10)
	obj.Find(10)
	obj.Change(1, 20)
	obj.Find(10)
}
