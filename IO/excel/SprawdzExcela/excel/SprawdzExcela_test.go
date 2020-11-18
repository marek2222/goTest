package excel

import (
	"log"
	"testing"
)

// func Test_ustawOs(t *testing.T) {
// 	type args struct {
// 		x int
// 		y int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
// 	}{
// 		{name: "success1", args: args{x: 1, y: 1}, want: "B"},
// 		{name: "success2", args: args{x: 0, y: 1}, want: "A"},
// 		{name: "success2", args: args{x: 3, y: 1}, want: "D"},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := ustawOs(tt.args.x, tt.args.y); got != tt.want {
// 				t.Errorf("ustawOs() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_czyPystyWiersz(t *testing.T) {
	type args struct {
		wiersz       []string
		liczbaKolumn int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "ok1", args: args{wiersz: []string{"", "", "", "", ""}, liczbaKolumn: 5}, want: true},
		{name: "false1", args: args{wiersz: []string{"", "", "asa", "", ""}, liczbaKolumn: 5}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := czyPystyWiersz(tt.args.wiersz, tt.args.liczbaKolumn)
			if got != tt.want {
				t.Errorf("czyPystyWiersz() = %v, want %v", got, tt.want)
			} else {
				log.Printf("name: %s, czyPystyWiersz() = %v, want: %v \n", tt.name, got, tt.want)
			}
		})
	}
}
