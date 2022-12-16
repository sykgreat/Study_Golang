package main

func main() {
	var zs Driver
	zs = new(ZhangSan)

	var ls Driver
	ls = new(LiSi)

	var bmw Car
	bmw = new(BMW)

	var benz Car
	benz = new(Benz)

	zs.Drive(bmw)
	ls.Drive(benz)
}
