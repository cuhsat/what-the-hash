package wth

import (
	"fmt"
	"strings"
	"testing"
)

func Example() {
	ch := make(chan string, len(DB))

	go Search([]byte("1234567890abcd"), ch)

	for m := range ch {
		fmt.Println(m)
	}

	// Output:
	// Cisco Type 7
	// BigCrypt
}

// Source: https://hashcat.net/wiki/doku.php?id=example_hashes
func TestSearch(t *testing.T) {
	for _, tt := range []struct {
		name, hash string
	}{
		{
			name: "CRC-32",
			hash: "c762de4a",
		},
		{
			name: "MD5",
			hash: "8743b52063cd84097a65d1633f5c74f5",
		},
		{
			name: "SHA-1",
			hash: "b89eaac7e61417341b710b727768294d0e6a277b",
		},
		{
			name: "SHA-256",
			hash: "127e6fbfe24a750e72930c220a8e138275656b8e5d8f48a98c3c92df2caba935",
		},
		{
			name: "SHA-512",
			hash: "82a9dda829eb7f8ffe9fbe49e45d47d2dad9664fbb7adf72492e3c81ebd3e29134d9bc12212bf83c6840f10e8246b9db54a4859b7ccd0123d86e5872c1e5082f",
		},
		{
			name: "RIPEMD-160",
			hash: "012cb9b334ec1aeb71a9c8ce85586082467f7eb6",
		},
		{
			name: "GRUB 2",
			hash: "grub.pbkdf2.sha512.10000.7d391ef48645f626b427b1fae06a7219b5b54f4f02b2621f86b5e36e83ae492bd1db60871e45bc07925cecb46ff8ba3db31c723c0c6acbd4f06f60c5b246ecbf.26d59c52b50df90d043f070bd9cbcd92a74424da42b3666fdeb08f1a54b8f1d2f4f56cf436f9382419c26798dc2c209a86003982b1e5a9fcef905f4dfaa4c524",
		},
		{
			name: "DNSSEC",
			hash: "7b5n74kq8r441blc2c5qbbat19baj79r:.lvdsiqfj.net:33164473:1",
		},
		{
			name: "SAM",
			hash: "aad3b435b51404eeaad3b435b51404ee:b4b9b02e6f09a9bd760f388b67351e2b",
		},
		{
			name: "NTLM",
			hash: "b4b9b02e6f09a9bd760f388b67351e2b",
		},
		{
			name: "NetNTLMv1",
			hash: "u4-netntlm::kNS:338d08f8e26de93300000000000000000000000000000000:9526fb8c23a90751cdd619b6cea564742e1e4bf33006ba41:cb8086049ec4736c",
		},
		{
			name: "NetNTLMv2",
			hash: "admin::N46iSNekpT:08ca45b7d7ea58ee:88dcbe4446168966a153a0064958dac6:5c7830315c7830310000000000000b45c67103d07d7b95acd12ffa11230e0000000052920b85f78d013c31cdb3b92f5d765c783030",
		},
	} {
		t.Run("Test Search with "+tt.name, func(t *testing.T) {
			ch := make(chan string, len(DB))

			Search([]byte(tt.hash), ch)

			for a := range ch {
				if strings.Contains(a, tt.name) {
					return
				}
			}

			t.Fatal("nothing found")
		})
	}
}

func BenchmarkSearch(b *testing.B) {
	b.Run("Benchmark Search", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Search([]byte(""), make(chan string, len(DB)))
		}
	})
}
