package compiler

import (
	"bytes"
	"os"
)

// Oh boi. Well, that escalated quickly.
// This map is basically Script -> Message -> field -> order
type fieldNumberMapping map[string]map[string]map[string]int

func getAssignedNumber(mapping fieldNumberMapping, script, message, field string) *int {
	if scriptMap, ok := mapping[script]; ok {
		if messageMap, ok := scriptMap[message]; ok {
			if number, ok := messageMap[field]; ok {
				return &number
			}
		}
	}
	return nil
}

func setAssignedNumber(mapping fieldNumberMapping, script, message, field string, number int) {
	if mapping[script] == nil {
		mapping[script] = map[string]map[string]int{}
	}
	if mapping[script][message] == nil {
		mapping[script][message] = map[string]int{}
	}
	mapping[script][message][field] = number
}

func deleteAssignedNumber(mapping fieldNumberMapping, script, message, field string) {
	if scriptMap, ok := mapping[script]; ok {
		if messageMap, ok := scriptMap[message]; ok {
			delete(messageMap, field)
			if len(messageMap) == 0 {
				delete(scriptMap, message)
				if len(scriptMap) == 0 {
					delete(mapping, script)
				}
			}
		}
	}
}

func saveFieldNumberMapping(mapping fieldNumberMapping, filePath string) error {
	buf := new(bytes.Buffer)

	for scriptName, script := range mapping {
		writeStrToBuf(buf, scriptName)
		writeIntToBuf(buf, len(script))

		for messageName, message := range script {
			writeStrToBuf(buf, messageName)
			writeIntToBuf(buf, len(message))

			for fieldName, field := range message {
				writeStrToBuf(buf, fieldName)
				writeIntToBuf(buf, field)
			}
		}
	}

	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		if err := os.Remove(filePath); err != nil {
			return err
		}
	}

	fs, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer fs.Close()

	_, err = fs.Write(buf.Bytes())
	return err
}

func readFieldNumberMapping(filePath string) (fieldNumberMapping, error) {
	mapping := fieldNumberMapping{}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return mapping, nil
	}

	fs, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer fs.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(fs)
	if err != nil {
		return nil, err
	}

	for buf.Len() > 0 {
		scriptName := readStrFromBuf(buf)
		scriptLen := readIntFromBuf(buf)

		script := map[string]map[string]int{}
		for i := 0; i < scriptLen; i++ {
			messageName := readStrFromBuf(buf)
			messageLen := readIntFromBuf(buf)

			message := map[string]int{}
			for j := 0; j < messageLen; j++ {
				fieldName := readStrFromBuf(buf)
				field := readIntFromBuf(buf)

				message[fieldName] = field
			}

			script[messageName] = message
		}

		mapping[scriptName] = script
	}

	return mapping, nil
}

func readStrFromBuf(buf *bytes.Buffer) string {
	strLen := readIntFromBuf(buf)
	return string(buf.Next(strLen))
}

func readIntFromBuf(buf *bytes.Buffer) int {
	b := buf.Next(4)
	return int(b[0]) | int(b[1])<<8 | int(b[2])<<16 | int(b[3])<<24
}

func writeStrToBuf(buf *bytes.Buffer, str string) {
	writeIntToBuf(buf, len(str))
	buf.WriteString(str)
}

func writeIntToBuf(buf *bytes.Buffer, i int) {
	buf.WriteByte(byte(i))
	buf.WriteByte(byte(i >> 8))
	buf.WriteByte(byte(i >> 16))
	buf.WriteByte(byte(i >> 24))
}
