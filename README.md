### Aven 🔐

Aven Provides Caesar And XOR Encryption Through A Clean, Modular CLI — Built For Speed And Extensibility.

#### 🚀 Features

- 🔐 Supports Caesar And XOR Encryption/Decryption
- 🧭 Intuitive CLI Interface Powered By Go's `flag` Package
- 🧱 Modular Architecture For Easy Expansion
- ⚡ Lightning-Fast Execution With Zero Dependencies
- 🌍 Fully Written In English For Global Accessibility

#### 📦 Installation

##### 📥 Clone The Repository

```bash
git clone https://github.com/shermsql/Aven.git
```

##### 📁 Go To Project Directory

```bash
cd Aven
```

##### ⚙️ Compile The Project

```bash
go build .\cmd\aven
```

#### 🧱 Platform Builds

Aven Can Be Built For Multiple Platforms Using Go's Cross-Compilation.

##### 🪟 Windows (64-Bit)

```bash
GOOS=windows GOARCH=amd64 go build -o Aven.exe .\cmd\aven
```

##### 🍎 macOS (Intel & Apple Silicon)

```bash
GOOS=darwin GOARCH=amd64 go build -o Aven .\cmd\aven
```

##### 🐧 Linux (64-Bit)

```bash
GOOS=linux GOARCH=amd64 go build -o Aven .\cmd\aven
```

##### 🐧 Linux (ARM)

```bash
GOOS=linux GOARCH=arm64 go build -o Aven .\cmd\aven
```

##### ℹ️ All Builds Produce A Standalone Binary With No External Dependencies.

##### 🧪 Tested On Windows, macOS, And Linux — Aven Runs Smoothly Across Platforms.

#### 🛠️ Usage

```bash
Aven -m <Mode> -a <Algorithm> -t <Text>
```

#### 📑 Parameters

| Flag     | Type    | Description       | Example               |
|----------|---------|-------------------|-----------------------|
| `-m`     | String  | Mode Of Operation | `Encrypt` / `Decrypt` |
| `-a`     | String  | Algorithm To Use  | `Caesar` / `XOR`      |
| `-t`     | String  | Text To Process   | `"Hello World"`       |
| `-h` | Bool    | Show Help Message | `-h`              |

#### 📖 Examples

##### 🔐 Caesar Encryption

Encrypts The Input Text Using Caesar Cipher With A Default Shift Of `3`.

```bash
go run .\cmd\aven -m Encrypt -a Caesar -t "Hello World!"
```

##### 🔓 Caesar Decryption

Decrypts A Caesar-Encrypted String Back To Its Original Form.

```bash
go run .\cmd\aven -m Decrypt -a Caesar -t "Khoor Zruog!"
```

##### 🔐 XOR Encryption

Encrypts The Input Using XOR Cipher With A Default Key Of `7`.

```bash
go run .\cmd\aven -m Encrypt -a XOR -t "Secret 123"
```

##### 🔓 XOR Decryption

Decrypts An XOR-Encrypted String Using The Same Key.

```bash
go run .\cmd\aven -m Decrypt -a XOR -t "<Encrypted Output>"
```

##### 🆘 Help Output

Displays Aven's ASCII Logo And Help Message When Parameters Are Missing Or `-h` Is Used.

```bash
go run .\cmd\aven -h
```

#### 🗂️ Project Structure

Aven Follows A Clean, Layered Architecture — Separating CLI Logic, Argument Parsing, And Cryptographic Algorithms.

```bash
aven/
├── cmd/aven/        # Entry Point
├── internal/args/   # CLI Argument Parser
├── internal/crypto/ # Encryption Algorithms
├── internal/ui/     # CLI Output
│   ├── Help.txt     # Help Text
│   ├── print.go     # Stylized CLI Output
│   └── ui.go        # Help Printer Logic
├── go.mod           # Go Module Definition
├── README.md        # Project Documentation
├── LICENSE          # MIT License
```

#### 📄 License

Licensed Under The MIT License — Feel Free To Use, Modify, And Share With Proper Attribution. Encryption Belongs To Everyone.

#### 💡 Philosophy & Motivation

Aven Was Born From A Desire To Combine Cryptographic Simplicity With Architectural Elegance.
It Aims To Be More Than Just A Tool — Aven Is A Statement That Encryption Can Be Both Accessible And Beautiful.
No External Libraries, No Clutter — Just Pure Go, Clean Design, And A Focus On Clarity.
The Project Reflects A Belief That Technical Precision And Emotional Resonance Can Coexist.

#### 🛣️ Roadmap

- 🔐 Add AES Encryption Support
- 📁 Enable File-Based Encryption And Decryption
- 🧪 Implement Unit Tests For Core Algorithms
- 📚 Expand Documentation With Visual Examples And Tutorials
- 🌐 Add Localization Support For CLI Messages
- 🧩 Introduce Plugin System For Custom Algorithms

#### 🤝 Contributing

##### 💬 Contributions Are Welcome — Whether It's Code, Documentation, Or Ideas.

##### 🛠️ To Get Started

- 🍴 Fork The Repository
- 🌿 Create A Feature Branch
- 📤 Submit A Pull Request With Clear Description
- 🧘 Respect The Project's Minimalist And Modular Philosophy

##### 🧪 No External Inspirations Were Used — Aven Is Built Entirely From Scratch, With Original Architecture And Design.