package main

import "fmt"
import "net"

type DNSHeader struct {
	ID int8 // random dns id
	QR bool // question == 1, answer == 1
	Opcode [4]bool // type of question
	        // 0 == a standard query (QUERY)
		// 1 == an inverse query (IQUERY)
		// 2 == a server status request (STATUS)
		// 3-15 == reserved for future use
	AA bool // privilege. authority answer == 1, other ==0
	TC bool // TrunCation notify bit
	RD bool // Recursion query bit
	RA bool // Recursion available bit
	Z bool  // Reserved for future use.  Must be zero.
	RCODE [4]bool // Respon code.
	              // 0 == No error condition
		      // 1 == Format error
		      // 2 == Server failure
		      // 3 == Name Error(Name not found from authority)
		      // 4 == Not Implemented
		      // 5 == Refused
		      // 6-15 == Reserved for future use
	QDCOUNT int16
	ANCOUNT int16
	NSCOUNT int16
	ARCOUNT int16
}

type DNSQuestion struct {
}

type DNSAnswer struct {
}

type DNSAuthority struct {
}

type DNSAdditional struct {
}

type DNSPacket struct {
	Header DNSHeader
	Question DNSQuestion
	Answer []DNSAnswer
	Authority []DNSAuthority
	Additional []DNSAdditional
}

func main() {
	fmt.Println("Hello, onokatio full DNS resolver.")
	conn, _ := net.Dial("udp", "198.41.0.4:53") // 198.41.0.4 == a.root-servers.net
	defer conn.Close()
	var query = []byte{0xa9, 0x4e, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x03, 0x77, 0x77, 0x77, 0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x03, 0x63, 0x6f, 0x6d, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x29, 0x05, 0xc8, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00}


	conn.Write(query)

	response := make([]byte, 2000)
	conn.Read(response)
	fmt.Println(response)
}
