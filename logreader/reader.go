package logreader

import (
	"eNBlogReader/core"
	"log"
	"strings"

	"github.com/hpcloud/tail"
)

func ReadLog(logFilePath string) {
	log.Printf("Invoke project 'eNB Log Reader' in version 1.0 invoked for log file: %v", logFilePath)

	t, err := tail.TailFile(logFilePath, tail.Config{Follow: true})
	log.Printf("Start to read file.")

	if err != nil {
		log.Printf("ERROR: %v", err)
	}

	for line := range t.Lines {
		// fmt.Println(line.Text)
		newLine := line.Text
		if !strings.Contains(newLine, "----DL") &&
			!strings.Contains(newLine, "UE_ID") &&
			!strings.Contains(newLine, "PRACH") &&
			!strings.Contains(newLine, "disconnected") &&
			!strings.Contains(newLine, "idle") {

			newLine = strings.Join(strings.Fields(newLine), " ")
			// log.Printf("line (trimmed) %s", newLine)
			s := strings.Split(newLine, " ")

			if len(s) > 15 {
				ueid := s[0]
				dlbr := numberConverter(s[9])
				ulbr := numberConverter(s[15])
				snr := s[12]
				// log.Printf("ue_id %v, dl-br %v, ul-br %v, snr %v", ueid, dlbr, ulbr, snr)

				newENBLog := core.Enb{UEid: ueid, DLbr: dlbr, ULbr: ulbr, SNR: snr}
				core.PushENBStats(newENBLog)
			} else {
				log.Printf("ERROR: Not able to parse this line: %s", newLine)
			}

		}
	}
	log.Printf("fin.")
}

func numberConverter(input string) string {
	result := input
	if strings.Contains(input, ".") {
		pointPosition := strings.Index(result, ".")
		if strings.Index(result, "k") > -1 {
			kPosition := strings.Index(result, "k")

			if kPosition-pointPosition == 2 {
				result = strings.Replace(result, "k", "00", -1)
			} else if kPosition-pointPosition == 3 {
				result = strings.Replace(result, "k", "0", -1)
			}
			result = strings.Replace(result, ".", "", -1)
		} else if strings.Index(result, "M") > -1 {
			mPosition := strings.Index(result, "M")
			//			log.Print("\tmPosition-pointPosition ", mPosition-pointPosition)
			if mPosition-pointPosition == 2 {
				result = strings.Replace(result, "M", "00000", -1)
			} else if mPosition-pointPosition == 3 {
				result = strings.Replace(result, "M", "0000", -1)
			}
			result = strings.Replace(result, ".", "", -1)
		}
	} else {
		result = strings.Replace(result, "k", "000", -1)
		result = strings.Replace(result, "M", "000000", -1)
	}

	return result
}
