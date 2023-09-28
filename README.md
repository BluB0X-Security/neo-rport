<!-- markdownlint-disable -->

## At a glance

<!-- markdownlint-restore -->

**Is Neo-Rport a replacement for TeamViewer?**
Yes and no. It depends on your needs.
TeamViewer and a couple of similar products are focused on giving access to a remote graphical desktop bypassing the
Remote Desktop implementation of Microsoft. They fall short in a heterogeneous environment where access to headless
Linux machines is needed. But they are without alternatives for Windows Home Editions.
Apart from remote management, they offer supplementary services like Video Conferences, desktop sharing, screen
mirroring, or spontaneous remote assistance for desktop users.

**Goal of Neo-Rport**
Neo-Rport focuses only on remote management of those operating systems where an existing login mechanism can be used.
It can be used for Linux and Windows, but also appliances and IoT devices providing a web-based configuration.
From a technological perspective, [Ngrok](https://ngrok.com/) and [openport.io](https://openport.io) are similar
products. Neo-Rport differs from them in many aspects.

* Neo-Rport is 100% open source.
* Neo-Rport will come with a user interface making the management of remote systems easy and user-friendly.
* Neo-Rport is made for all operating systems with native and small binaries. No need for Python or similar heavyweights.
* Neo-Rport allows you to self-host the server.
* Neo-Rport allows clients to wait in standby mode without an active tunnel. Tunnels can be requested on-demand by the user remotely.

**Supported operating systems**
For the client almost all operating systems are supported, and we provide binaries for a variety of Linux architectures
and Microsoft Windows.
Also, the server can run on any operating system supported by the golang compiler. At the moment we provide server
binaries only for Linux X64 because this is the ideal platform for running it securely and cost-effective.

## Instantly launch an Neo-RPort server

[![Button Launch RPort  Now](https://img.shields.io/badge/RPort_Server-Launch_Now-brightgreen?style=for-the-badge&logo=Windows%20Terminal)](https://kb.rport.io/install-the-rport-server)

🚀 If you are curious and want to try Neo-RPort, install your server now in no time. Use our
[server installation script](https://kb.rport.io/install-the-rport-server).

## 📖 Documentation

### End-User documentation

[![end-user-documentation](https://img.shields.io/badge/End--User_Documentation-Read_Now-green?style=for-the-badge&logo=Gitbook)](https://kb.rport.io)

our [end-user documentation](https://kb.rport.io) with articles on user-friendly installation, settings and secure operation of RPort.

### Technical documentation

[![technical-documentation](https://img.shields.io/badge/Technical_Documentation-Read_Now-orange?style=for-the-badge&logo=Github)](https://oss.rport.io/)

our [technical documentation](https://oss.rport.io) with all background information and underlying concepts
