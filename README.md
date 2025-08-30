# 🛡️ PHPUnit Go Scanner (CVE-2017-9841)

A fast multithreaded scanner written in Go for detecting exposed and vulnerable `eval-stdin.php` endpoints in PHPUnit (CVE-2017-9841). Works across multiple domains with support for parallel scanning and auto-protocol detection.

---

## 🔍 CVE-2017-9841 Summary

> **CVE-2017-9841** is a critical remote code execution (RCE) vulnerability in **PHPUnit**, caused by the public exposure of the `eval-stdin.php` script.  
> Attackers can execute arbitrary PHP code on the server by sending crafted input to this file.

**Affected versions:**
- PHPUnit ≤ 4.8.28
- PHPUnit ≤ 5.6.2

**Exploitable file path:**
```
/vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php
```

---

## ⚙️ Features

- ✅ Detects exposed `eval-stdin.php` in multiple common paths
- ✅ Supports HTTP/HTTPS auto-detection
- ✅ Multithreaded for speed (`-t` flag)
- ✅ Saves results to organized output files
- ✅ Simple, fast, no dependencies

---

## 🧱 Command Run with Open Source (GO)

Just run your Go source file directly with:

```bash
go run PHPUnit_GoScan.go -l list.txt -t 20
```

---

## 🚀 How to Use

### Basic syntax:
```bash
./PHPUnit_GoScan_amd64_linux -l list.txt -t 20
```

Or on Windows:
```cmd
PHPUnit_GoScan_amd64_windows.exe -l list.txt -t 20
```

### Parameters

| Flag       | Description                                                  |
|------------|--------------------------------------------------------------|
| `-l`      | Path to input file containing one domain per line            |
| `-t` | Number of concurrent scan threads (default: 10 recommended)  |

---

## 📁 Example `list.txt`

```
example.com
http://target.org
https://vulnerable.site/
testdomain.net
```

You can mix raw domains and full URLs. The tool auto-detects protocol if missing.

---

## 📂 Paths Scanned

These common PHPUnit paths will be checked against each domain:

```
/vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php
/phpunit/phpunit/src/Util/PHP/eval-stdin.php
/phpunit/src/Util/PHP/eval-stdin.php
```

You can customize this list in `PHPUnit_PayloadList` inside your source code.

---

## 📦 Output Files

After scanning, the following files will be generated:

| File Name              | Description                                 |
|------------------------|---------------------------------------------|
| `Domain_Online.txt`    | Domains that responded successfully         |
| `PHPUnit_Injected.txt` | Domains confirmed vulnerable to CVE-2017-9841 |

---

## 🧪 Example Console Output

```
[1/100] [Domain Online]   ==> https://target.com
[1/100] [PHPUnit == PWNED] ==> https://target.com/vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php

[5/100] [Domain Offline]  ==> http://dead.site
[6/100] [Fail Injection]  ==> https://clean.site
```

---

## 📥 Download Prebuilt Binaries

If you don’t want to build manually, use these (place in root folder):

| Platform  | File Name               |
|-----------|-------------------------|
| Windows   | `PHPUnit_GoScan_amd64_windows.exe`  |
| Linux     | `PHPUnit_GoScan_amd64_linux`    |

Make sure to run `chmod +x PHPUnit_GoScan_amd64_linux` before executing on Linux.

---

## ⚠️ Legal Warning

This tool is for **educational purposes and authorized security testing only**.  
**Do NOT scan domains you do not own or have permission to test.**  
Unauthorized scanning and exploitation may be illegal and punishable by law.

---

## ✍️ Author

Developed by [DRCrypter.ru](https://drcypter.ru)  
Telegram: `@drcrypterd0tru`  
GitHub: [https://github.com/drcrypterdotru](https://github.com/drcrypterdotru)

---

## 📘 License

MIT License — free to use, modify, distribute.