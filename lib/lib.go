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

func ReadConfig(){
	viper.SetConfigName(".qqw")
	viper.AddConfigPath("./")
	viper.AddConfigPath("$HOME/")
	viper.SetConfigType("yaml")
	log.Print("読み込んだぞ")
}
