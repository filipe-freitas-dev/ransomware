
# 🛡️ Ransomware Simulation Tool

This project is a simulated ransomware application developed in Go, designed for educational and research purposes. It aims to demonstrate the behavior of ransomware in a controlled environment, helping cybersecurity students and professionals understand its mechanics.

## ⚠️ Disclaimer

**This tool is intended solely for educational and ethical research purposes. Do not deploy this code in any unauthorized or live environment. The author is not responsible for any misuse of this software.**

## 🎯 Objectives

- Demonstrate how ransomware encrypts files in a target directory
- Simulate key generation and storage for decryption
- Showcase command-line control and execution of ransomware behavior
- Encourage awareness and preparation against ransomware threats

## 🛠️ Technologies Used

- **Go (Golang)** — The core language used to implement file encryption and CLI interface
- **AES Encryption** — For simulating the ransomware behavior
- **Command-Line Interface** — For control and operation

## 📦 Project Structure

```
ransomware/
├── main.go             # Main application logic
├── crypto.go           # AES encryption/decryption functions
├── keygen.go           # Key generation and storage
└── README.md
```

## 🚀 How to Use

> **Note:** Only run this in a safe and isolated environment such as a virtual machine.

1. **Clone the repository**
```bash
git clone https://github.com/filipe-freitas-dev/ransomware.git
cd ransomware
```

2. **Build the project**
```bash
go build -o ransomware
```

3. **Run the simulation**
```bash
./ransomware encrypt /path/to/test-directory
```
The program will encrypt files inside the specified directory.

or 

```bash
./ransomware decrypt /path/to/test-directory
```
The program will decrypt files inside the specified directory encrypted by this program.

## 🔐 Decryption

A simulated decryption process is available for recovery testing. Make sure to save the encryption key that is generated and displayed during execution.

## 📚 Educational Use Cases

- Cybersecurity labs and simulations
- Red team / blue team training exercises
- Understanding ransomware mechanisms in a safe environment

## 🧪 Testing

*Tests and verification steps can be added as needed to validate encryption and decryption logic.*

## 📄 License

MIT License. See [LICENSE](LICENSE) file for details.

