package logs

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLog() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
