package instruments

import (
	"bufio"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

const csvURL = "https://public.fyers.in/sym_details/NSE_FO.csv"

type Instrument struct {
	Name       string
	Symbol     string
	Underlying string
	Expiry     string
	OptionType string
	Strike     float64
	Instrument string
}

var (
	cache = map[string]Instrument{}
	mu    sync.RWMutex
)

func Load() error {

	resp, err := http.Get(csvURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	tmp := make(map[string]Instrument)

	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {

		line := scanner.Text()
		fields := strings.Split(line, ",")

		if len(fields) < 18 {
			continue
		}

		symbol := strings.TrimSpace(fields[9])
		name := strings.TrimSpace(fields[1])
		underlying := strings.TrimSpace(fields[11])
		expiry := strings.TrimSpace(fields[7])

		strike, _ := strconv.ParseFloat(fields[13], 64)
		optionType := strings.TrimSpace(fields[14])

		instType := "INDEX"

		if strings.HasSuffix(symbol, "FUT") {
			instType = "FUTURE"
		}

		if optionType == "CE" || optionType == "PE" {
			instType = "OPTION"
		}

		tmp[symbol] = Instrument{
			Name:       name,
			Symbol:     symbol,
			Underlying: underlying,
			Expiry:     expiry,
			Strike:     strike,
			OptionType: optionType,
			Instrument: instType,
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	mu.Lock()
	cache = tmp
	mu.Unlock()

	return nil
}

func Get(symbol string) (Instrument, bool) {

	mu.RLock()
	defer mu.RUnlock()

	inst, ok := cache[symbol]
	return inst, ok
}

func All() []Instrument {

	mu.RLock()
	defer mu.RUnlock()

	out := make([]Instrument, 0, len(cache))

	for _, inst := range cache {
		out = append(out, inst)
	}

	return out
}

func Search(prefix string) []Instrument {

	prefix = strings.ToUpper(prefix)

	mu.RLock()
	defer mu.RUnlock()

	out := make([]Instrument, 0)

	for _, inst := range cache {

		if strings.Contains(strings.ToUpper(inst.Symbol), prefix) ||
			strings.Contains(strings.ToUpper(inst.Name), prefix) ||
			strings.Contains(strings.ToUpper(inst.Underlying), prefix) {

			out = append(out, inst)

			if len(out) >= 100 {
				break
			}
		}
	}

	return out
}
