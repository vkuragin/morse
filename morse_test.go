package morse

import "testing"

var m *Morse

func init() {
	var err error
	m, err = New()
	if err != nil {
		panic(err)
	}
}

func TestEncode(t *testing.T) {
	tests := []struct {
		str  string
		want string
	}{
		{"SOS", "... --- ..."},
		{"huY piZda", ".... ..- -.--   .--. .. --.. -.. .-"},
		{"a", ".-"},
		{"A", ".-"},
		{" ", " "},
		{"1.7f", ".---- .-.-.- --... ..-."},
		{"wtf?", ".-- - ..-. ..--.."},
		{"1/7", ".---- -..-. --..."},
	}
	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			if got := m.Encode(tt.str); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		str  string
		want string
	}{
		{"... --- ...", "SOS"},
		{".... ..- -.--  .--. .. --.. -.. .-", "HUY PIZDA"},
		{".-", "A"},
		{" ", " "},
		{".---- .-.-.- --... ..-.", "1.7F"},
		{".-- - ..-. ..--..", "WTF?"},
		{".---- -..-. --...", "1/7"},
	}
	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			if got := m.Decode(tt.str); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
