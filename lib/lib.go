package lib

import(
	"log"
	"github.com/spf13/viper"
)

func Check(e error){
	if e != nil {
		log.Print(e)
	}
}
