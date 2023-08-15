// @author:llychao<lychao_vip@163.com>

// @contributor: Junyi<me@junyi.pw>

// @date:2020-02-18

// @功能:golang m3u8 video Downloader

package main

import (
	"testing"
)

func TestDrawProgressBar(t *testing.T) {
	type args struct {
		prefix     string
		proportion float32
		width      int
		suffix     []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "0",
			args: args{
				prefix:     "1",
				proportion: float32(0.1),
				width:      20,
				suffix:     []string{"123.ts"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DrawProgressBar(tt.args.prefix, tt.args.proportion, tt.args.width, tt.args.suffix...)
		})
	}
}
