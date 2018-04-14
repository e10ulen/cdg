package lib

import(
	"fmt"
	"github.com/spf13/viper"
)

func Check(e error){
	if e != nil {
		panic(e)
	}
}

func ReadConfig(){
	viper.SetConfigName(".qqw")
	viper.AddConfigPath("./")
	viper.AddConfigPath("$HOME/")
	viper.SetConfigType("yaml")
	fmt.Print("読み込んだぞ")
}
