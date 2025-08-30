# 🛡️ PHPUnit Go Scanner (CVE-2017-9841)

A fast, multithreaded scanner written in Go for detecting exposed and vulnerable `eval-stdin.php` endpoints in PHPUnit (CVE-2017-9841). Supports scanning across multiple domains with parallel execution and automatic protocol detection.

---

## 🖼 Demo Screenshot

![demo](https://raw.githubusercontent.com/drcrypterdotru/PHPUnit-GoScan/refs/heads/main/demo.png)

---

## 🔍 CVE-2017-9841 Summary

> **CVE-2017-9841** is a critical remote code execution (RCE) vulnerability in **PHPUnit**, caused by the public exposure of the `eval-stdin.php` script.  
> Attackers can execute arbitrary PHP code on the server by sending crafted input to this endpoint.

**Affected versions:**
- PHPUnit ≤ 4.8.28
- PHPUnit ≤ 5.6.2

**Common vulnerable path:**
```
/vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php
```

---

## ⚙️ Features

- ✅ Detects exposed `eval-stdin.php` files using multiple common paths
- ✅ Automatically detects HTTP/HTTPS protocol
- ✅ High-speed multithreaded scanning via `-t` flag
- ✅ Clean, categorized output to result files
- ✅ No third-party dependencies — just Go

---

## 🚀 Usage

### 🔧 Command-line Execution
```bash
go run PHPUnit_GoScan.go -l list.txt -t 20
```

Or use precompiled binaries:

#### On Linux:
```bash
chmod +x PHPUnit_GoScan_amd64_linux
./PHPUnit_GoScan_amd64_linux -l list.txt -t 20
```

#### On Windows:
```cmd
PHPUnit_GoScan_amd64_windows.exe -l list.txt -t 20
```

---

### 📌 Parameters

| Flag   | Description                                         |
|--------|-----------------------------------------------------|
| `-l`   | Path to input file with one domain per line         |
| `-t`   | Number of concurrent threads (default: 10, recommended: 20) |

---

## 📁 Input: `list.txt`

Example domain list:
```
example.com
http://target.org
https://vulnerable.site/
testdomain.net
```

- Supports raw domains and full URLs
- Automatically adds protocol if missing

---

## 🔎 Paths Scanned

The scanner checks for the following common vulnerable paths:
```
/vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php
/phpunit/phpunit/src/Util/PHP/eval-stdin.php
/phpunit/src/Util/PHP/eval-stdin.php
```

You can customize these in the `PHPUnit_PayloadList` section of the Go source file.

---

## 📦 Output Files

After scanning, results are written to:

| File Name              | Description                                 |
|------------------------|---------------------------------------------|
| `Domain_Online.txt`    | Domains that responded with HTTP 200        |
| `PHPUnit_Injected.txt` | Domains confirmed vulnerable to CVE-2017-9841 |

---

## 🧪 Example Console Output

```
[1/100] [Domain Online]     ==> https://target.com
[1/100] [PHPUnit == PWNED]  ==> https://target.com/vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php

[5/100] [Domain Offline]    ==> http://dead.site
[6/100] [Fail Injection]    ==> https://clean.site
```

---

## 📥 Download Prebuilt Binaries

| Platform | File Name                          |
|----------|-------------------------------------|
| Linux    | [PHPUnit_GoScan_amd64_linux](https://github.com/drcrypterdotru/PHPUnit-GoScan/releases/download/v1.0.0/PHPUnit_GoScan_amd64_linux) |
| Windows  | [PHPUnit_GoScan_amd64_windows.exe](https://github.com/drcrypterdotru/PHPUnit-GoScan/releases/download/v1.0.0/PHPUnit_GoScan_amd64_windows.exe) |

> 🛠 On Linux:  
> Run `chmod +x PHPUnit_GoScan_amd64_linux` before executing.


---

## ⚠️ Legal Warning

This tool is for **educational and authorized security testing only**.  
Do **not** scan domains you do not own or lack permission to test.  
Unauthorized use may be illegal and punishable under applicable laws.

---

## ✍️ Author

Developed by [DRCrypter.ru](https://drcypter.ru)  
Telegram: [`@drcrypterd0tru`](https://t.me/drcrypterd0tru)  
GitHub: [@drcrypterdotru](https://github.com/drcrypterdotru)

---

## 📘 License

**MIT License** — Free to use, modify, and distribute.  
Links:
- [Source Code: PHPUnit_GoScan.go](https://github.com/drcrypterdotru/PHPUnit-GoScan/blob/main/PHPUnit_GoScan.go)
- [Linux Binary](https://github.com/drcrypterdotru/PHPUnit-GoScan/blob/main/PHPUnit_GoScan_amd64_linux)
- [Windows Binary](https://github.com/drcrypterdotru/PHPUnit-GoScan/blob/main/PHPUnit_GoScan_amd64_windows.exe)