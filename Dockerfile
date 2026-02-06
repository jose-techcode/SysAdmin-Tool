# Use an official Ubuntu image

FROM debian:latest

# Update and install every utility of CLI tool Luasys

RUN apt update && apt install -y \
    dmidecode \
    pciutils \
    lm-sensors \
    usbutils \
    && apt clean

# Copy the binary

COPY sys /usr/local/bin/sys

# Gives permission to the binary

RUN chmod +x /usr/local/bin/sys

# Command of binary

ENTRYPOINT ["/usr/local/bin/sys"]