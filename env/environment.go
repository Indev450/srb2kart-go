package srb2k_env

import (
	"log"
	"os"
)

var HOSTADDR, ok = os.LookupEnv("SRB2K_LIB_HOST")

func ValidateEnvironment() {
    if !ok {
        log.Fatal("SRB2K_LIB_HOST env variable not set!")
    }
}
