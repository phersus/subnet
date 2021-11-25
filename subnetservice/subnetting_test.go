package subnetservice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrtSubNets_CrtSubNets(t *testing.T) {

	cases := []struct {
		name   string
		in     string
		expect []string
	}{
		{name: "should return all the subnets from /32 to /1 for a given IPv4 address",
			in: "172.16.10.37",
			expect: []string{
				"172.16.10.37/32",
				"172.16.10.36/31",
				"172.16.10.36/30",
				"172.16.10.32/29",
				"172.16.10.32/28",
				"172.16.10.32/27",
				"172.16.10.0/26",
				"172.16.10.0/25",
				"172.16.10.0/24",
				"172.16.10.0/23",
				"172.16.8.0/22",
				"172.16.8.0/21",
				"172.16.0.0/20",
				"172.16.0.0/19",
				"172.16.0.0/18",
				"172.16.0.0/17",
				"172.16.0.0/16",
				"172.16.0.0/15",
				"172.16.0.0/14",
				"172.16.0.0/13",
				"172.16.0.0/12",
				"172.0.0.0/11",
				"172.0.0.0/10",
				"172.0.0.0/9",
				"172.0.0.0/8",
				"172.0.0.0/7",
				"172.0.0.0/6",
				"168.0.0.0/5",
				"160.0.0.0/4",
				"160.0.0.0/3",
				"128.0.0.0/2",
				"128.0.0.0/1",
			},
		},
		{
			name: "ip address out of range, should return an empty slice []",
			in:   "172.302.1.9",
			expect: []string{
				"",
			},
		},
		{
			name: "should return an empty slice []",
			in:   "",
			expect: []string{
				"",
			},
		},
		{
			name: "bad formatted ip address, should return an empty slice []",
			in:   "172.15.20",
			expect: []string{
				"",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := CrtSubNets(tc.in)
			assert.Equal(t, tc.expect, got)
		})
	}
}
