package logs

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func ReadLog() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		time.Sleep(2000)
		fmt.Println("run")
	}

}
