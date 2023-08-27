package providers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/robertkrimen/otto"
)

type deepAI struct {
	provider
}

var deepAIprovider = deepAI{
	provider: provider{
		url:       "https://api.deepai.org/make_me_a_pizza",
		active:    true,
		canStream: true,
	},
}

const (
	deepAIJSCode = `var agent = 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36'
	var a, b, c, d, e, h, f, l, g, k, m, n, r, x, C, E, N, F, T, O, P, w, D, G, Q, R, W, I, aa, fa, na, oa, ha, ba, X, ia, ja, ka, J, la, K, L, ca, S, U, M, ma, B, da, V, Y;
	h = Math.round(1E11 * Math.random()) + "";
	f = function () {
		for (var p = [], q = 0; 64 > q;) p[q] = 0 | 4294967296 * Math.sin(++q % Math.PI);
		
		return function (t) {
			var v, y, H, ea = [v = 1732584193, y = 4023233417, ~v, ~y],
				Z = [],
				A = unescape(encodeURI(t)) + "\u0080",
				z = A.length;
			t = --z / 4 + 2 | 15;
			for (Z[--t] = 8 * z; ~z;) Z[z >> 2] |= A.charCodeAt(z) << 8 * z--;
			for (q = A = 0; q < t; q += 16) {
				for (z = ea; 64 > A; z = [H = z[3], v + ((H = z[0] + [v & y | ~v & H, H & v | ~H & y, v ^ y ^ H, y ^ (v | ~H)][z = A >> 4] + p[A] + ~~Z[q | [A, 5 * A + 1, 3 * A + 5, 7 * A][z] & 15]) << (z = [7, 12, 17, 22, 5, 9, 14, 20, 4, 11, 16, 23, 6, 10, 15, 21][4 * z + A++ % 4]) | H >>> -z), v, y]) v = z[1] | 0, y = z[2];
				for (A = 4; A;) ea[--A] += z[A]
			}
			for (t = ""; 32 > A;) t += (ea[A >> 3] >> 4 * (1 ^ A++) & 15).toString(16);
			return t.split("").reverse().join("")
		}
	}();
	
	"tryit-" + h + "-" + f(agent + f(agent + f(agent + h + "x")));`
)

func (d *deepAI) NewCompletion(messages []Message) (*string, error) {
	stream, err := d.NewCompletionStream(messages)
	if err != nil {
		return nil, err
	}
	var res string
	for {
		str, ok := <-stream
		if !ok {
			break
		}
		res += str
	}
	return &res, nil
}

func (d *deepAI) NewCompletionStream(messages []Message) (chan string, error) {

	vm := otto.New()
	apiKey, err := vm.Run(deepAIJSCode)
	if err != nil {
		return nil, err
	}
	headers := map[string]string{
		"api-key":      apiKey.String(),
		"User-Agent":   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36",
		"Content-Type": "application/x-www-form-urlencoded",
	}
	payload := url.Values{}
	payload.Add("chas_style", "chat")
	messageString, err := json.Marshal(messages)
	if err != nil {
		return nil, err
	}
	payload.Add("chatHistory", string(messageString))
	req, err := http.NewRequest("POST", d.url, strings.NewReader(payload.Encode()))
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	responseChan := make(chan string)
	go func() {
		defer resp.Body.Close()
		defer close(responseChan)
		for {
			buf := make([]byte, 32)
			n, err := resp.Body.Read(buf)
			if n > 0 {
				responseChan <- string(buf[:n])
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println("Error:", err)
				break
			}
		}
	}()
	return responseChan, nil
}
