package persist

import "log"

func ItemSave() chan interface{} {
	out := make(chan interface{})
	itemCount:=0
	go func() {
		for {
			item := <- out
			itemCount++
			log.Printf("item save got a item %d -- %v \n",itemCount,item)
		}
	}()
	return out
}
