# Speed Snitch Agent

[![Go Report Card](https://goreportcard.com/badge/github.com/silinternational/speed-snitch-agent)](https://goreportcard.com/report/github.com/silinternational/speed-snitch-agent)

This repo includes the source for our Speed Snitch agent which we use to monitor and report on internet speeds
at remote locations. The agent communicates with a management API to receive configuration details related to Tasks 
(speed tests, latency tests, etc.). 

This project is intended to work in conjunction with the `speed-snitch-admin-api` and `speed-snitch-admin-ui` projects
for the central management capabilities.  They can be found at github.com/silinternational.

## Build Instructions
1. Run `make dep`
2. Build binary (this will create `cmd/speedsnitch/speedsnitch`):
    - Mac: `make mac`
    - Linux: `make linux`
    - Windows: `make windows`

## Debian Package
1. Edit fpm.env based on the contents of fpm.env.dist but using the correct values
2. Run `make fpm`
3. The package will be created in ./docker/fpm as speedsnitch.linux_0.1.2_amd64.deb

## Release Process
1. Create new release branch for version number from `develop` following convention `x.x.x`
   (note: don't include `release` in the name.)
2. Update source with new version number
4. Commit/push changes
5. Create PR from `x.x.x` to `develop`
6. After PR reviewed and merged, create PR from `develop` to `master`
7. Review and merge PR to `master`


## Special Instructions for systemd on a Raspberry Pi
On a Raspberry Pi, you can have systemd ensure the speedsnitch executable is always running and
restarts following an update.

1. Edit extras/pi_lib_systemd_system/speedSnitchAgent.service to have the correct values.
2. Copy (as root) the following files into /usr/lib/systemd/system/ ...
    - extras/pi_lib_systemd_system/speedSnitchAgent.service
    - extras/pi_lib_systemd_system/speedSnitchWatcher.service
    - extras/pi_lib_systemd_system/speedSnitchWatcher.path
3. Create symlinks for these files ...
    - $ cd /etc/systemd/system
    - $ mkdir -p default.target.wants
    - $ cd default.target.wants
    - $ sudo ln -s /usr/lib/systemd/system/speedSnitchAgent.service speedSnitchAgent.service
    - $ sudo ln -s /usr/lib/systemd/system/speedSnitchWatcher.service speedSnitchWatcher.service
    - $ sudo ln -s /usr/lib/systemd/system/speedSnitchWatcher.path speedSnitchWatcher.path
4. In the /usr/lib/systemd/system/speedSnitchAgent.service file replace the parameters in the `ExecStart` value.
6. Reload the systemd daemons and start the new services
    - $ sudo systemctl daemon-reload
    - $ sudo systemctl start speedSnitchWatcher.path
    - $ sudo systemctl start speedSnitchWatcher.service
7. In order to view the status ... $ sudo systemctl status speedSnitchAgent.service
8. In order to see more lines from its output ... $ journalctl -e -u speedSnitchAgent.service


## License - MIT
MIT License

Copyright (c) 2018 SIL International

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.