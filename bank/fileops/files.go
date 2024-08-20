package fileops
 
import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, error := os.ReadFile(fileName)
	if error != nil {
		return 0, errors.New("the file was not found")
	}
	var textValue = string(data)
	bal, error := strconv.ParseFloat(textValue, 64)
	if error != nil {
		return 0, errors.New("failed to parse string to a float")
	}
	return bal, nil
}
