package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

const ENV_FILE = ".env"

func main() {
	godotenv.Load(ENV_FILE)

	envVals := map[string]string{
		"PASSWORD": "",
		"IP":       "",
		"REACTIVE": "",
	}

	password := initPw()
	envVals["PASSWORD"] = password

	ip := initIp()
	envVals["IP"] = ip

	reactive := initReactive()
	envVals["REACTIVE"] = reactive

	godotenv.Write(envVals, ENV_FILE)
}

func initPw() string {
	password := os.Getenv("PASSWORD")
	if password != "" {
		var res string
		fmt.Print("Previous password detected. Would you like to create a new one? y/N ")
		_, err := fmt.Scanln(&res)
		if err != nil {
			return password
		}
		res = strings.TrimSpace(res)
		if slices.Contains([]string{"yes", "y"}, strings.ToLower(res)) {
			password = ""
		}

	}
	for password == "" {
		fmt.Print("Input a password: ")
		_, err := fmt.Scanln(&password)
		if err == nil {
			password = strings.TrimSpace(password)
			if password != "" {
				break
			}
		}
	}

	return password
}

type ipOpt struct {
	name string
	ip   string
}

func publicIp() string {
	client := http.Client{
		Timeout: 1 * time.Second,
	}
	res, err := client.Get("https://ifconfig.me/ip")
	if err == nil {
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err == nil {
			return string(body)
		}
	}
	return ""
}

func initIp() string {
	maxW := len("public")
	ipOpts := make([]ipOpt, 0)
	pubIp := publicIp()
	if pubIp != "" {
		ipOpts = append(ipOpts,
			ipOpt{
				name: "public",
				ip:   pubIp,
			},
		)
	}

	// get all network interfaces
	interfaces, err := net.Interfaces()
	if err == nil {
		// iterate over network interfaces
		for _, iface := range interfaces {
			// check if this interface is connected
			if iface.Flags&net.FlagUp == 0 {
				continue
			}
			// get this interface's ip addresses
			addrs, err := iface.Addrs()
			if err != nil {
				continue
			}

			// loop through each address
			for _, addr := range addrs {
				switch v := addr.(type) {
				case *net.IPNet:
					// only ipv4
					if v.IP.To4() != nil {
						ipOpts = append(ipOpts,
							ipOpt{
								name: iface.Name,
								ip:   v.IP.String(),
							},
						)
						if len(iface.Name) > maxW {
							maxW = len(iface.Name)
						}
					}
				}
			}
		}
	}

	var ipv4Regex = `^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`
	re := regexp.MustCompile(ipv4Regex)

	fmt.Println("Which network interface and ip address do you want to use?")
	if pubIp != "" {
		for i, opt := range ipOpts {
			fmt.Printf("%d) %-*s  %s\n", i, maxW, opt.name, opt.ip)
		}
	} else {
		for i, opt := range ipOpts {
			fmt.Printf("%d) %-*s  %s\n", i+1, maxW, opt.name, opt.ip)
		}
	}

	var ip string
	for ip == "" {
		fmt.Printf("Select %d-%d or manually input ip address: ", 0, len(ipOpts)-1)
		_, err = fmt.Scanln(&ip)
		if err != nil {
			continue
		}
		ip = strings.TrimSpace(ip)
		// check if ip is a manual ip address
		if re.MatchString(ip) {
			return ip
		}
		ipIndex, err := strconv.Atoi(ip)
		if err == nil {
			if pubIp == "" {
				ipIndex--
			}
			if ipIndex >= 0 && ipIndex < len(ipOpts) {
				return ipOpts[ipIndex].ip
			}
		}
		ip = ""
	}

	return ip
}

func initReactive() string {
	var res string
	fmt.Print("Reactive buzzers? Y/n ")
	_, err := fmt.Scanln(&res)
	if err != nil {
		return "true"
	}
	res = strings.ToLower(strings.TrimSpace(res))
	if slices.Contains([]string{"no", "n"}, res) {
		return "false"
	}

	return "true"
}
