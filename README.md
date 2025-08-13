# 1. Installation

Luasys tool .deb package.

- To install, download the .deb file, navigate to the directory where the .deb file is located and run:

sudo dpkg -i luasys_1.0.7.deb

- Then, run:

luasys --help

- **Note:**

Run the CLI tool with administrator permissions (sudo) for proper operation.

# 2. CLI Luasys

The "CLI Luasys" is a multifunctional CLI tool for auditing, diagnosing, and monitoring hardware resources, such as CPU, GPU, motherboard, BIOS, memory (ram and swap), disk and partitions,
temperature, USB devices, host (kernel and operating system), net (network), and processes. Net and processes do not appear in the general command, they only work in subcommands. You need administrator or sudo permission to run the CLI Luasys. By cloning the repository, you will have control of a local REST GET API, Docker, Makefile and one Bash script.

# 3. Features

- luasys - Displays general information, excluding net and process.

- api - Starts the API HTTP server. Localhost. The api subcommand does not work inside Docker.

- cpu - Shows CPU information.

- gpu - Shows GPU information.

- motherboard - Shows motherboard information.

- bios - Shows BIOS information.

- mem - Shows memory information.

- disk - Shows disk information.

- temp - Shows temperature information.

- usb - Shows USB device information.

- host - Shows kernel and operating system information.

- net - Shows network information. It works only through subcommand.

- processes - Shows process information. It works only through subcommand.

# 4. Technologies

- Language: Golang
- Shell Scripting: Bash
- Framework: Echo
- Libraries: Cobra & Gopsutil
- Environment: Linux (Debian & Ubuntu based)
- File format: Makefile & Yaml 
- Code versioning: Git
- Containerization: Docker
- CI: Github Actions
- Scanner (docker image): Trivy
- Others: dmidecode, lspci, sensors, and lsusb

# 5. API

- GET: Displays auditing, diagnostic, and monitoring information via an HTTP server via API.

- / - General path, excluding net and process.

- /cpu - CPU path.

- /gpu - GPU path.

- /motherboard - Motherboard path.

- /bios - BIOS path.

- /mem - Memory path.

- /disk - Disk path.

- /temperature - Temperature path.

- /usb - USB device path.

- /host - Kernel and operating system path.

- /net - Network path. It only works through a specific route.

- /processes - Process path. It only works through a specific route.

# 6. Clone the Repository

- Bash

git clone https://github.com/jose-techcode/CLI_Luasys

# 7. Project Folder

- Bash

cd CLI_Luasys

# 8. Run in Docker

- Build the binary in the same directory of Dockerfile

go build -o luasys

- Build the image

docker build -t luasys:latest .

- Scan the Docker image with Trivy (false positives are expected)

trivy image luasys:latest

- Run the image with privileges

docker run --rm --privileged luasys:latest (optional: < subcommand >)

# 8.1. Run in Docker (with bash)

- Execute this script to build binary and Docker image and scan it with Trivy

./docker_trivy.sh

- Run the image with privileges

docker run --rm --privileged luasys:latest (optional: < subcommand >)

# 8.2. Run in Docker (with makefile)

- Execute this makefile command to build binary and Docker image and scan it with Trivy

make image

- Run the image with privileges

docker run --rm --privileged luasys:latest (optional: < subcommand >)

# 9. Contribute

Feel free to open Issues or submit Pull Requests.

# 10. License

This project is licensed under the GPL 3.0.

# 11. Notes

The api subcommand does not work inside Docker. In the future, maybe goreleaser will be added to manage package versions, mainly for a possible cross-platform idea as well. Tests of new technologies or features may appear.