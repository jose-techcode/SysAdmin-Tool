# 1. Installation

Sys tool .deb package.

- To install, download the .deb file, navigate to the directory where the .deb file is located and run:

`sudo dpkg -i sys_1.0.9.deb`

- Then, run:

`sys --help`

- **Note:**

Run the CLI tool with administrator permissions (sudo) for proper operation.

# 2. CLI Luasys

The "SysAdmin-Tool" is am CLI tool for auditing, diagnosing, and monitoring hardware resources, such as CPU, GPU, motherboard, BIOS, memory (ram and swap), disk and partitions, USB devices, host (kernel and operating system), net (network), and processes. Net and processes do not appear in the general command, they only work in subcommands. You need administrator or sudo permission to run the CLI Sys

# 3. Features

- `sys` - Displays general information, excluding net and process. Admin/Sudo.

- `api` - Starts the API HTTP server. Localhost. The api subcommand does not work inside Docker. Admin/Sudo.

- `cpu` - Shows CPU information.

- `gpu` - Shows GPU information.

- `motherboard` - Shows motherboard information. Admin/Sudo.

- `bios` - Shows BIOS information. Admin/Sudo.

- `mem` - Shows memory information.

- `disk` - Shows disk information.

- `usb` - Shows USB device information.

- `host` - Shows kernel and operating system information.

- `net` - Shows network information. It works only through subcommand.

- `processes` - Shows process information. It works only through subcommand.

# 4. Technologies

- Language: Go (1.25.7+)
- Shell Scripting: Bash (5.3+)
- Framework: Echo (4.15.0+)
- Libraries: Cobra (1.10.2+) & Gopsutil(3.24.5+)
- Environment: Linux/Debian(13+)/Fedora(42+)
- Code versioning: Git (2.53.0+)
- Containerization: Docker (29.2.1+)
- Scanner: Trivy (0.69.1+)
- CI: Github Actions
- Others: dmidecode, pciutils, usbutils

# 5. API

- GET: Displays auditing, diagnostic, and monitoring information via an HTTP server via API.

- `/` - General path, excluding net and process. Admin/Sudo.

- `/cpu` - CPU path.

- `/gpu` - GPU path.

- `/motherboard` - Motherboard path. Admin/Sudo.

- `/bios` - BIOS path. Admin/Sudo.

- `/mem` - Memory path.

- `/disk` - Disk path.

- `/usb` - USB device path.

- `/host` - Kernel and operating system path.

- `/net` - Network path. It only works through a specific route.

- `/processes` - Process path. It only works through a specific route.

# 6. Clone the Repository

- Bash

`git clone https://github.com/jose-techcode/SysAdmin-Tool`

# 7. Project Folder

- Bash

`cd SysAdmin-Tool`

# 8. Run in Docker

- Build the binary in the same directory of Dockerfile

`go build -o sys`

- Build the image

`docker build -t sys:latest .`

- Scan the Docker image with Trivy (false positives are expected)

`trivy image sys:latest`

- Run the image with privileges

`docker run --rm --privileged sys:latest` (optional: < subcommand >)

# 8.1. Run in Docker (with bash)

- Execute this script to build binary and Docker image and scan it with Trivy

`./docker_trivy.sh`

- Run the image with privileges

`docker run --rm --privileged sys:latest` (optional: < subcommand >)

# 8.2. Run in Docker (with makefile)

- Execute this makefile command to build binary and Docker image and scan it with Trivy

`make image`

- Run the image with privileges

`docker run --rm --privileged sys:latest` (optional: < subcommand >)

# 9. Contribute

Feel free to open Issues or submit Pull Requests.

# 10. License

This project is licensed under the GPL 3.0.

# 11. Notes

The api subcommand does not work inside Docker. In the future, maybe goreleaser will be added to manage package versions, mainly for a possible cross-platform idea as well. Tests of new technologies or features may appear. Finished.