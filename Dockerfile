# Use an official Debian image

FROM debian:latest

# Update and install every utility of CLI Tool Sys (Debian)

RUN apt update && apt install -y \
    dmidecode \
    pciutils \
    usbutils \
    && apt clean

# Copy the binary

COPY sys /usr/local/bin/sys

# Gives permission to the binary

RUN chmod +x /usr/local/bin/sys

# Command of binary

ENTRYPOINT ["/usr/local/bin/sys"]